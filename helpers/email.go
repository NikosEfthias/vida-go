package helpers

import (
	"fmt"
	"net/smtp"
	"strings"

	"gitlab.mugsoft.io/vida/go-api/config"
)

func SendMail(uname, password, from string, to []string, server string, subject, msg string) error {
	//{{{
	headers := "Subject: " + subject +
		"\r\nFrom: <" + from + "> vida\r\n" +
		"Content-Type: Text/HTML\r\n"
	if len(to) > 0 {
		for _, m := range to {
			headers += fmt.Sprintf("To: %s\r\n", m)
		}
	}
	headers += "\r\n\r\n"
	return smtp.SendMail(server,
		smtp.PlainAuth("", uname, password, strings.Split(server, ":")[0]),
		from,
		to,
		[]byte(headers+msg),
	)
	//}}}
}
func SendMailPreconfigured(to []string, subject, msg string) error {
	return SendMail(
		//{{{
		config.Get("APP_EMAIL_ADDR"),
		config.Get("APP_EMAIL_PASSWD"),
		config.Get("APP_EMAIL_ADDR"),
		to,
		config.Get("SMTP_ADDR"),
		subject,
		msg,
		//}}}
	)
}
func SendOneMailPreconfigured(to string, subject, msg string) error {
	//{{{
	return SendMailPreconfigured([]string{to}, subject, msg)
	//}}}
}
