package user

import (
	"fmt"

	"github.com/mugsoft/vida/helpers"
	"github.com/mugsoft/vida/models"
	"github.com/mugsoft/vida/services/storage"
)

func Service_login(email, phone, password string) (string, error) {
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
	u.Token = helpers.Unique_id()
	storage.Add_or_update_user(u)
	return u.Token, nil
}
