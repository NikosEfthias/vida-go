package user

import (
	"github.com/mugsoft/vida/services"
	"github.com/mugsoft/vida/services/storage"
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
