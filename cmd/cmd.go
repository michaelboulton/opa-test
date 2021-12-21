package main

import (
	"context"

	"github.com/michaelboulton/opa-test/pkg/logging"
	"github.com/michaelboulton/opa-test/pkg/opa"
	"github.com/open-policy-agent/opa/sdk"
	"github.com/pkg/errors"
	"github.com/spf13/cobra"
)

var logger = logging.Logger

var cmd = &cobra.Command{
	Use:     "opa-test",
	Example: "opa-test my-config.yaml",
	Short:   "Runs an OPA config",
	Args:    cobra.ExactArgs(1),
	RunE:    doRun,
}

func doRun(command *cobra.Command, args []string) error {
	ctx, cancel := context.WithCancel(command.Context())
	defer cancel()

	filename := args[0]

	instance, err := opa.NewOpa(ctx, filename)
	if err != nil {
		return err
	}
	defer instance.Stop(ctx)

	decision, err := instance.Decision(ctx, sdk.DecisionOptions{
		Path: "example/allow",
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

func main() {
	err := cmd.Execute()
	if err != nil {
		logger.Panic(err)
	}
}
