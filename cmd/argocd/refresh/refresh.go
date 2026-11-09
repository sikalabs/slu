package refresh

import (
	"fmt"
	"time"

	argocd_cmd "github.com/sikalabs/slu/cmd/argocd"
	"github.com/sikalabs/slu/utils/argocd_utils"

	"github.com/spf13/cobra"
)

var FlagNamespace string
var FlagHard bool
var FlagNoWait bool
var FlagTimeout time.Duration

var Cmd = &cobra.Command{
	Use:     "refresh <app>",
	Short:   "Refresh ArgoCD Application (using the refresh annotation)",
	Aliases: []string{"r"},
	Args:    cobra.ExactArgs(1),
	Run: func(c *cobra.Command, args []string) {
		appName := args[0]

		refreshType := "normal"
		if FlagHard {
			refreshType = "hard"
		}
		wait := FlagTimeout
		if FlagNoWait {
			wait = 0
		}

		err := argocd_utils.ArgoCDRefresh(c.Context(), FlagNamespace, appName, refreshType, wait)
		cobra.CheckErr(err)
		fmt.Printf("Application '%s' refreshed (%s)\n", appName, refreshType)
	},
}

func init() {
	argocd_cmd.Cmd.AddCommand(Cmd)
	Cmd.Flags().StringVarP(
		&FlagNamespace,
		"namespace",
		"n",
		"argocd",
		"Namespace of the ArgoCD Application",
	)
	Cmd.Flags().BoolVar(
		&FlagHard,
		"hard",
		false,
		"Hard refresh (invalidate manifest cache)",
	)
	Cmd.Flags().BoolVar(
		&FlagNoWait,
		"no-wait",
		false,
		"Do not wait for ArgoCD to finish the refresh",
	)
	Cmd.Flags().DurationVar(
		&FlagTimeout,
		"timeout",
		2*time.Minute,
		"How long to wait for ArgoCD to finish the refresh",
	)

	// Deprecated flags from the API-based implementation, kept so existing
	// scripts don't break. They are ignored.
	var ignoredString string
	var ignoredBool bool
	Cmd.Flags().StringVarP(&ignoredString, "server", "s", "", "")
	Cmd.Flags().BoolVar(&ignoredBool, "insecure", false, "")
	Cmd.Flags().StringVarP(&ignoredString, "password", "p", "", "")
	Cmd.Flags().StringVar(&ignoredString, "service-token-name", "", "")
	for _, f := range []string{"server", "insecure", "password", "service-token-name"} {
		Cmd.Flags().MarkDeprecated(f, "refresh now uses the Kubernetes API (refresh annotation), this flag is ignored")
	}
}
