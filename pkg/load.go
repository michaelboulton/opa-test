package pkg

import (
	"context"

	"github.com/open-policy-agent/opa/sdk"
	"github.com/pkg/errors"
)

func NewOpa(ctx context.Context) (*sdk.OPA, error) {
	opa, err := sdk.New(ctx, sdk.Options{
		Config:        nil,
		Logger:        nil,
		ConsoleLogger: nil,
		Ready:         nil,
		Plugins:       nil,
	})
	if err != nil {
		return nil, errors.Wrap(err, "oh noes")
	}

	return opa, nil
}
