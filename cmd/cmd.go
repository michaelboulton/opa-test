package main

import (
	"context"

	"github.com/michaelboulton/opa-test/pkg"
	"github.com/spf13/cobra"
)

func main() {
	err := cmd.Execute()
	if err != nil {
		panic(err)
	}
}

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

	opa, err := pkg.NewOpa(ctx, filename)
	if err != nil {
		return err
	}

	defer opa.Stop(ctx)

	return nil
}
