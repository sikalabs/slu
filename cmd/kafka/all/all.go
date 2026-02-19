package all

import (
	"context"
	"fmt"
	"log"

	kafka_cmd "github.com/sikalabs/slu/cmd/kafka"
	"github.com/spf13/cobra"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/runtime/schema"
	"k8s.io/client-go/dynamic"
	"k8s.io/client-go/tools/clientcmd"
)

var kafkaGVR = schema.GroupVersionResource{
	Group:    "kafka.strimzi.io",
	Version:  "v1beta2",
	Resource: "kafkas",
}

var Cmd = &cobra.Command{
	Use:   "all",
	Short: "List all Kafka clusters and their bootstrap servers",
	Args:  cobra.NoArgs,
	Run: func(c *cobra.Command, args []string) {
		kubeConfig := clientcmd.NewNonInteractiveDeferredLoadingClientConfig(
			clientcmd.NewDefaultClientConfigLoadingRules(),
			&clientcmd.ConfigOverrides{},
		)
		restconfig, err := kubeConfig.ClientConfig()
		if err != nil {
			log.Fatal(err)
		}
		dynamicClient, err := dynamic.NewForConfig(restconfig)
		if err != nil {
			log.Fatal(err)
		}

		kafkaList, err := dynamicClient.Resource(kafkaGVR).Namespace("").List(
			context.TODO(), metav1.ListOptions{},
		)
		if err != nil {
			log.Fatal(err)
		}

		for _, kafka := range kafkaList.Items {
			namespace := kafka.GetNamespace()
			name := kafka.GetName()

			status, ok := kafka.Object["status"].(map[string]interface{})
			if !ok {
				fmt.Printf("%s/%s (no status)\n", namespace, name)
				continue
			}

			listeners, ok := status["listeners"].([]interface{})
			if !ok || len(listeners) == 0 {
				fmt.Printf("%s/%s (no listeners)\n", namespace, name)
				continue
			}

			fmt.Printf("%s/%s\n", namespace, name)
			for _, l := range listeners {
				listener, ok := l.(map[string]interface{})
				if !ok {
					continue
				}
				listenerName, _ := listener["name"].(string)
				bootstrapServers, _ := listener["bootstrapServers"].(string)
				fmt.Printf("  %s: %s\n", listenerName, bootstrapServers)
			}
		}
	},
}

func init() {
	kafka_cmd.Cmd.AddCommand(Cmd)
}
