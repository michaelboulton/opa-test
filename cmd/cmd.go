package main

import (
	"github.com/michaelboulton/opa-test/pkg/cmd"
	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use: "opa-test",
}

func main() {
	r := &registrar{rootCmd: rootCmd}
	cmd.AddServeCmd(r)
	cmd.AddRunCmd(r)

	err := rootCmd.Execute()
	if err != nil {
		panic(err)
	}
}

type registrar struct {
	rootCmd *cobra.Command
}

func (r *registrar) Register(cmd *cobra.Command) {
	r.rootCmd.AddCommand(cmd)
}
