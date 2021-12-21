package opa

import (
	"context"
	"os"

	"github.com/michaelboulton/opa-test/pkg/logging"
	"github.com/open-policy-agent/opa/sdk"
	"github.com/pkg/errors"
)

var logger = logging.Logger

func NewOpa(ctx context.Context, filename string) (*sdk.OPA, error) {
	file, err := os.Open(filename)
	if err != nil {
		return nil, errors.Wrap(err, "opening config file")
	}

	opa, err := sdk.New(ctx, sdk.Options{
		Config:        file,
		Logger:        logger.WithSkip(1),
		ConsoleLogger: logger.WithSkip(1),
	})
	if err != nil {
		return nil, errors.Wrap(err, "creating OPA instance")
	}

	logger.Infof("loaded config from %s", filename)

	return opa, nil
}

// Service defines a service
// services:
//  - name: acmecorp
//    url: https://example.com/service/v1
//    credentials:
//      bearer:
//        token: "bGFza2RqZmxha3NkamZsa2Fqc2Rsa2ZqYWtsc2RqZmtramRmYWxkc2tm"
//
type Service struct {
	Name        string                 `json:"name,omitempty"`
	URL         string                 `json:"url,omitempty"`
	Credentials map[string]interface{} `json:"credentials,omitempty"`
}

// Bundle defines a bundle
// bundles:
//  authz:
//    service: acmecorp
//    resource: somedir/bundle.tar.gz
//    persist: true
//    polling:
//      min_delay_seconds: 10
//      max_delay_seconds: 20
//    signing:
//      keyid: my_global_key
//      scope: read
type Bundle struct {
	BundleSource *BundleSource `json:"authz,omitempty"`
}
type BundleSource struct {
	Service  string   `json:"service,omitempty"`
	Resource string   `json:"resource,omitempty"`
	Persist  bool     `json:"persist,omitempty"`
	Polling  *Polling `json:"polling,omitempty"`
	Signing  *Signing `json:"signing,omitempty"`
}
type Polling struct {
	MinDelaySeconds int `json:"min_delay_seconds,omitempty"`
	MaxDelaySeconds int `json:"max_delay_seconds,omitempty"`
}
type Signing struct {
	Keyid string `json:"keyid,omitempty"`
	Scope string `json:"scope,omitempty"`
}

// OpaConfig defines the top level OPA config to go to json
type OpaConfig struct {
	Services []Service         `json:"services,omitempty"`
	Bundles  map[string]Bundle `json:"bundles,omitempty"`
}
