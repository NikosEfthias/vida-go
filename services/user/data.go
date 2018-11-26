package user

import (
	//{{{
	"fmt"
	"io"
	"strconv"

	"github.com/mugsoft/tools/bytesize"
	"gitlab.mugsoft.io/vida/api/go-api/helpers"
	"gitlab.mugsoft.io/vida/api/go-api/models"
	"gitlab.mugsoft.io/vida/api/go-api/services"
	"gitlab.mugsoft.io/vida/api/go-api/services/storage"
	//}}}
)

func Service_get(token string) (interface{}, error) {
	u := storage.Get_user_by_token(token)
	if nil == u {
		return nil, services.ERR_N_LOGIN
	}
	return u, nil
}

//Service_update updates the user data
/*
	All three parameters passed at first as a string but then value can be reassigned with a different type internally and thats the reason its defined as an interface{}
*/
func Service_update(key, token string, value interface{}) (string, error) {
	//{{{
	u := storage.Get_user_by_token(token)
	if nil == u {
		return "", services.ERR_N_LOGIN
	}
	var err error
	if value.(string) == "" {
		return "", fmt.Errorf("empty value")
	}
	switch key {
	//check update key
	//{{{
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
		//}}}
	}
	if key != "password" && u.PassReset {
		return "", fmt.Errorf("pass reset key cannot be used for anything other than password reset")
	}
	u.PassReset = false
	err = models.User_update(u.Id, map[string]interface{}{key: value, "pass_reset": false}, u)
	u.Token = token
	if nil != err {
		return "", fmt.Errorf("err updating the user: %s", err.Error())
	}
	storage.Extend_token(u)
	return "success", nil
	//}}}
}

func Service_profile_pic(token string, file io.Reader) (string, error) {
	//{{{
	const LIMIT_FILESIZE = bytesize.MB * 10
	var ALLOWED_MIMES = []string{"jpeg", "jpg", "png", "jpeg"}
	if file == nil {
		return "", fmt.Errorf("cannot read the file")
	}
	u := storage.Get_user_by_token(token)
	if nil == u {
		return "", services.ERR_N_LOGIN
	}
	if u.PassReset {
		return "", fmt.Errorf("pass reset key cannot be used for anything other than password reset")
	}
	__data_url, err := helpers.Multipart_to_data_url(file, LIMIT_FILESIZE, ALLOWED_MIMES)
	if nil != err {
		return "", fmt.Errorf("cannot process the file error : %s", err.Error())
	}
	__token := u.Token //update destroys the old token so lets save it
	err = models.User_update(u.Id, map[string]interface{}{"profile_pic_url": __data_url}, u)
	u.Token = __token
	if nil != err {
		return "", fmt.Errorf("err updating the user: %s", err.Error())
	}
	storage.Extend_token(u)
	return "success", nil
	//}}}
}
