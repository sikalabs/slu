package install_cluster_issuer_zerossl

import (
	parent_cmd "github.com/sikalabs/slu/cmd/scripts/kubernetes"
	"github.com/sikalabs/slu/utils/k8s_scripts"
	"github.com/spf13/cobra"
)

var FlagDry bool
var FlagEmail string
var FlagKeyID string
var FlagKeySecret string

var Cmd = &cobra.Command{
	Use:     "install-cluster-issuer-zerossl",
	Short:   "Install ZeroSSL Cluster Issuer",
	Aliases: []string{"iciz"},
	Args:    cobra.NoArgs,
	Run: func(c *cobra.Command, args []string) {
		k8s_scripts.InstallClusterIssuerZeroSSL(FlagEmail, FlagKeyID, FlagKeySecret, FlagDry)
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
		"",
		"Email of ZeroSSL account",
	)
	Cmd.MarkFlagRequired("email")
	Cmd.Flags().StringVarP(
		&FlagKeyID,
		"key-id",
		"i",
		"",
		"ZeroSSL KeyID",
	)
	Cmd.MarkFlagRequired("key-id")
	Cmd.Flags().StringVarP(
		&FlagKeySecret,
		"key-secret",
		"s",
		"",
		"ZeroSSL KeySecret",
	)
	Cmd.MarkFlagRequired("key-secret")
}
