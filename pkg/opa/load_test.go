package opa

import (
	"context"
	"encoding/base64"
	"encoding/json"
	"fmt"
	"io"
	"io/ioutil"
	"math/rand"
	"net/http"
	"os"
	"testing"
	"time"

	"github.com/bazelbuild/rules_go/go/tools/bazel"
	"github.com/gin-gonic/gin"
	"github.com/pkg/errors"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"go.uber.org/zap/zapcore"
	"gopkg.in/yaml.v3"
)

func TestNewOpa(t *testing.T) {
	tests := []struct {
		name string
	}{
		{
			name: "create with bundle",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			ctx, cancel := context.WithTimeout(context.Background(), time.Second*3)
			defer cancel()

			randInput := make([]byte, 20)
			_, err := rand.Read(randInput)
			require.NoError(t, err)
			token := base64.URLEncoding.EncodeToString(randInput)

			i := rand.Int()%40000 + 10000
			addr := fmt.Sprintf("127.0.0.1:%d", i)
			policy := "policies"
			bundle := "authz"

			configFile := createConfigFile(t, addr, policy, bundle, token)

			server := startServingBundles(t, addr, policy, bundle, token)
			defer server.Close()

			opa, err := NewOpa(ctx, configFile.Name())
			require.NoError(t, err)
			require.NotNil(t, opa)
		})
	}
}

func createConfigFile(t *testing.T, addr string, policy string, bundle string, token string) *os.File {
	exampleConfig := OpaConfig{
		Services: []Service{
			{
				Name: "mytestservice",
				URL:  "http://" + addr,
				Credentials: map[string]interface{}{
					"bearer": map[string]string{
						"token": token,
					},
				},
			},
		},
		Bundles: map[string]Bundle{
			bundle: {
				BundleSource: &BundleSource{
					Service:  "mytestservice",
					Resource: policy,
					Persist:  false,
					Polling:  nil,
					Signing:  nil,
				},
			},
		},
	}
	asJson, err := json.Marshal(exampleConfig)
	require.NoError(t, err)
	intermediary := map[string]interface{}{}
	err = yaml.Unmarshal(asJson, &intermediary)
	require.NoError(t, err)
	asYaml, err := yaml.Marshal(intermediary)

	file, err := ioutil.TempFile(bazel.TestTmpDir(), "*.yaml")
	require.NoError(t, err)

	_, err = file.Write(asYaml)
	require.NoError(t, err)

	_, err = file.Seek(0, io.SeekStart)
	require.NoError(t, err)

	t.Logf("File for config: %s", file.Name())

	return file
}

func startServingBundles(t *testing.T, addr string, policy string, bundle string, token string) *http.Server {
	infoLogger := logger.WriteAtLevel(zapcore.InfoLevel)
	// gin.DefaultWriter = infoLogger
	gin.DefaultWriter = ioutil.Discard
	gin.DefaultErrorWriter = logger.WriteAtLevel(zapcore.ErrorLevel)

	runfile, err := bazel.Runfile("policies/bundle.tar.gz")
	if !assert.NoError(t, err) {
		av, err := bazel.ListRunfiles()
		require.NoError(t, err)
		for _, entry := range av {
			t.Logf("%s: %s", entry.Workspace, entry.ShortPath)
		}
		require.FailNow(t, "ohnoes")
	}

	router := gin.Default()
	router.
		Use(gin.LoggerWithConfig(gin.LoggerConfig{
			Formatter: func(params gin.LogFormatterParams) string {
				return params.ErrorMessage
			},
			Output: infoLogger,
		})).
		Use(func(context *gin.Context) {
			t0 := time.Now()
			context.Next()
			duration := time.Since(t0)

			logger.
				WithFields(map[string]interface{}{
					"duration": duration.Seconds(),
					"method":   context.Request.Method,
				}).
				Info("Path: %s", context.Request.URL.Path)
		}).
		GET(fmt.Sprintf("/bundles/%s", bundle), func(context *gin.Context) {
			if context.GetHeader("authorization") == "" {
				_ = context.AbortWithError(401, errors.New("no auth header"))
				return
			}
			if context.GetHeader("authorization") != "Bearer "+token {
				_ = context.AbortWithError(403, errors.New("bad auth header"))
				return
			}

			context.File(runfile)
			context.Status(200)
		})

	server := &http.Server{
		Handler: router,
		Addr:    addr,
	}
	go func() {
		_ = server.ListenAndServe()
	}()

	return server
}
