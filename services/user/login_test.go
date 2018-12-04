package user

import (
	"testing"

	"gitlab.mugsoft.io/vida/go-api/models"
)

func TestService_forgot_password(t *testing.T) {
	//before hook {{{
	models.User_new(&models.User{
		Id:    "1234",
		Email: "testing@mugsoft.io",
	})
	defer models.User_delete("testing@mugsoft.io")
	//}}}
	var cases = map[string]bool{
		"fake@email.com":     true,
		"fake@emailcom":      true,
		"testing@mugsoft.io": false,
		"":                   true,
	}
	for mail, errored := range cases {
		_, err := Service_forgot_password(mail)
		var has_errored = (err != nil)
		if errored != has_errored {
			t.Fatalf("expecter error to be %v found %v,%v  %v", errored, has_errored, mail, err)
		}
	}
}
