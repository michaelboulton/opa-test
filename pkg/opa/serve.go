package opa

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/pkg/errors"
	"go.uber.org/zap/zapcore"
)

// ServePolicy serves a policy file with the given name
func ServePolicy(bundleName string, token string, policyFile string, addr string) *http.Server {
	infoLogger := logger.WriteAtLevel(zapcore.InfoLevel)
	// gin.DefaultWriter = infoLogger
	gin.DefaultWriter = ioutil.Discard
	gin.DefaultErrorWriter = logger.WriteAtLevel(zapcore.ErrorLevel)

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
		GET(fmt.Sprintf("/bundles/%s", bundleName), func(context *gin.Context) {
			if context.GetHeader("authorization") == "" {
				_ = context.AbortWithError(401, errors.New("no auth header"))
				return
			}
			if context.GetHeader("authorization") != "Bearer "+token {
				_ = context.AbortWithError(403, errors.New("bad auth header"))
				return
			}

			context.File(policyFile)
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