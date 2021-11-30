package secret

import (
	"context"
	"time"

	wait_for_k8s_cmd "github.com/sikalabs/slu/cmd/wait_for/k8s"
	"github.com/sikalabs/slu/utils/k8s"
	"github.com/sikalabs/slu/utils/wait_for_utils"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"

	"github.com/spf13/cobra"
)

var CmdFlagName string
var CmdFlagNamespace string
var CmdFlagTimeout int

var Cmd = &cobra.Command{
	Use:   "pod",
	Short: "Wait for pod",
	Args:  cobra.NoArgs,
	Run: func(c *cobra.Command, args []string) {
		clientset, defaultNamespace, _ := k8s.KubernetesClient()

		namespace := defaultNamespace
		if CmdFlagNamespace != "" {
			namespace = CmdFlagNamespace
		}

		podClient := clientset.CoreV1().Pods(namespace)

		wait_for_utils.WaitFor(
			CmdFlagTimeout, 100*time.Millisecond,
			func() (bool, bool, string, error) {
				pod, err := podClient.Get(context.TODO(), CmdFlagName, metav1.GetOptions{})
				if err != nil {
					return wait_for_utils.WaitForResponseWaiting(err.Error())
				}

				if pod.Status.Phase == "Failed" {
					return wait_for_utils.WaitForResponseFailed("Failed")
				}

				if pod.Status.Phase == "Succeeded" {
					return wait_for_utils.WaitForResponseSucceeded("Succeeded")
				}

				return wait_for_utils.WaitForResponseWaiting("Running")
			},
		)

	},
}

func init() {
	wait_for_k8s_cmd.Cmd.AddCommand(Cmd)
	Cmd.Flags().StringVarP(
		&CmdFlagName,
		"pod-name",
		"p",
		"",
		"Pod Name",
	)
	Cmd.MarkFlagRequired("pod-name")
	Cmd.Flags().StringVarP(
		&CmdFlagNamespace,
		"namespace",
		"n",
		"",
		"Kubernetes Namespace",
	)
	Cmd.Flags().IntVarP(
		&CmdFlagTimeout,
		"timeout",
		"t",
		5*60, // 5 min
		"Timeout",
	)
}
