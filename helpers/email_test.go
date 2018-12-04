package helpers

import (
	"testing"

	"gitlab.mugsoft.io/vida/go-api/config"
)

func TestSendMail(t *testing.T) {
	err := SendMail(config.Get("APP_EMAIL_ADDR"), config.Get("APP_EMAIL_PASSWD"), config.Get("APP_EMAIL_ADDR"),
		[]string{"testing@mugsoft.io", "nikos@mugsoft.io"},
		config.Get("SMTP_ADDR"),
		"Hello from vida mailing",
		"testing",
	)
	if nil != err {
		t.Fatal("err:", err)
	}
}
