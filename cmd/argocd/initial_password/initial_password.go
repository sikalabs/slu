package initial_password

import (
	"context"
	"encoding/json"
	"fmt"
	"log"

	argocd_cmd "github.com/sikalabs/slu/cmd/argocd"
	rootcmd "github.com/sikalabs/slu/cmd/root"
	"github.com/sikalabs/slu/utils/k8s"

	"github.com/spf13/cobra"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

var CmdFlagNamespace string

var Cmd = &cobra.Command{
	Use:     "initial-password",
	Short:   "Get ArgoCD Initial Passowrd",
	Aliases: []string{"ip"},
	Args:    cobra.NoArgs,
	Run: func(c *cobra.Command, args []string) {
		clientset, _, _ := k8s.KubernetesClient()

		secretClient := clientset.CoreV1().Secrets(CmdFlagNamespace)

		secret, err := secretClient.Get(context.TODO(), "argocd-initial-admin-secret", metav1.GetOptions{})
		if err != nil {
			log.Fatal(err)
		}

		if rootcmd.RootCmdFlagJson {
			outJson, err := json.Marshal(string(secret.Data["password"]))
			if err != nil {
				panic(err)
			}
			fmt.Println(string(outJson))
		} else {
			fmt.Println(string(secret.Data["password"]))
		}
	},
}

func init() {
	argocd_cmd.Cmd.AddCommand(Cmd)
	Cmd.Flags().StringVarP(
		&CmdFlagNamespace,
		"namespace",
		"n",
		"argocd",
		"ArgoCD Namespace",
	)
}
