package delete_ns

import (
	k8s_cmd "github.com/sikalabs/slu/cmd/k8s"
	"github.com/sikalabs/slu/utils/k8s"

	"github.com/spf13/cobra"
)

var FlagNamespace string

var Cmd = &cobra.Command{
	Use:   "delete-ns",
	Short: "Delete (stucked) terminating namespace",
	Args:  cobra.NoArgs,
	Run: func(c *cobra.Command, args []string) {
		clientset, _, _ := k8s.KubernetesClient()
		k8s.DeleteTerminatingNamespace(clientset, FlagNamespace)
	},
}

func init() {
	k8s_cmd.Cmd.AddCommand(Cmd)
	Cmd.Flags().StringVarP(
		&FlagNamespace,
		"namespace",
		"n",
		"",
		"Stucked Kubernetes Namespace",
	)
	Cmd.MarkFlagRequired("namespace")
}
