package install_kafdrop_openshift

import (
	parent_cmd "github.com/sikalabs/slu/cmd/scripts/kubernetes"
	"github.com/sikalabs/slu/utils/k8s_scripts"
	"github.com/spf13/cobra"
)

var FlagDry bool
var FlagName string
var FlagNamespace string
var FlagHost string
var FlagKafkaBootstrap string

var Cmd = &cobra.Command{
	Use:     "install-kafdrop-openshift",
	Short:   "Install Kafdrop on OpenShift",
	Aliases: []string{"ikdo"},
	Args:    cobra.NoArgs,
	Run: func(c *cobra.Command, args []string) {
		k8s_scripts.InstallKafdropOpenshift(FlagName, FlagNamespace, FlagHost, FlagKafkaBootstrap, FlagDry)
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
	Cmd.Flags().StringVar(
		&FlagName,
		"name",
		"",
		"Kafdrop release name",
	)
	Cmd.Flags().StringVarP(
		&FlagNamespace,
		"namespace",
		"n",
		"kafdrop",
		"Kubernetes Namespace",
	)
	Cmd.Flags().StringVar(
		&FlagHost,
		"host",
		"",
		"Kafdrop host",
	)
	Cmd.Flags().StringVar(
		&FlagKafkaBootstrap,
		"kafka-bootstrap",
		"",
		"Kafka bootstrap server",
	)
	Cmd.MarkFlagRequired("name")
	Cmd.MarkFlagRequired("host")
	Cmd.MarkFlagRequired("kafka-bootstrap")
}
