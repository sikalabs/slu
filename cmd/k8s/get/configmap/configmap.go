package configmap

import (
	"context"
	"encoding/json"
	"fmt"
	"log"

	k8s_get_cmd "github.com/sikalabs/slu/cmd/k8s/get"
	"github.com/sikalabs/slu/utils/k8s"

	"github.com/spf13/cobra"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

var FlagJson bool
var CmdFlagName string
var CmdFlagNamespace string
var CmdFlagKey string

var Cmd = &cobra.Command{
	Use:     "configmap",
	Short:   "Get data from ConfigMap",
	Aliases: []string{"cm"},
	Args:    cobra.NoArgs,
	Run: func(c *cobra.Command, args []string) {
		clientset, defaultNamespace, _ := k8s.KubernetesClient()

		namespace := defaultNamespace
		if CmdFlagNamespace != "" {
			namespace = CmdFlagNamespace
		}

		configmapClient := clientset.CoreV1().ConfigMaps(namespace)

		configmap, err := configmapClient.Get(context.TODO(), CmdFlagName, metav1.GetOptions{})
		if err != nil {
			log.Fatal(err)
		}

		if CmdFlagKey != "" {
			if FlagJson {
				outJson, err := json.Marshal(string(configmap.Data[CmdFlagKey]))
				if err != nil {
					panic(err)
				}
				fmt.Println(string(outJson))
			} else {
				fmt.Println(string(configmap.Data[CmdFlagKey]))
			}
		} else {
			if FlagJson {
				out := make(map[string]string)
				for key, val := range configmap.Data {
					out[key] = string(val)
				}
				outJson, err := json.Marshal(out)
				if err != nil {
					panic(err)
				}
				fmt.Println(string(outJson))
			} else {
				for key, val := range configmap.Data {
					fmt.Printf("KEY:   %s\nVALUE: %s\n---\n", key, val)
				}
			}
		}
	},
}

func init() {
	k8s_get_cmd.Cmd.AddCommand(Cmd)
	Cmd.Flags().StringVarP(
		&CmdFlagName,
		"configmap",
		"c",
		"default",
		"ConfigMap Name",
	)
	Cmd.MarkFlagRequired("configmap")
	Cmd.Flags().StringVarP(
		&CmdFlagNamespace,
		"namespace",
		"n",
		"",
		"Kubernetes Namespace",
	)
	Cmd.Flags().StringVarP(
		&CmdFlagKey,
		"key",
		"k",
		"",
		"Get only specific key from data",
	)
	Cmd.PersistentFlags().BoolVar(
		&FlagJson,
		"json",
		false,
		"Format output to JSON",
	)
}
