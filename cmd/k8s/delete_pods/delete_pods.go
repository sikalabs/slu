package delete_pods

import (
	k8s_cmd "github.com/sikalabs/slu/cmd/k8s"
	"github.com/sikalabs/slu/utils/k8s"

	"github.com/spf13/cobra"
)

var Cmd = &cobra.Command{
	Use:   "delete-pods",
	Short: "Delete \"not ready\" pods in all namespaces",
	Args:  cobra.NoArgs,
	Run: func(c *cobra.Command, args []string) {
		clientset, _, _ := k8s.KubernetesClient()
		k8s.DeleteNoReadyPods(clientset, false)
	},
}

func init() {
	k8s_cmd.Cmd.AddCommand(Cmd)
}
