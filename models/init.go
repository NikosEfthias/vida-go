package models

import (
	"time"

	"github.com/mugsoft/vida/config"
	mgo "gopkg.in/mgo.v2"
)

var db *mgo.Database

func db__get() *mgo.Database {
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
