package models

import (
	"fmt"
	"time"

	"github.com/mugsoft/vida/helpers"
	"gopkg.in/mgo.v2/bson"
)

const __COL_USER_STR = "users"

var __col = db__get().C(__COL_USER_STR)

type User struct {
	Id             string    `bson:"id" json:"id"`
	Name           string    `bson:"name" json:"name"`
	Lastname       string    `bson:"lastname" json:"lastname"`
	Phone          string    `bson:"phone" json:"phone"`
	Email          string    `bson:"email" json:"email"`
	Notification   int       `bson:"notification" json:"notification"`
	FbAccountName  string    `json:"fb_account_name" bson:"fb_account_name"`
	FbProfilePic   string    `json:"fb_profile_pic" bson:"fb_profile_pic"`
	Password       string    `bson:"password" json:"-"`
	CreatedAt      time.Time `bson:"created_at" json:"created_at"`
	UpdatedAt      time.Time `bson:"updated_at" json:"updated_at"`
	Login__expires time.Time `bson:"-" json:"-"`
	Token          string    `bson:"-" json:"token"`
	ProfilePicURL  string    `bson:"profile_pic_url" json:"profile_pic_url"`
}

//User_new generates id and date fields of the user and hashes password then saves
func User_new(u *User) error {
	if "" == u.Email && "" == u.Phone {
		return fmt.Errorf("missing email and phone")
	}
	if nil == User_get(u) {
		return fmt.Errorf("user exists")
	}
	u.Id = helpers.Unique_id()
	u.Password = Hash_password(u, u.Password)
	u.CreatedAt = time.Now()
	u.UpdatedAt = time.Now()
	return __col.Insert(u)
}

//Hash_password hashes user password with salt and generates id if it is empty
func Hash_password(u *User, pass string) string {
	if "" == u.Id {
		u.Id = helpers.Unique_id()
	}
	return helpers.MD5(u.Id + pass)
}

func User_get(u *User) error {
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
	return __col.Find(bson.M{
		"$or": _q,
	}).One(u)
}

func User_update(userid string, fields map[string]interface{}, updatedU *User) error {
	err := __col.Update(bson.M{"id": userid}, bson.M{"$set": fields})
	if nil != err {
		return err
	}
	if nil == updatedU {
		return nil
	}
	updatedU.Id = userid
	return User_get(updatedU)
}
