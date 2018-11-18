package user

import "testing"

func TestService_forgot_password(t *testing.T) {
	var cases = map[string]bool{
		"fake@email.com":   true,
		"fake@emailcom":    true,
		"nikos@mugsoft.io": false,
		"":                 true,
	}
	for mail, errored := range cases {
		_, err := Service_forgot_password(mail)
		var has_errored = (err != nil)
		if errored != has_errored {
			t.Fatalf("expecter error to be %v found %v,%v  %v", errored, has_errored, mail, err)
		}
	}
}
