package run_kafka

import (
	parent_cmd "github.com/sikalabs/slu/cmd/scripts/docker"
	"github.com/sikalabs/slu/utils/exec_utils"
	"github.com/spf13/cobra"
)

var Cmd = &cobra.Command{
	Use:   "run-kafka",
	Short: "Run Kafka in Docker (apache/kafka:latest on port 9092)",
	Args:  cobra.NoArgs,
	Run: func(c *cobra.Command, args []string) {
		exec_utils.ExecOut(
			"docker",
			"run",
			"-d",
			"--name", "kafka",
			"-p", "9092:9092",
			"-e", "KAFKA_NODE_ID=1",
			"-e", "KAFKA_PROCESS_ROLES=broker,controller",
			"-e", "KAFKA_LISTENERS=PLAINTEXT://0.0.0.0:9092,CONTROLLER://0.0.0.0:9093",
			"-e", "KAFKA_ADVERTISED_LISTENERS=PLAINTEXT://localhost:9092",
			"-e", "KAFKA_CONTROLLER_LISTENER_NAMES=CONTROLLER",
			"-e", "KAFKA_CONTROLLER_QUORUM_VOTERS=1@localhost:9093",
			"-e", "KAFKA_OFFSETS_TOPIC_REPLICATION_FACTOR=1",
			"apache/kafka:latest",
		)
	},
}

func init() {
	parent_cmd.Cmd.AddCommand(Cmd)
}
