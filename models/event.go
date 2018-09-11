package models

import (
	"fmt"
	"time"

	"github.com/mugsoft/vida/helpers"
)

const __COL_EVENT_STR = "events"

var __col_event = db__get().C(__COL_EVENT_STR)

type Event struct {
	Id        string        `json:"id" bson:"id"`
	Owner     string        `json:"owner" bson:"owner"`
	Invitees  []*Invitation `json:"invitees" bson:"-"` //i1_mail;i1_phone:i2_mail;i2_phone syntax for compatibility with other dbs without owerhead this wont be used to search just for visual purposes
	Title     string        `json:"title" bson:"title"`
	Loc       string        `json:"loc" bson:"loc"`
	Detail    string        `json:"detail" bson:"detail"`
	MaxGuest  int           `json:"max_guest" bson:"max_guest"`
	MinGuest  int           `json:"min_guest" bson:"min_guest"`
	Cost      float64       `json:"cost" bson:"cost"`
	Img       string        `json:"img" bson:"img"`
	EventType int           `json:"event_type" bson:"event_type"` //1:static single date 2: votable single date (date range )
	StartDate time.Time     `json:"start_date" bson:"start_date"`
	EndDate   time.Time     `json:"end_date" bson:"end_date"`
	Defaults  `json:"defaults" bson:"defaults"`
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
