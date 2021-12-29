package cmd

import "github.com/spf13/cobra"

// Registrar adds a command to the command line
type Registrar interface {
	Register(cmd *cobra.Command)
}
