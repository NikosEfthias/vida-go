package helpers

import "testing"

func TestSendMail(t *testing.T) {
	err := SendMail("info@vidaevents.org", "_test_vida_", "info@vidaevents.org",
		[]string{"nikos@mugsoft.io", "furkanaydin@mugsoft.io", "paulchavez.uw@icloud.com"},
		"smtp.gmail.com:587",
		"Hello from vida mailing",
		"testing",
	)
	if nil != err {
		t.Fatal("err:", err)
	}
}
