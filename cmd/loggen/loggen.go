package loggen

import (
	"log"
	"math/rand"
	"os"
	"time"

	"github.com/rs/zerolog"
	"github.com/sikalabs/slu/cmd/root"
	"github.com/spf13/cobra"
)

var FlagLogFile string
var FlagLogPrefix string
var FlagSleepTime int
var FlagNoError bool
var FlagNoWarn bool
var FlagNoInfo bool
var FlagNoDebug bool

var Cmd = &cobra.Command{
	Use:   "loggen",
	Short: "Log Generator",
	Args:  cobra.NoArgs,
	Run: func(c *cobra.Command, args []string) {
		if root.RootCmdFlagJson {
			var logger zerolog.Logger
			if FlagLogFile != "" {
				f, err := os.OpenFile(FlagLogFile,
					os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
				if err != nil {
					panic(err)
				}
				defer f.Close()
				logger = zerolog.New(f).With().Timestamp().Logger()
				logger.Info().Msg("Logging into file " + FlagLogFile)

			} else {
				logger = zerolog.New(os.Stdout).With().Timestamp().Logger()
				logger.Info().Msg("Logging into STDOUT")
			}

			if FlagNoError && FlagNoWarn && FlagNoInfo && FlagNoDebug {
				logger.Error().Msg("ERROR No logging output enabled.")
				os.Exit(1)
			}

			for {
				time.Sleep(time.Duration(FlagSleepTime) * time.Millisecond)

				randomNumber := rand.Intn(100)
				if randomNumber > 90 && !FlagNoError {
					logger.Error().Msg("An error is usually an exception that has been caught and not handled.")
					continue
				}
				if randomNumber > 70 && !FlagNoWarn {
					logger.Warn().Msg("WARN A warning that should be ignored is usually at this level and should be actionable.")
					continue
				}
				if randomNumber > 30 && !FlagNoInfo {
					logger.Info().Msg("INFO This is less important than debug log and is often used to provide context in the current task.")
					continue
				}
				if !FlagNoDebug {
					logger.Debug().Msg("DEBUG This is a debug log that shows a log that can be ignored.")
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
				time.Sleep(time.Duration(FlagSleepTime) * time.Millisecond)

				randomNumber := rand.Intn(100)
				if randomNumber > 90 && !FlagNoError {
					logger.Println("ERROR An error is usually an exception that has been caught and not handled.")
					continue
				}
				if randomNumber > 70 && !FlagNoWarn {
					logger.Println("WARN A warning that should be ignored is usually at this level and should be actionable.")
					continue
				}
				if randomNumber > 30 && !FlagNoInfo {
					logger.Println("INFO This is less important than debug log and is often used to provide context in the current task.")
					continue
				}
				if !FlagNoDebug {
					logger.Println("DEBUG This is a debug log that shows a log that can be ignored.")
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
}
