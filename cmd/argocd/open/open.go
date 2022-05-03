package open

import (
	"context"
	"fmt"
	"log"

	"github.com/pkg/browser"
	argocd_cmd "github.com/sikalabs/slu/cmd/argocd"
	"github.com/sikalabs/slu/utils/k8s"

	"github.com/spf13/cobra"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

var FlagNamespace string

var Cmd = &cobra.Command{
	Use:   "open",
	Short: "Open ArgoCD in browser",
	Args:  cobra.NoArgs,
	Run: func(c *cobra.Command, args []string) {
		clientset, _, _ := k8s.KubernetesClient()

		ingressClient := clientset.NetworkingV1().Ingresses(FlagNamespace)

		ingress, err := ingressClient.Get(context.TODO(), "argocd-server", metav1.GetOptions{})
		if err != nil {
			log.Fatal(err)
		}

		rule := ingress.Spec.Rules[0]
		url := "http://" + rule.Host + rule.HTTP.Paths[0].Path
		fmt.Println(url)

		browser.OpenURL(url)
	},
}

func init() {
	argocd_cmd.Cmd.AddCommand(Cmd)
	Cmd.Flags().StringVarP(
		&FlagNamespace,
		"namespace",
		"n",
		"argocd",
		"ArgoCD Namespace",
	)
}
