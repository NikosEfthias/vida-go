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
	Id        string    `bson:"id"`
	Name      string    `bson:"name"`
	Lastname  string    `bson:"lastname"`
	Phone     string    `bson:"phone"`
	Email     string    `bson:"email"`
	Password  string    `bson:"password"`
	CreatedAt time.Time `bson:"created_at"`
	UpdatedAt time.Time `bson:"updated_at"`
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
