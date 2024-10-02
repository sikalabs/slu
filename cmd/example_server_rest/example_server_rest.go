package example_server_rest

import (
	"fmt"
	"io"
	"net/http"
	"os"

	"github.com/sikalabs/slu/cmd/root"
	"github.com/spf13/cobra"

	"github.com/rs/zerolog"
	"github.com/rs/zerolog/log"
)

var FlagPort int
var FlagJson bool

func init() {
	root.RootCmd.AddCommand(Cmd)
	Cmd.PersistentFlags().IntVarP(&FlagPort, "port", "p", 8000, "Listen on port")
}

var Cmd = &cobra.Command{
	Use:   "example-server-rest",
	Short: "Run API server which logs method, path and body",
	Args:  cobra.NoArgs,
	Run: func(c *cobra.Command, args []string) {
		Server(FlagPort)
	},
}

func Server(port int) {
	// Setup zerolog logger
	log.Logger = log.Output(zerolog.ConsoleWriter{Out: os.Stdout})

	// Handle any route
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		// Read request body
		body, err := io.ReadAll(r.Body)
		if err != nil {
			log.Error().Err(err).Msg("Error reading request body")
			http.Error(w, "Internal Server Error", http.StatusInternalServerError)
			return
		}
		// Log the request details
		log.Info().
			Str("method", r.Method).
			Str("path", r.URL.Path).
			Str("body", string(body)).
			Msg("Received request")

		// Send a response
		fmt.Fprintln(w, "{}")
	})

	// Start the server
	log.Info().Msgf("Starting server on http:0.0.0.0:%d, see: http://127.0.0.1:%d", port, port)
	if err := http.ListenAndServe(fmt.Sprintf(":%d", port), nil); err != nil {
		log.Fatal().Err(err).Msg("Server failed to start")
	}
}
