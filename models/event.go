package models

import (
	"fmt"
	"time"

	"gitlab.mugsoft.io/vida/api/go-api/helpers"
)

const _COL_EVENT_STR = "events"

var _col_event = db_get().C(_COL_EVENT_STR)

type Event struct {
	//{{{
	Id        string    `json:"id" bson:"id"`
	Owner     string    `json:"owner" bson:"owner"`
	Guests    []string  `json:"guests" bson:"guests"`
	Title     string    `json:"title" bson:"title"`
	Loc       string    `json:"loc" bson:"loc"`
	Detail    string    `json:"detail" bson:"detail"`
	MaxGuest  int       `json:"max_guest" bson:"max_guest"`
	MinGuest  int       `json:"min_guest" bson:"min_guest"`
	Cost      float64   `json:"cost" bson:"cost"`
	Img       string    `json:"img" bson:"img"`
	StartDate time.Time `json:"start_date" bson:"start_date"`
	EndDate   time.Time `json:"end_date" bson:"end_date"`
	Votable   bool      `json:"votable" bson:"votable"`
	Defaults
	//}}}
}

func Event_new(e *Event) error {
	//{{{
	e.Id = helpers.Unique_id()
	if "" == e.Owner {
		return fmt.Errorf("missing event owner")
	}
	e.CreatedAt = time.Now()
	e.UpdatedAt = time.Now()
	return _col_event.Insert(e)
	//}}}
}
func Event_delete(id string) error {
	//{{{
	err := helpers.Check_id_format(id)
	if nil != err {
		return err
	}
	return _col_event.Remove(map[string]string{"id": id})
	//}}}
}
func Event_get_by_owner(owner_id string) {
	//{{{

	//}}}
}

func Event_get_by_id(id string) (*Event, error) {
	//{{{
	err := helpers.Check_id_format(id)
	if nil != err {
		return nil, err
	}
	_e := new(Event)
	err = _col_event.Find(map[string]string{"id": id}).One(_e)
	return _e, err
	//}}}
}

func Event_get_by_guest(guest_id string) {
	//{{{

	//}}}
}
