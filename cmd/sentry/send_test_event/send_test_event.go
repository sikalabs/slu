package send_test_event

import (
	"log"
	"time"

	"github.com/getsentry/sentry-go"
	parent_cmd "github.com/sikalabs/slu/cmd/sentry"
	"github.com/spf13/cobra"
)

var Cmd = &cobra.Command{
	Use:   "send-test-event <sentry_dsn>",
	Short: "Send test event to Sentry",
	Args:  cobra.ExactArgs(1),
	Run: func(c *cobra.Command, args []string) {
		send(args[0])
	},
}

func init() {
	parent_cmd.Cmd.AddCommand(Cmd)
}

func send(dsn string) {
	err := sentry.Init(sentry.ClientOptions{
		Dsn:   dsn,
		Debug: true,
	})
	if err != nil {
		log.Fatalf("sentry.Init: %s", err)
	}
	defer sentry.Flush(2 * time.Second)

	sentry.CaptureMessage("It works!")
}
