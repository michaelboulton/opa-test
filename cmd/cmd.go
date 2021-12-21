package main

import (
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
	Use:  "opa-test",
	Args: cobra.NoArgs,
	RunE: doRun,
}

func doRun(command *cobra.Command, args []string) error {
	logger := pkg.Logger

	opa, err := pkg.NewOpa(command.Context())
	if err != nil {
		return err
	}

	logger.Infof("%#v", opa)

	return nil
}
