package user

import (
	"fmt"
	"strconv"

	"github.com/mugsoft/vida/models"
	"github.com/mugsoft/vida/services/storage"
)

func Service_get(token string) (interface{}, error) {
	u := storage.Get_user_by_token(token)
	if nil == u {
		return nil, fmt.Errorf("not logged in")
	}
	return u, nil
}

//Service_update updates the user data
/*
	All three parameters passed at first as a string but then value can be reassigned with a different type internally and thats the reason its defined as an interface{}
*/
func Service_update(key, token string, value interface{}) (string, error) {
	u := storage.Get_user_by_token(token)
	if nil == u {
		return "", fmt.Errorf("not logged in")
	}
	var err error
	if value.(string) == "" {
		return "", fmt.Errorf("empty value")
	}
	switch key {
	case "name":
	case "lastname":
	case "email":
		if nil == models.User_get(&models.User{Email: value.(string)}) {
			return "", fmt.Errorf("in use")
		}
	case "password":
		value = models.Hash_password(u, value.(string))
	case "phone":
		if nil == models.User_get(&models.User{Phone: value.(string)}) {
			return "", fmt.Errorf("in use")
		}
	case "notification":
		var i int
		i, err = strconv.Atoi(value.(string))
		if nil != err {
			return "", fmt.Errorf("invalid data %s, error: %s", value.(string), err.Error())
		}
		value = i
	case "fb_account_name":
	case "fb_profile_pic":
	default:
		return "", fmt.Errorf("unknown update field %s", key)
	}
	err = models.User_update(u.Id, map[string]interface{}{key: value}, u)
	if nil != err {
		return "", fmt.Errorf("err updating the user: %s", err.Error())
	}
	u.Token = token
	storage.Extend_token(u)
	return "success", nil
}
