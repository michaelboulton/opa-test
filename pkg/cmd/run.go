package cmd

import (
	"context"

	"github.com/michaelboulton/opa-test/pkg/opa"
	"github.com/open-policy-agent/opa/sdk"
	"github.com/pkg/errors"
	"github.com/spf13/cobra"
)

var RunCmd = &cobra.Command{
	Use:     "run",
	Example: "opa-test run my-config.yaml",
	Short:   "Runs an OPA config",
	Args:    cobra.ExactArgs(1),
	RunE: func(cmd *cobra.Command, args []string) error {
		return doRun(cmd.Context(), args[0])
	},
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
		Path: "/policies/allow_post",
		Input: map[string]interface{}{
			"a": "b",
		},
	})
	if err != nil {
		return errors.Wrap(err, "making decision")
	}

	logger.Infof("Decision: %#v", decision)

	return nil
}
