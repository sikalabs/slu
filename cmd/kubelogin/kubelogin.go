package df

import (
	"context"
	"os"
	"os/signal"
	"syscall"

	"github.com/int128/kubelogin/pkg/di"
	"github.com/sikalabs/slu/cmd/root"
	"github.com/spf13/cobra"
)

var Cmd = &cobra.Command{
	Use:                "kubelogin",
	Short:              "Embedded int128/kubelogin, oidc-login command",
	DisableFlagParsing: true,
	Run: func(c *cobra.Command, args []string) {
		version := "v1.34.0+slu"

		ctx := context.Background()
		ctx, stop := signal.NotifyContext(ctx, os.Interrupt, syscall.SIGTERM)
		defer stop()
		os.Exit(di.NewCmd().Run(ctx, os.Args[1:], version))
	},
}

func init() {
	root.RootCmd.AddCommand(Cmd)
}
