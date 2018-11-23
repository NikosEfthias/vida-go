package models

import (
	"fmt"
	"time"

	"gitlab.mugsoft.io/vida/api/go-api/helpers"
	"gopkg.in/mgo.v2/bson"
)

const _COL_USER_STR = "users"

var _col = db_get().C(_COL_USER_STR)

type User struct {
	//{{{
	Id            string    `bson:"id" json:"id"`
	Name          string    `bson:"name" json:"name"`
	Lastname      string    `bson:"lastname" json:"lastname"`
	Phone         string    `bson:"phone" json:"phone"`
	Email         string    `bson:"email" json:"email"`
	Notification  int       `bson:"notification" json:"notification"`
	FbAccountName string    `json:"fb_account_name" bson:"fb_account_name"`
	FbProfilePic  string    `json:"fb_profile_pic" bson:"fb_profile_pic"`
	Password      string    `bson:"password" json:"-"`
	Login_expires time.Time `bson:"-" json:"-"`
	Token         string    `bson:"-" json:"token"`
	ProfilePicURL string    `bson:"profile_pic_url" json:"profile_pic_url"`
	PassReset     bool      `bson:"pass_reset" json:"pass_reset,omitempty"`
	Tmp           bool      `json:"tmp,omitempty" bson:"-"`
	Defaults
	//}}}
}

//User_new generates id and date fields of the user and hashes password then saves
func User_new(u *User) error {
	//{{{
	//{{{ error checks
	if "" == u.Email && "" == u.Phone {
		return fmt.Errorf("missing email and phone")
	}
	if nil == User_get(u) {
		return fmt.Errorf("user exists")
	} //}}}
	u.Id = helpers.Unique_id()
	u.Password = Hash_password(u, u.Password)
	u.CreatedAt = time.Now()
	u.UpdatedAt = time.Now()
	return _col.Insert(u)
	//}}}
}

//Hash_password hashes user password with salt and generates id if it is empty
func Hash_password(u *User, pass string) string {
	//{{{
	if "" == u.Id {
		u.Id = helpers.Unique_id()
	}
	return helpers.MD5(u.Id + pass)
	//}}}
}

func User_get(u *User) error {
	//{{{
	_q := []bson.M{}
	if "" != u.Email {
		_q = append(_q, bson.M{"email": u.Email})
	}
	if "" != u.Phone {
		_q = append(_q, bson.M{"phone": u.Phone})
	}
	if "" == u.Email && "" == u.Phone {
		return fmt.Errorf("missing email and phone")
	}
	return _col.Find(bson.M{
		"$or": _q,
	}).One(u)
	//}}}
}
func User_get_by_email(email string) (*User, error) {
	//{{{
	if !helpers.Is_email_valid(email) {
		return nil, fmt.Errorf("invalid email address")
	}
	usr := &User{
		Email: email,
	}
	err := User_get(usr)
	if nil != err {
		return nil, err
	}
	return usr, nil
	//}}}
}
func User_update(userid string, fields map[string]interface{}, updatedU *User) error {
	//{{{
	var _fields_with_pdatedAt = map[string]interface{}{
		"updated_at": time.Now(),
	}
	for k, v := range fields {
		_fields_with_pdatedAt[k] = v
	}
	err := _col.Update(bson.M{"id": userid}, bson.M{"$set": _fields_with_pdatedAt})
	if nil != err {
		return err
	}
	if nil == updatedU {
		return nil
	}
	updatedU.Id = userid
	return User_get(updatedU)
	//}}}
}
func User_new_tmp(email string) (*User, error) {
	//{{{
	u := &User{
		Email: email,
	}
	//{{{ error checks
	if nil == User_get(u) {
		return nil, fmt.Errorf("user exists")
	} //}}}
	u.Id = helpers.Unique_id()
	u.Token = helpers.Unique_id()
	u.Tmp = true
	return u, nil //}}}
}
