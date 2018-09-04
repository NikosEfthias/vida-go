package models

import (
	"fmt"
	"time"

	"github.com/mugsoft/vida/helpers"
)

const __COL_EVENT_STR = "events"

var __col_event = db__get().C(__COL_EVENT_STR)

type Event struct {
	Id        string
	Owner     string
	Invitees  string //i1_mail;i1_phone:i2_mail;i2_phone syntax for compatibility with other dbs without owerhead this wont be used to search just for visual purposes
	Title     string
	Loc       string
	Detail    string
	MaxGuest  int
	MinGuest  int
	Cost      float64
	Img       string
	StartDate time.Time
	EndDate   time.Time
	Defaults
}

func Event_new(e *Event) error {
	e.Id = helpers.Unique_id()
	if "" == e.Owner {
		return fmt.Errorf("missing event owner")
	}
	e.CreatedAt = time.Now()
	e.UpdatedAt = time.Now()
	return __col_event.Insert(e)
}
