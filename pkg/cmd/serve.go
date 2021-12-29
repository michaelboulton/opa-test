package cmd

import (
	"context"

	"github.com/michaelboulton/opa-test/pkg/opa"
	"github.com/spf13/cobra"
)

func AddServeCmd(r Registrar) {
	var address string
	var authToken string
	var bundleFile string
	var bundleName string

	var ServeCmd = &cobra.Command{
		Use:     "serve",
		Example: "opa-test serve mybundlename bundle.tar.gz",
		Short:   "serves an OPA bundle",
		Args:    cobra.NoArgs,
		RunE: func(cmd *cobra.Command, args []string) error {
			return doServe(cmd.Context(), address, authToken, bundleName, bundleFile)
		},
	}

	flags := ServeCmd.Flags()
	flags.StringVar(&address, "address", "127.0.0.1:8898", "address to serve on")
	flags.StringVar(&authToken, "auth-token", "", "authorization")
	flags.StringVar(&bundleFile, "bundle-file", "", "location of bundle file to server")
	flags.StringVar(&bundleName, "bundle-name", "", "name of bundle")

	_ = ServeCmd.MarkFlagRequired("auth-token")
	_ = ServeCmd.MarkFlagRequired("bundle-file")
	_ = ServeCmd.MarkFlagRequired("bundle-name")

	r.Register(ServeCmd)
}

func doServe(ctx context.Context, addr string, token string, bundleName string, bundleFile string) error {
	ctx, cancel := context.WithCancel(ctx)
	defer cancel()

	opa.ServePolicy(ctx, addr, token, bundleName, bundleFile)

	return nil
}
