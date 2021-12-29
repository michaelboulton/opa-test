package main

import (
	"github.com/michaelboulton/opa-test/pkg/cmd"
	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use: "opa-test",
}

func main() {
	rootCmd.AddCommand(cmd.RunCmd, cmd.ServeCmd)

	err := rootCmd.Execute()
	if err != nil {
		panic(err)
	}
}
