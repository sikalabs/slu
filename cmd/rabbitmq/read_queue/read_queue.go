package read_queue

import (
	"log"
	"os"
	"os/signal"
	"syscall"

	parent_cmd "github.com/sikalabs/slu/cmd/rabbitmq"
	rabbitmqUtisl "github.com/sikalabs/slu/utils/rabbitmq"
	"github.com/spf13/cobra"
)

var (
	FlagHost        string
	FlagPort        int
	FlagUser        string
	FlagPassword    string
	FlagVirtualHost string
	FlagQueue       string
	FlagSsl         bool
	FlagSslCert     string
	FlagSslKey      string
)

var Cmd = &cobra.Command{
	Use:     "read-queue",
	Aliases: []string{"rq"},
	Short:   "Read messages from RabbitMQ queue",
	Long:    "Read messages from RabbitMQ queue. It will connect to the server and read messages from the specified queue",
	Example: "slu rabbitmq read-queue -host localhost --port 5672 --user guest --password guest --vhost / --queue my_queue",
	Run: func(cmd *cobra.Command, args []string) {
		readQueue(FlagHost, FlagPort, FlagUser, FlagPassword, FlagVirtualHost, FlagQueue, FlagSsl, FlagSslCert, FlagSslKey)
	},
}

func init() {
	parent_cmd.Cmd.AddCommand(Cmd)
	Cmd.Flags().StringVarP(&FlagHost, "host", "H", "localhost", "RabbitMQ server host")
	Cmd.MarkFlagRequired("host")
	Cmd.Flags().IntVarP(&FlagPort, "port", "P", 5672, "RabbitMQ server port")
	Cmd.Flags().StringVarP(&FlagUser, "user", "u", "guest", "RabbitMQ user")
	Cmd.Flags().StringVarP(&FlagPassword, "password", "p", "guest", "RabbitMQ password")
	Cmd.Flags().StringVarP(&FlagVirtualHost, "vhost", "v", "/", "RabbitMQ virtual host")
	Cmd.Flags().StringVarP(&FlagQueue, "queue", "q", "", "RabbitMQ queue name to read messages from")
	Cmd.MarkFlagRequired("queue")
	Cmd.Flags().BoolVarP(&FlagSsl, "ssl", "s", false, "Use SSL for RabbitMQ connection")
	Cmd.Flags().StringVarP(&FlagSslCert, "ssl-cert", "c", "", "Path to SSL certificate file")
	Cmd.Flags().StringVarP(&FlagSslKey, "ssl-key", "k", "", "Path to SSL key file")
}

func readQueue(host string, port int, user, password, virtualHost, queue string, ssl bool, sslCert, sslKey string) {
	con, ch, err := rabbitmqUtisl.ConnectToRabbitMQ(ssl, user, password, host, port, virtualHost, sslCert, sslKey)
	if err != nil {
		log.Fatalf("Connection to RabbitMQ failed: %s", err.Error())
	}
	defer con.Close()
	defer ch.Close()

	msgs, err := ch.Consume(queue, "", true, false, false, false, nil)
	if err != nil {
		log.Fatalf("Failed to register consumer: %s", err.Error())
	}

	sigs := make(chan os.Signal, 1)
	signal.Notify(sigs, syscall.SIGINT, syscall.SIGTERM)

	go func() {
		<-sigs
		log.Println("Shutting down...")
		con.Close()
		ch.Close()
		os.Exit(0)
	}()

	for msg := range msgs {
		log.Printf("Message from the queue: %s", string(msg.Body))
	}
}
