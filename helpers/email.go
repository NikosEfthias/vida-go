package helpers

import (
	"net/smtp"
	"strings"
)

func SendMail(uname, password, from string, to []string, server string, subject, msg string) error {
	headers := "Subject: " + subject + "\r\n\r\n"
	return smtp.SendMail(server,
		smtp.PlainAuth("", uname, password, strings.Split(server, ":")[0]),
		from,
		to,
		[]byte(headers+msg),
	)
}
