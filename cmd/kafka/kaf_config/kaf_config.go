package kaf_config

import (
	"context"
	"log"
	"os/exec"
	"strings"

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

var Cmd = &cobra.Command{
	Use:   "kaf-config <namespace/kafka-cluster>",
	Short: "Configure kaf for a Kafka cluster using the first working listener",
	Args:  cobra.ExactArgs(1),
	Run: func(c *cobra.Command, args []string) {
		parts := strings.SplitN(args[0], "/", 2)
		if len(parts) != 2 {
			log.Fatal("argument must be in the form namespace/kafka-cluster")
		}
		namespace, kafkaName := parts[0], parts[1]

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

		kafka, err := dynamicClient.Resource(kafkaGVR).Namespace(namespace).Get(
			context.TODO(), kafkaName, metav1.GetOptions{},
		)
		if err != nil {
			log.Fatal(err)
		}

		status, ok := kafka.Object["status"].(map[string]interface{})
		if !ok {
			log.Fatalf("kafka %s/%s has no status", namespace, kafkaName)
		}

		listeners, ok := status["listeners"].([]interface{})
		if !ok || len(listeners) == 0 {
			log.Fatalf("kafka %s/%s has no listeners", namespace, kafkaName)
		}

		for _, l := range listeners {
			listener, ok := l.(map[string]interface{})
			if !ok {
				continue
			}
			bootstrapServers, _ := listener["bootstrapServers"].(string)
			if bootstrapServers == "" {
				continue
			}

			addCmd := exec.Command("kaf", "config", "add-cluster", kafkaName, "-b", bootstrapServers)
			addCmd.Stdout = nil
			addCmd.Stderr = nil
			if err := addCmd.Run(); err != nil {
				log.Fatalf("kaf config add-cluster: %v", err)
			}

			useCmd := exec.Command("kaf", "config", "use-cluster", kafkaName)
			useCmd.Stdout = nil
			useCmd.Stderr = nil
			if err := useCmd.Run(); err != nil {
				log.Fatalf("kaf config use-cluster: %v", err)
			}

			return
		}

		log.Fatalf("no working listener found in kafka %s/%s", namespace, kafkaName)
	},
}

func init() {
	kafka_cmd.Cmd.AddCommand(Cmd)
}
