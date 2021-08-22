package parse_k8s_secret

import (
	"context"
	"log"

	tls_cmd "github.com/sikalabs/slu/cmd/tls"
	"github.com/sikalabs/slu/utils/k8s"
	"github.com/sikalabs/slu/utils/tls_utils"

	"github.com/spf13/cobra"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

var CmdFlagSecretName string
var CmdFlagNamespace string

var Cmd = &cobra.Command{
	Use:     "parse-k8s-secret",
	Short:   "Parse TLS Certificate from Kubernetes Secret",
	Aliases: []string{"parse-k8s", "parse-secret", "parse-sec"},
	Args:    cobra.NoArgs,
	Run: func(c *cobra.Command, args []string) {
		clientset, defaultNamespace, _ := k8s.KubernetesClient()

		namespace := defaultNamespace
		if CmdFlagNamespace != "" {
			namespace = CmdFlagNamespace
		}

		secretClient := clientset.CoreV1().Secrets(namespace)

		secret, err := secretClient.Get(
			context.TODO(),
			CmdFlagSecretName,
			metav1.GetOptions{},
		)
		if err != nil {
			log.Fatal(err)
		}

		tls_utils.PrintCertificateFromBytes(
			secret.Data["tls.crt"],
			secret.Data["tls.key"],
		)
	},
}

func init() {
	tls_cmd.Cmd.AddCommand(Cmd)
	Cmd.Flags().StringVarP(
		&CmdFlagSecretName,
		"secret-name",
		"s",
		"",
		"Secret Name",
	)
	Cmd.Flags().StringVarP(
		&CmdFlagNamespace,
		"namespace",
		"n",
		"",
		"Kubernetes Namespace",
	)
}
