package df

import (
	"context"
	"os"
	"os/signal"
	"syscall"
	"fmt"

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
		log(fmt.Sprintf("%s\n", os.Args))
		log(fmt.Sprintf("%s\n", os.Args[1:]))
		os.Exit(di.NewCmd().Run(ctx, os.Args[1:], version))
	},
}

func init() {
	root.RootCmd.AddCommand(Cmd)
}

func log(line string) {
	filePath := "example.log"

	// Open the file in append mode, create if it doesn't exist
	f, err := os.OpenFile(filePath, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	if err != nil {
		fmt.Println("Error opening file:", err)
		return
	}
	defer f.Close()

	// Write the line
	if _, err := f.WriteString(line); err != nil {
		fmt.Println("Error writing to file:", err)
		return
	}
}
