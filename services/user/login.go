package user

import (
	//{{{
	"bytes"
	"fmt"
	"text/template"

	"gitlab.mugsoft.io/vida/api/go-api/helpers"
	"gitlab.mugsoft.io/vida/api/go-api/models"
	"gitlab.mugsoft.io/vida/api/go-api/services/storage"
	//}}}
)

func Service_login(email, phone, password string) (string, error) {
	//{{{
	//{{{
	if ("" == email && "" == phone) || "" == password {
		return "", fmt.Errorf("missing login fields")
	}
	var u = &models.User{
		Email: email,
		Phone: phone,
	}
	err := models.User_get(u)
	if nil != err {
		return "", fmt.Errorf("user does not exists")
	}
	if u.Password != models.Hash_password(u, password) {
		return "", fmt.Errorf("invalid password")
	}
	//}}}
	u.Token = helpers.Unique_id()
	storage.Add_or_update_user(u)
	return u.Token, nil
	//}}}
}
func Service_forgot_password(email string) (string, error) {
	//{{{
	u := &models.User{Email: email}
	err := models.User_get(u)
	if nil != err {
		return "", fmt.Errorf("no such user")
	}
	u.Token = helpers.Unique_id()
	u.PassReset = true
	err = models.User_update(u.Id, map[string]interface{}{"pass_reset": true}, nil)
	if nil != err {
		return "", fmt.Errorf("cannot complete the request")
	}
	var mail = new(bytes.Buffer)
	err = template.Must(template.New("mail").Parse("your token is {{.Token}}")).Execute(mail, u)
	if nil != err {
		return "", err
	}
	err = helpers.SendMailPreconfigured([]string{u.Email}, "Reset your password", mail.String())
	storage.Add_or_update_user(u)
	return "please check your email", err
	//}}}
}
