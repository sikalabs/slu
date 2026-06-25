package signpost

import (
	"context"
	"fmt"
	"log"

	"github.com/pkg/browser"
	open_cmd "github.com/sikalabs/slu/cmd/scripts/open"
	"github.com/sikalabs/slu/utils/k8s"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"

	"github.com/spf13/cobra"
)

var Cmd = &cobra.Command{
	Use:   "signpost",
	Short: "Open Signpost in browser",
	Args:  cobra.NoArgs,
	Run: func(c *cobra.Command, args []string) {
		clientset, _, err := k8s.KubernetesClient()
		if err != nil {
			log.Fatal(err)
		}

		ingress, err := clientset.NetworkingV1().Ingresses("signpost").Get(
			context.TODO(), "signpost", metav1.GetOptions{},
		)
		if err != nil {
			log.Fatal(err)
		}

		host := ingress.Spec.Rules[0].Host
		url := "https://" + host
		fmt.Println(url)
		browser.OpenURL(url)
	},
}

func init() {
	open_cmd.Cmd.AddCommand(Cmd)
}
