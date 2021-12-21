package pkg

import (
	"context"
	"os"

	"github.com/open-policy-agent/opa/sdk"
	"github.com/pkg/errors"
)

func NewOpa(ctx context.Context, filename string) (*sdk.OPA, error) {
	file, err := os.Open(filename)
	if err != nil {
		return nil, errors.Wrap(err, "opening config file")
	}

	opa, err := sdk.New(ctx, sdk.Options{
		Config:        file,
		Logger:        Logger,
		ConsoleLogger: Logger,
	})
	if err != nil {
		return nil, errors.Wrap(err, "oh noes")
	}

	Logger.Infof("loaded config from %s", filename)

	return opa, nil
}
