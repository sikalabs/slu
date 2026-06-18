package list_containers

import (
	"context"
	"log"
	"os"

	"github.com/olekukonko/tablewriter"
	k8s_cmd "github.com/sikalabs/slu/cmd/k8s"
	"github.com/sikalabs/slu/utils/k8s"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"

	"github.com/spf13/cobra"
)

var FlagNamespace string
var FlagAllNamespaces bool

var Cmd = &cobra.Command{
	Use:     "list-containers",
	Short:   "List all containers with their images",
	Aliases: []string{"lc"},
	Args:    cobra.NoArgs,
	Run: func(c *cobra.Command, args []string) {
		clientset, defaultNamespace, _ := k8s.KubernetesClient()

		namespace := defaultNamespace
		if FlagNamespace != "" {
			namespace = FlagNamespace
		}

		if FlagAllNamespaces {
			namespace = ""
		}

		podClient := clientset.CoreV1().Pods(namespace)
		pods, err := podClient.List(context.TODO(), metav1.ListOptions{})
		if err != nil {
			log.Fatal(err)
		}

		table := tablewriter.NewWriter(os.Stdout)
		table.Header([]string{
			"Namespace",
			"Pod",
			"Container Image",
		})
		for _, pod := range pods.Items {
			for _, container := range pod.Spec.Containers {
				table.Append([]string{pod.Namespace, pod.Name, container.Image})
			}
		}
		table.Render()
	},
}

func init() {
	k8s_cmd.Cmd.AddCommand(Cmd)
	Cmd.Flags().StringVarP(
		&FlagNamespace,
		"namespace",
		"n",
		"",
		"Kubernetes Namespace",
	)
	Cmd.Flags().BoolVarP(
		&FlagAllNamespaces,
		"all-namespaces",
		"A",
		false,
		"All Namespaces",
	)
}
