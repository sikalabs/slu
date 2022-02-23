package bad_pods

import (
	"context"
	"fmt"
	"log"

	k8s_get_cmd "github.com/sikalabs/slu/cmd/k8s/get"
	"github.com/sikalabs/slu/utils/k8s"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"

	"github.com/spf13/cobra"
)

var FlagNamespace string
var FlagAllNamespaces bool

var Cmd = &cobra.Command{
	Use:     "bad-pods",
	Short:   "Get bad pods (all pods without them with Phase in: Running, Succeeded",
	Aliases: []string{"bp"},
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
		for _, pod := range pods.Items {
			if pod.Status.Phase == "Running" {
				continue
			}
			if pod.Status.Phase == "Succeeded" {
				continue
			}
			fmt.Println(pod.Namespace, pod.Name, pod.Status.Phase)
		}
	},
}

func init() {
	k8s_get_cmd.Cmd.AddCommand(Cmd)
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
