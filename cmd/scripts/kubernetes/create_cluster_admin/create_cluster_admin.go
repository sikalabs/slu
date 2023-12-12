package create_cluster_admin

import (
	"context"
	"fmt"
	"log"
	"strconv"
	"time"

	parent_cmd "github.com/sikalabs/slu/cmd/scripts/kubernetes"
	"github.com/sikalabs/slu/utils/k8s"
	"github.com/sikalabs/slu/utils/k8s_scripts"
	"github.com/spf13/cobra"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

var FlagDry bool

var Cmd = &cobra.Command{
	Use:     "create-cluster-admin",
	Short:   "Create Cluster Admin (RBAC)",
	Aliases: []string{"cca"},
	Args:    cobra.NoArgs,
	Run: func(c *cobra.Command, args []string) {
		suffix := strconv.Itoa(int(time.Now().Unix()))
		fmt.Println("cluster-admin-" + suffix)
		k8s_scripts.CreateClusterAdmin(suffix, FlagDry)
	},
}

func init() {
	parent_cmd.Cmd.AddCommand(Cmd)
	Cmd.Flags().BoolVar(
		&FlagDry,
		"dry",
		false,
		"Dry run",
	)
}

func getTokenOrDie(namespace string, serviceAccount string) string {
	clientset, _, _ := k8s.KubernetesClient()

	saClient := clientset.CoreV1().ServiceAccounts(namespace)
	secretClient := clientset.CoreV1().Secrets(namespace)

	sa, err := saClient.Get(context.TODO(), serviceAccount, metav1.GetOptions{})
	if err != nil {
		log.Fatal(err)
	}
	secret, err := secretClient.Get(context.TODO(), sa.Secrets[0].Name, metav1.GetOptions{})
	if err != nil {
		log.Fatal(err)
	}

	return string(secret.Data["token"])
}
