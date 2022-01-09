package mail_utils

import "net/smtp"

func RawSendMail(
	smtpHost string,
	smtpPort string,
	user string,
	password string,
	from string,
	to string,
	rawMessage []byte,
) error {
	var auth smtp.Auth

	if password != "" {
		auth = smtp.PlainAuth(from, user, password, smtpHost)
	}

	return smtp.SendMail(smtpHost+":"+smtpPort, auth, from,
		[]string{to}, rawMessage)
}

func SendSimpleMail(
	smtpHost string,
	smtpPort string,
	user string,
	password string,
	from string,
	to string,
	subject string,
	message string,
) error {
	rawMessage := []byte("To: " + to + "\r\n" +
		"From: " + from + "\r\n" +
		"Subject: " + subject + "\r\n" +
		"\r\n" +
		message + "\r\n")
	return RawSendMail(smtpHost, smtpPort, user, password, from, to, rawMessage)
}
