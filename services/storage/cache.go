package storage

import (
	"sync"
	"time"

	"github.com/mugsoft/vida/helpers"
	"github.com/mugsoft/vida/models"
)

var __cache__user = struct {
	sync.Mutex
	by__id    map[string]*models.User
	by__token map[string]*models.User
}{
	sync.Mutex{},
	map[string]*models.User{},
	map[string]*models.User{},
}

func Get_user_by_id(k string) *models.User {
	__cache__user.Lock()
	defer __cache__user.Unlock()
	return __is__expired(__cache__user.by__id[k])
}

func Get_user_by_token(k string) *models.User {
	__cache__user.Lock()
	defer __cache__user.Unlock()
	return __is__expired(__cache__user.by__token[k])
}

func Add_or_update_user(u *models.User) {
	if u.Token == "" {
		helpers.Log(helpers.ERR, "missing token in add or update for cache")
		return
	}
	u.Login__expires = time.Now().Add(time.Hour * 3)
	__cache__user.Lock()
	defer __cache__user.Unlock()
	old_u := __cache__user.by__id[u.Id]
	__cache__user.by__id[u.Id] = u
	if nil != old_u {
		//dont allow multi login
		//new login voids the old one
		delete(__cache__user.by__token, old_u.Token)
	}
	__cache__user.by__token[u.Token] = u
}

func __is__expired(u *models.User) *models.User {
	if nil == u {
		return u
	}
	if u.Login__expires.Unix() < time.Now().Unix() {
		return nil
	}
	u.Login__expires = time.Now().Add(time.Hour * 4)
	return u
}

func Extend_token(u *models.User) {
	if nil == u {
		return
	}
	u.Login__expires = time.Now().Add(time.Hour * 4)
}

//TODO:  add setter so data can be locked
//TODO:  add function to remove old docs for every 5 mins
