package smtp_proxy_utils

import (
	"bytes"
	"log"
	"net"
	"net/mail"
	"net/smtp"

	"github.com/mhale/smtpd"
)

func RunSimpleSMTPProxy(local string, remote string) {
	var err error
	mailHandler := func(origin net.Addr, from string, to []string, data []byte) error {
		msg, err := mail.ReadMessage(bytes.NewReader(data))
		if err != nil {
			log.Println(err)
		}
		subject := msg.Header.Get("Subject")
		log.Printf("Received mail from %s for %s with subject: %s", from, to[0], subject)
		smtp.SendMail(remote, nil, from, to, data)
		return nil
	}
	err = smtpd.ListenAndServe(local, mailHandler, "MySMTPServer", "")
	if err != nil {
		log.Fatalln(err)
	}
}
