package secret

import (
	"context"
	"fmt"
	"os"
	"time"

	wait_for_k8s_cmd "github.com/sikalabs/slu/cmd/wait_for/k8s"
	"github.com/sikalabs/slu/utils/k8s"
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

		started := time.Now()

		latestPhase := ""

		for {
			pod, err := podClient.Get(context.TODO(), CmdFlagName, metav1.GetOptions{})
			if err != nil {
				fmt.Println(err)
				time.Sleep(100 * time.Millisecond)
				continue
			}

			phase := string(pod.Status.Phase)

			if latestPhase != phase {
				fmt.Println("pods \"" + CmdFlagName + "\" phase " + phase)
			}
			latestPhase = phase

			if phase == "Failed" {
				os.Exit(1)
			}

			if phase == "Succeeded" {
				os.Exit(0)
			}

			if time.Since(started) > time.Duration(CmdFlagTimeout*int(time.Second)) {
				os.Exit(1)
			}

			time.Sleep(100 * time.Millisecond)
		}

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
