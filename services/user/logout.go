package user

import (
	"gitlab.mugsoft.io/vida/go-api/services"
	"gitlab.mugsoft.io/vida/go-api/services/storage"
)

func Service_logout(token string) (string, error) {
	//{{{
	u := storage.Get_user_by_token(token)
	if nil == u {
		return "", services.ERR_N_LOGIN
	}
	storage.Remove_user_by_token(token)
	return "success", nil
	//}}}
}
