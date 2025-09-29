package put_message

import (
	"github.com/rabbitmq/amqp091-go"
	parent_cmd "github.com/sikalabs/slu/cmd/rabbitmq"
	rabbitmqUtisl "github.com/sikalabs/slu/utils/rabbitmq"
	"github.com/spf13/cobra"
	"log"
)

var (
	FlagHost        string
	FlagPort        int
	FlagUser        string
	FlagPassword    string
	FlagVirtualHost string
	FlagExchange    string
	FlagRoutingKey  string
	FlagSsl         bool
	FlagSslCert     string
	FlagSslKey      string
	FlagMessage     string
)

var Cmd = &cobra.Command{
	Use:     "put-message",
	Aliases: []string{"pm"},
	Short:   "Put a message to RabbitMQ server",
	Long:    "Put a message to RabbitMQ server. It will connect to the server and send a message to the specified exchange and routing key",
	Example: "slu rabbitmq put-message --host localhost --port 5672 --user guest --password guest --vhost / --exchange my_exchange --routing-key my_routing_key --message '{\"message\": \"Hello RabbitMQ! This is default test message\"}'",
	Run: func(cmd *cobra.Command, args []string) {
		putMessage(FlagHost, FlagPort, FlagUser, FlagPassword, FlagVirtualHost, FlagExchange, FlagRoutingKey, FlagSsl, FlagSslCert, FlagSslKey)
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
	Cmd.Flags().StringVarP(&FlagExchange, "exchange", "e", "", "RabbitMQ exchange name to send the message to")
	Cmd.MarkFlagRequired("exchange")
	Cmd.Flags().StringVarP(&FlagRoutingKey, "routing-key", "r", "", "RabbitMQ routing key")
	Cmd.Flags().BoolVarP(&FlagSsl, "ssl", "s", false, "Use SSL for RabbitMQ connection")
	Cmd.Flags().StringVarP(&FlagSslCert, "ssl-cert", "c", "", "Path to SSL certificate file")
	Cmd.Flags().StringVarP(&FlagSslKey, "ssl-key", "k", "", "Path to SSL key file")
	Cmd.Flags().StringVarP(&FlagMessage, "message", "m", "", "Message to send to RabbitMQ")
}

func putMessage(host string, port int, user, password, virtualHost, exchange, routingKey string, ssl bool, sslCert, sslKey string) {
	con, ch, err := rabbitmqUtisl.ConnectToRabbitMQ(ssl, user, password, host, port, virtualHost, sslCert, sslKey)
	if err != nil {
		log.Fatalf("Connection to RabbitMQ failed: %s", err.Error())
	}

	var messageUsed string
	if FlagMessage != "" {
		messageUsed = FlagMessage
	} else {
		messageUsed = `{"message": "Hello RabbitMQ! This is default test message"}`
	}

	defer con.Close()
	defer ch.Close()
	if err := ch.Publish(exchange,
		routingKey,
		false,
		false,
		amqp091.Publishing{
			ContentType: "application/json",
			Body:        []byte(messageUsed),
		},
	); err != nil {
		log.Fatalf("Failed to put message: %s", err.Error())
	}
	log.Printf("Message sent to exchange '%s' with routing key '%s'", exchange, routingKey)
}
