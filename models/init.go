package models

import (
	"time"

	"gitlab.mugsoft.io/vida/api/go-api/config"
	mgo "gopkg.in/mgo.v2"
)

var db *mgo.Database

func db_get() *mgo.Database {
	if nil != db {
		return db
	}
	_ses, err := mgo.DialWithTimeout(config.Get("DB_ADDR"), time.Second*5)
	if nil != err {
		panic(err)
	}
	db = _ses.DB(config.Get("DB"))
	return db
}

type Defaults struct {
	CreatedAt time.Time `bson:"created_at" json:"created_at"`
	UpdatedAt time.Time `bson:"updated_at" json:"updated_at"`
}
