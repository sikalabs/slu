package bootstrap

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
	Version:  "v1",
	Resource: "kafkas",
}

var flagNamespace string
var flagKafka string
var flagListener string

var Cmd = &cobra.Command{
	Use:   "bootstrap",
	Short: "Get bootstrap server for a Kafka cluster listener",
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

		kafka, err := dynamicClient.Resource(kafkaGVR).Namespace(flagNamespace).Get(
			context.TODO(), flagKafka, metav1.GetOptions{},
		)
		if err != nil {
			log.Fatal(err)
		}

		status, ok := kafka.Object["status"].(map[string]interface{})
		if !ok {
			log.Fatalf("kafka %s/%s has no status", flagNamespace, flagKafka)
		}

		listeners, ok := status["listeners"].([]interface{})
		if !ok || len(listeners) == 0 {
			log.Fatalf("kafka %s/%s has no listeners", flagNamespace, flagKafka)
		}

		for _, l := range listeners {
			listener, ok := l.(map[string]interface{})
			if !ok {
				continue
			}
			name, _ := listener["name"].(string)
			if name != flagListener {
				continue
			}
			bootstrapServers, _ := listener["bootstrapServers"].(string)
			fmt.Println(bootstrapServers)
			return
		}

		log.Fatalf("listener %q not found in kafka %s/%s", flagListener, flagNamespace, flagKafka)
	},
}

func init() {
	Cmd.Flags().StringVarP(&flagNamespace, "namespace", "n", "kafka", "Kubernetes namespace")
	Cmd.Flags().StringVarP(&flagKafka, "kafka", "k", "", "Kafka cluster name")
	Cmd.Flags().StringVarP(&flagListener, "listener", "l", "", "Listener name")
	_ = Cmd.MarkFlagRequired("kafka")
	_ = Cmd.MarkFlagRequired("listener")

	kafka_cmd.Cmd.AddCommand(Cmd)
}
