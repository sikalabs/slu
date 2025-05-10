package run_cron_job

import (
	"os"
	"os/exec"
	"strconv"

	parent_cmd "github.com/sikalabs/slu/cmd/k8s"
	"github.com/sikalabs/slu/utils/time_utils"

	"github.com/spf13/cobra"
)

var FlagNamespace string

var Cmd = &cobra.Command{
	Use:     "run-cron-job <cronjob-name>",
	Short:   "Run a CronJob now",
	Aliases: []string{"rcj"},
	Args:    cobra.ExactArgs(1),
	Run: func(c *cobra.Command, args []string) {
		cronJobName := args[0]

		kubectlArgs := []string{
			"create",
			"job",
			"--from=cronjob/" + cronJobName,
			cronJobName + "-" + strconv.Itoa(time_utils.Unix()),
		}

		if FlagNamespace != "" {
			kubectlArgs = append(kubectlArgs, "--namespace="+FlagNamespace)
		}

		cmd := exec.Command("kubectl", kubectlArgs...)
		cmd.Stdout = os.Stdout
		cmd.Stderr = os.Stderr
		cmd.Stdin = os.Stdin
		cmd.Run()
	},
}

func init() {
	parent_cmd.Cmd.AddCommand(Cmd)
	Cmd.Flags().StringVarP(
		&FlagNamespace,
		"namespace",
		"n",
		"",
		"Kubernetes Namespace",
	)
}
