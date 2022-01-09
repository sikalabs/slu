package send

import (
	"log"
	"strconv"

	mail_cmd "github.com/sikalabs/slu/cmd/mail"
	"github.com/sikalabs/slu/utils/mail_utils"
	"github.com/spf13/cobra"
)

var FlagSmtpHost string
var FlagSmtpPort int
var FlagFrom string
var FlagSmtpUser string
var FlagPassword string
var FlagTo string
var FlagSubject string
var FlagMessage string

var Cmd = &cobra.Command{
	Use:     "send",
	Short:   "Send mail",
	Aliases: []string{"s"},
	Args:    cobra.NoArgs,
	Run: func(c *cobra.Command, args []string) {
		user := FlagFrom
		if FlagSmtpUser != "" {
			user = FlagSmtpUser
		}
		err := mail_utils.SendSimpleMail(
			FlagSmtpHost,
			strconv.Itoa(FlagSmtpPort),
			user,
			FlagPassword,
			FlagFrom,
			FlagTo,
			FlagSubject,
			FlagMessage,
		)
		if err != nil {
			log.Fatal(err)
		}
	},
}

func init() {
	mail_cmd.Cmd.AddCommand(Cmd)
	Cmd.Flags().StringVarP(
		&FlagSmtpHost,
		"smtp-host",
		"H",
		"",
		"SMTP host (smtp.google.com)",
	)
	Cmd.MarkFlagRequired("smtp-host")
	Cmd.Flags().IntVarP(
		&FlagSmtpPort,
		"smtp-port",
		"P",
		0,
		"SMTP port (587, 25, ...)",
	)
	Cmd.MarkFlagRequired("smtp-port")
	Cmd.Flags().StringVarP(
		&FlagSmtpUser,
		"smtp-user",
		"U",
		"",
		"SMTP user (default from --from or -f)",
	)
	Cmd.Flags().StringVarP(
		&FlagFrom,
		"from",
		"f",
		"",
		"from (john@doe.com)",
	)
	Cmd.MarkFlagRequired("from")
	Cmd.Flags().StringVarP(
		&FlagPassword,
		"password",
		"p",
		"",
		"password",
	)
	Cmd.Flags().StringVarP(
		&FlagTo,
		"to",
		"t",
		"",
		"to (john@acme.com)",
	)
	Cmd.MarkFlagRequired("to")
	Cmd.Flags().StringVarP(
		&FlagSubject,
		"subject",
		"s",
		"",
		"Email subject",
	)
	Cmd.MarkFlagRequired("subject")
	Cmd.Flags().StringVarP(
		&FlagMessage,
		"message",
		"m",
		"",
		"email message",
	)
	Cmd.MarkFlagRequired("message")
}
