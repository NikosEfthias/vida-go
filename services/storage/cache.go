package storage

import (
	"sync"
	"time"

	"gitlab.mugsoft.io/vida/api/go-api/helpers"
	"gitlab.mugsoft.io/vida/api/go-api/models"
)

var _cache_user = struct {
	sync.Mutex
	by_id    map[string]*models.User
	by_token map[string]*models.User
}{
	//{{{
	sync.Mutex{},
	map[string]*models.User{},
	map[string]*models.User{},
	//}}}
}

func Get_user_by_id(k string) *models.User {
	//{{{
	_cache_user.Lock()
	defer _cache_user.Unlock()
	return _is_expired(_cache_user.by_id[k])
	//}}}
}

func Get_user_by_token(k string) *models.User {
	//{{{
	_cache_user.Lock()
	defer _cache_user.Unlock()
	return _is_expired(_cache_user.by_token[k])
	//}}}
}

func Add_or_update_user(u *models.User) {
	//{{{
	if nil == u {
		helpers.Log(helpers.ERR, "null pointer passed into add or update")
		return
	}
	if u.Token == "" {
		helpers.Log(helpers.ERR, "missing token in add or update for cache")
		return
	}
	u.Login_expires = time.Now().Add(time.Hour * 3)
	_cache_user.Lock()
	defer _cache_user.Unlock()
	old_u := _cache_user.by_id[u.Id]
	_cache_user.by_id[u.Id] = u
	if nil != old_u {
		//dont allow multi login
		//new login voids the old one
		delete(_cache_user.by_token, old_u.Token)
	}
	_cache_user.by_token[u.Token] = u
	//}}}
}

func _is_expired(u *models.User) *models.User {
	//{{{
	if nil == u {
		return u
	}
	if u.Login_expires.Unix() < time.Now().Unix() {
		return nil
	}
	u.Login_expires = time.Now().Add(time.Hour * 4)
	return u
	//}}}
}

func Extend_token(u *models.User) {
	//{{{
	if nil == u {
		return
	}
	u.Login_expires = time.Now().Add(time.Hour * 4)
	//}}}
}
func Remove_user_by_token(token string) {
	//{{{
	_cache_user.Lock()
	u, ok := _cache_user.by_token[token]
	delete(_cache_user.by_token, token)
	if ok {
		delete(_cache_user.by_id, u.Id)
	}
	_cache_user.Unlock()
	//}}}
}

//TODO:  add setter so data can be locked
//TODO:  add function to remove old docs for every 5 mins
