package cmd

import (
	"context"
	"reflect"

	"github.com/michaelboulton/opa-test/pkg/opa"
	"github.com/open-policy-agent/opa/sdk"
	"github.com/pkg/errors"
	"github.com/spf13/cobra"
)

func AddRunCmd(r Registrar) {
	var RunCmd = &cobra.Command{
		Use:     "run",
		Example: "opa-test run my-config.yaml",
		Short:   "Runs an OPA config",
		Args:    cobra.ExactArgs(1),
		RunE: func(cmd *cobra.Command, args []string) error {
			return doRun(cmd.Context(), args[0])
		},
	}

	r.Register(RunCmd)
}

func doRun(ctx context.Context, filename string) error {
	ctx, cancel := context.WithCancel(ctx)
	defer cancel()

	instance, err := opa.NewOpa(ctx, filename)
	if err != nil {
		return err
	}
	defer instance.Stop(ctx)

	decision, err := instance.Decision(ctx, sdk.DecisionOptions{
		Path: "authz/allow",
		Input: map[string]interface{}{
			"a":      "b",
			"path":   "/users",
			"method": "POST",
		},
	})
	if err != nil {
		return errors.Wrap(err, "making decision")
	}

	logger.Infof("Decision: %#v", decision)

	allowed, ok := decision.Result.(bool)
	if !ok {
		return errors.Errorf("Expected boolean decision result but got %#v", reflect.TypeOf(decision.Result))
	}
	if !allowed {
		return errors.New("Was not allowed")
	}

	return nil
}
