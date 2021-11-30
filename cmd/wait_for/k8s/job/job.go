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
	Use:   "job",
	Short: "Wait for job",
	Args:  cobra.NoArgs,
	Run: func(c *cobra.Command, args []string) {
		clientset, defaultNamespace, _ := k8s.KubernetesClient()

		namespace := defaultNamespace
		if CmdFlagNamespace != "" {
			namespace = CmdFlagNamespace
		}

		jobsClient := clientset.BatchV1().Jobs(namespace)

		wait_for_utils.WaitFor(
			CmdFlagTimeout, 100*time.Millisecond,
			func() (bool, bool, string, error) {
				job, err := jobsClient.Get(context.TODO(), CmdFlagName, metav1.GetOptions{})
				if err != nil {
					return wait_for_utils.WaitForResponseWaiting(err.Error())
				}

				if job.Status.Failed == 1 {
					return wait_for_utils.WaitForResponseFailed("Failed")
				}

				if job.Status.Succeeded == 1 {
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
		"job-name",
		"j",
		"",
		"Job Name",
	)
	Cmd.MarkFlagRequired("job-name")
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
