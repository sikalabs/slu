package connect_from_vault

import (
	parent_cmd "github.com/sikalabs/slu/cmd/scripts/kubernetes"
	"github.com/sikalabs/slu/utils/k8s_scripts"
	"github.com/spf13/cobra"
)

var FlagVaultAddress string
var FlagClusterName string

var Cmd = &cobra.Command{
	Use:     "connect-from-vault",
	Short:   "Connect to Kubernetes cluster using kubeconfig from Vault",
	Aliases: []string{"cfv"},
	Args:    cobra.NoArgs,
	Run: func(c *cobra.Command, args []string) {
		k8s_scripts.ConnectFromVault(FlagVaultAddress, FlagClusterName)
	},
}

func init() {
	parent_cmd.Cmd.AddCommand(Cmd)
	Cmd.Flags().StringVarP(
		&FlagVaultAddress,
		"vault-address",
		"a",
		"",
		"Vault address",
	)
	Cmd.MarkFlagRequired("vault-address")
	Cmd.Flags().StringVarP(
		&FlagClusterName,
		"cluster-name",
		"c",
		"",
		"Cluster name",
	)
	Cmd.MarkFlagRequired("cluster-name")
}
