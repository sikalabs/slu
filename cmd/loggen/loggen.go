package loggen

import (
	"fmt"
	"log"
	"math/rand"
	"net/url"
	"os"
	"time"

	"github.com/grafana/loki-client-go/loki"
	"github.com/grafana/loki-client-go/pkg/urlutil"
	"github.com/prometheus/common/model"
	"github.com/rs/zerolog"
	"github.com/sikalabs/slu/cmd/root"
	"github.com/sikalabs/slu/internal/error_utils"
	"github.com/spf13/cobra"
)

var FlagJson bool
var FlagLogFile string
var FlagLokiURL string
var FlagLogPrefix string
var FlagSleepTime int
var FlagNoError bool
var FlagNoWarn bool
var FlagNoInfo bool
var FlagNoDebug bool
var FlagLokiLabelInstance string
var FlagLimit int

var Cmd = &cobra.Command{
	Use:   "loggen",
	Short: "Log Generator",
	Args:  cobra.NoArgs,
	Run: func(c *cobra.Command, args []string) {
		u, err := url.Parse(FlagLokiURL)
		error_utils.HandleError(err, "Failed to parse Loki URL")
		cfg := loki.Config{
			URL: urlutil.URLValue{
				URL: u,
			},
			BatchWait: 5 * time.Second,
			BatchSize: 1024 * 1024,
			Timeout:   2 * time.Second,
		}

		client, err := loki.New(cfg)
		error_utils.HandleError(err, "Failed to create Loki client")
		defer client.Stop()

		var i int = 0
		if FlagJson {
			var logger zerolog.Logger
			if FlagLogFile != "" {
				f, err := os.OpenFile(FlagLogFile,
					os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
				if err != nil {
					panic(err)
				}
				defer f.Close()
				logger = zerolog.New(f).With().Timestamp().Logger()
				logger.Info().Str("prefix", FlagLogPrefix).Msg("Logging into file " + FlagLogFile)
			} else {
				logger = zerolog.New(os.Stdout).With().Timestamp().Logger()
				logger.Info().Str("prefix", FlagLogPrefix).Msg("Logging into STDOUT")
			}

			if FlagNoError && FlagNoWarn && FlagNoInfo && FlagNoDebug {
				logger.Error().Str("prefix", FlagLogPrefix).Msg("ERROR No logging output enabled.")
				os.Exit(1)
			}

			for {
				if FlagLimit > 0 && i >= FlagLimit {
					logger.Info().Str("prefix", FlagLogPrefix).Int("i", i).Msg("Reached limit, exiting.")
					os.Exit(0)
				}

				time.Sleep(time.Duration(FlagSleepTime) * time.Millisecond)

				randomNumber := rand.Intn(100)
				if randomNumber > 90 && !FlagNoError {
					logger.Error().Str("prefix", FlagLogPrefix).Int("i", i).Msg("ERROR An error is usually an exception that has been caught and not handled.")
					if FlagLokiURL != "" {
						client.Handle(model.LabelSet{
							"prefix":   model.LabelValue(FlagLogPrefix),
							"level":    "error",
							"instance": model.LabelValue(FlagLokiLabelInstance),
						}, time.Now(), "An error is usually an exception that has been caught and not handled. "+fmt.Sprintf("i=%d", i))
					}
					i++
					continue
				}
				if randomNumber > 70 && !FlagNoWarn {
					logger.Warn().Str("prefix", FlagLogPrefix).Int("i", i).Msg("WARN A warning that should be ignored is usually at this level and should be actionable.")
					if FlagLokiURL != "" {
						client.Handle(model.LabelSet{
							"prefix":   model.LabelValue(FlagLogPrefix),
							"level":    "warn",
							"instance": model.LabelValue(FlagLokiLabelInstance),
						}, time.Now(), "WARN A warning that should be ignored is usually at this level and should be actionable. "+fmt.Sprintf("i=%d", i))
					}
					i++
					continue
				}
				if randomNumber > 30 && !FlagNoInfo {
					logger.Info().Str("prefix", FlagLogPrefix).Int("i", i).Msg("INFO This is less important than debug log and is often used to provide context in the current task.")
					if FlagLokiURL != "" {
						client.Handle(model.LabelSet{
							"prefix":   model.LabelValue(FlagLogPrefix),
							"level":    "info",
							"instance": model.LabelValue(FlagLokiLabelInstance),
						}, time.Now(), "INFO This is less important than debug log and is often used to provide context in the current task. "+fmt.Sprintf("i=%d", i))
					}
					i++
					continue
				}
				if !FlagNoDebug {
					logger.Debug().Str("prefix", FlagLogPrefix).Int("i", i).Msg("DEBUG This is a debug log that shows a log that can be ignored.")
					if FlagLokiURL != "" {
						client.Handle(model.LabelSet{
							"prefix":   model.LabelValue(FlagLogPrefix),
							"level":    "debug",
							"instance": model.LabelValue(FlagLokiLabelInstance),
						}, time.Now(), "DEBUG This is a debug log that shows a log that can be ignored. "+fmt.Sprintf("i=%d", i))
					}
					i++
					continue
				}
			}
		} else {
			var logger *log.Logger
			if FlagLogFile != "" {
				f, err := os.OpenFile(FlagLogFile,
					os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
				if err != nil {
					panic(err)
				}
				defer f.Close()
				logger = log.New(f, FlagLogPrefix+" ", log.LstdFlags)
				logger.Println("Logging into file " + FlagLogFile)

			} else {
				logger = log.New(os.Stdout, FlagLogPrefix+" ", log.LstdFlags)
				logger.Println("Logging into STDOUT")
			}

			if FlagNoError && FlagNoWarn && FlagNoInfo && FlagNoDebug {
				logger.Println("ERROR No logging output enabled.")
				os.Exit(1)
			}

			for {
				if FlagLimit > 0 && i >= FlagLimit {
					// Log the limit reached message
					logger.Printf("INFO Reached limit, exiting. (i=%d)\n", i)
					os.Exit(0)
				}

				time.Sleep(time.Duration(FlagSleepTime) * time.Millisecond)

				randomNumber := rand.Intn(100)
				if randomNumber > 90 && !FlagNoError {
					logger.Printf("ERROR An error is usually an exception that has been caught and not handled. (i=%d)\n", i)
					if FlagLokiURL != "" {
						client.Handle(model.LabelSet{
							"prefix":   model.LabelValue(FlagLogPrefix),
							"level":    "error",
							"instance": model.LabelValue(FlagLokiLabelInstance),
						}, time.Now(), "An error is usually an exception that has been caught and not handled. "+fmt.Sprintf("i=%d", i))
					}
					i++
					continue
				}
				if randomNumber > 70 && !FlagNoWarn {
					logger.Printf("WARN A warning that should be ignored is usually at this level and should be actionable. (i=%d)\n", i)
					if FlagLokiURL != "" {
						client.Handle(model.LabelSet{
							"prefix":   model.LabelValue(FlagLogPrefix),
							"level":    "warn",
							"instance": model.LabelValue(FlagLokiLabelInstance),
						}, time.Now(), "WARN A warning that should be ignored is usually at this level and should be actionable. "+fmt.Sprintf("i=%d", i))
					}
					i++
					continue
				}
				if randomNumber > 30 && !FlagNoInfo {
					logger.Printf("INFO This is less important than debug log and is often used to provide context in the current task (i=%d)\n", i)
					if FlagLokiURL != "" {
						client.Handle(model.LabelSet{
							"prefix":   model.LabelValue(FlagLogPrefix),
							"level":    "info",
							"instance": model.LabelValue(FlagLokiLabelInstance),
						}, time.Now(), "INFO This is less important than debug log and is often used to provide context in the current task. "+fmt.Sprintf("i=%d", i))
					}
					i++
					continue
				}
				if !FlagNoDebug {
					logger.Printf("DEBUG This is a debug log that shows a log that can be ignored. (i=%d)\n", i)
					if FlagLokiURL != "" {
						client.Handle(model.LabelSet{
							"prefix":   model.LabelValue(FlagLogPrefix),
							"level":    "debug",
							"instance": model.LabelValue(FlagLokiLabelInstance),
						}, time.Now(), "DEBUG This is a debug log that shows a log that can be ignored. "+fmt.Sprintf("i=%d", i))
					}
					i++
					continue
				}
			}
		}
	},
}

func init() {
	root.RootCmd.AddCommand(Cmd)
	Cmd.Flags().StringVarP(
		&FlagLogFile,
		"log-file",
		"f",
		"",
		"Output log file (default: STDOUT)",
	)
	Cmd.Flags().StringVar(
		&FlagLokiURL,
		"loki-url",
		"",
		"Log also to Loki, use full URL (e.g. http://127.0.0.1:3100/loki/api/v1/push)",
	)
	Cmd.Flags().StringVarP(
		&FlagLogPrefix,
		"log-prefix",
		"p",
		"loggen",
		"Log prefix",
	)
	Cmd.Flags().IntVarP(
		&FlagSleepTime,
		"sleep-time",
		"s",
		1000,
		"Sleep time (in ms)	",
	)
	Cmd.Flags().BoolVarP(
		&FlagNoError,
		"no-error",
		"e",
		false,
		"No errors",
	)
	Cmd.Flags().BoolVarP(
		&FlagNoWarn,
		"no-warn",
		"w",
		false,
		"No warnings",
	)
	Cmd.Flags().BoolVarP(
		&FlagNoInfo,
		"no-info",
		"i",
		false,
		"No infos",
	)
	Cmd.Flags().BoolVarP(
		&FlagNoDebug,
		"no-debug",
		"d",
		false,
		"No debugs",
	)
	Cmd.PersistentFlags().BoolVar(
		&FlagJson,
		"json",
		false,
		"Format output to JSON",
	)
	Cmd.Flags().StringVar(
		&FlagLokiLabelInstance,
		"loki-label-instance",
		"0",
		"Loki label instance",
	)
	Cmd.Flags().IntVar(
		&FlagLimit,
		"limit",
		0,
		"Limit number of logs to generate (default: no limit)",
	)
}
