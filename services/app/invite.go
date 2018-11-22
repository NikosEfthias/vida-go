package app

import (
	//{{{
	"fmt"

	"gitlab.mugsoft.io/vida/api/go-api/helpers"
	"gitlab.mugsoft.io/vida/api/go-api/models"
	"gitlab.mugsoft.io/vida/api/go-api/services"
	"gitlab.mugsoft.io/vida/api/go-api/services/storage"
	//}}}
)

func Service_invite_people(token string, people []string) (string, error) {
	//{{{
	u := storage.Get_user_by_token(token)
	//error checks {{{
	if nil == u {
		return "", services.ERR_N_LOGIN
	}
	if len(people) < 1 {
		return "", fmt.Errorf("no people to invite")
	}
	//}}}
	var errs = make([]string, 0, len(people))
	for _, p := range people {
		//for each people do
		//issue a token for a temporary user {{{
		usr, err := models.User_new_tmp(p)
		if nil != err {
			helpers.Log(helpers.ERR, "Cannot send app invitation to ", p, "err:", err.Error())
			errs = append(errs, err.Error())
			continue
		}
		if nil == usr {
			helpers.Log(helpers.ERR, "weird null usr behaviour must be checked")
			errs = append(errs, "weird null pointer check this")
			continue
		}
		storage.Add_or_update_user(usr)
		//}}}
		//add a new invitation to the db {{{
		inv, err := models.Invitation_create(models.INV_APP, nil, u.Id, usr.Id, "PLEASE JOIN THE AWESOME VIDA "+usr.Token)
		if nil != err {
			helpers.Log(helpers.ERR, "Cannot send app invitation to ", p, "err:", err.Error())
			errs = append(errs, err.Error())
			continue
		}
		//}}}
		//send email to the user {{{
		err = helpers.SendOneMailPreconfigured(p, "welcome to vida", inv.Message)
		if nil != err {
			return "", err
		}
		//}}}
	}
	if len(errs) > 0 {
		return "partial success may be ", fmt.Errorf("%v", errs)
	}
	return "success", nil //}}}
}
