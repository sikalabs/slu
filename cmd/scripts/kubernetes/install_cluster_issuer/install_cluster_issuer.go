package install_cluster_issuer

import (
	"fmt"

	parent_cmd "github.com/sikalabs/slu/cmd/scripts/kubernetes"
	"github.com/sikalabs/slu/utils/k8s_scripts"
	"github.com/sikalabs/slu/utils/sh_utils"
	"github.com/spf13/cobra"
)

const DEFAULT_LETS_ENCRYPT_EMAIL = "lets-encrypt-slu@sikamail.com"

var FlagDry bool
var FlagEmail string

var Cmd = &cobra.Command{
	Use:     "install-cluster-issuer",
	Short:   "Install Let's Encrypt Cluster Issuer",
	Aliases: []string{"ici"},
	Args:    cobra.NoArgs,
	Run: func(c *cobra.Command, args []string) {
		k8s_scripts.InstallClusterIssuer(FlagEmail, FlagDry)
	},
}

func init() {
	parent_cmd.Cmd.AddCommand(Cmd)
	Cmd.Flags().BoolVar(
		&FlagDry,
		"dry",
		false,
		"Dry run",
	)
	Cmd.Flags().StringVarP(
		&FlagEmail,
		"email",
		"e",
		DEFAULT_LETS_ENCRYPT_EMAIL,
		"Email for Let's Encrypt account & notifications",
	)
}

func sh(script string, dry bool) {
	if dry {
		fmt.Println(script)
		return
	}
	err := sh_utils.ExecShOutDir("", script)
	if err != nil {
		sh_utils.HandleError(err)
	}
}
