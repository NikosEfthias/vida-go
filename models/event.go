package models

import "time"

type Event struct {
	Id        string
	Owner     string
	Invitees  string //i1_mail;i1_phone:i2_mail;i2_phone syntax for compatibility with other dbs without owerhead this wont be used to search just for visual purposes
	StartDate time.Time
	EndDate   time.Time
	Defaults
}
