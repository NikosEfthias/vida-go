package user

import (
	"encoding/base64"
	"fmt"
	"io"
	"net/http"
	"strconv"
	"strings"

	"github.com/mugsoft/tools/bytesize"
	"github.com/mugsoft/vida/models"
	"github.com/mugsoft/vida/services"
	"github.com/mugsoft/vida/services/storage"
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
	u := storage.Get_user_by_token(token)
	if nil == u {
		return "", services.ERR_N_LOGIN
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
	u.Token = token
	if nil != err {
		return "", fmt.Errorf("err updating the user: %s", err.Error())
	}
	storage.Extend_token(u)
	return "success", nil
}

func Service_profile_pic(token string, file io.ReadCloser) (string, error) {
	const LIMIT_FILESIZE = int64(bytesize.MB * 10)
	var ALLOWED_MIMES = []string{"jpeg", "jpg", "png", "jpeg"}
	var MIME string
	if file == nil {
		return "", fmt.Errorf("cannot read the file")
	}
	defer file.Close()
	u := storage.Get_user_by_token(token)
	if nil == u {
		return "", services.ERR_N_LOGIN
	}
	var magic = make([]byte, 512)
	var filebuf = make([]byte, LIMIT_FILESIZE)
	n, err := file.Read(magic)
	if nil != err {
		return "", fmt.Errorf("cannot read the file, error: %s", err.Error())
	} else if n < 512 {
		return "", fmt.Errorf("too small")
	} else {
		//else is just for scoping variables here
		var valid_mime bool
		MIME = http.DetectContentType(magic)
		for i := range ALLOWED_MIMES {
			if strings.Contains(MIME, ALLOWED_MIMES[i]) {
				valid_mime = true
				break
			}
		}
		if !valid_mime {
			return "", fmt.Errorf("invalid image type")
		}
	}
	n, err = file.Read(filebuf)
	if nil != err {
		return "", fmt.Errorf("cannot read the file, error: %s", err.Error())
	} else if int64(n) > LIMIT_FILESIZE-512 {
		//we already consumed the first 512 bytes so if the read amount is big file is bigger no matter what
		return "", fmt.Errorf("too big")
	}
	var __data_url = "data:" + MIME + ";base64," + base64.StdEncoding.EncodeToString(append(magic, filebuf[:n]...))
	__token := u.Token //update destroys the old token so lets save it
	err = models.User_update(u.Id, map[string]interface{}{"profile_pic_url": __data_url}, u)
	u.Token = __token
	if nil != err {
		return "", fmt.Errorf("err updating the user: %s", err.Error())
	}
	storage.Extend_token(u)
	return "success", nil
}
