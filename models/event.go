package models

import (
	"fmt"
	"time"

	"gitlab.mugsoft.io/vida/go-api/helpers"
	"gopkg.in/mgo.v2/bson"
)

const _COL_EVENT_STR = "events"

var _col_event = db_get().C(_COL_EVENT_STR)

type Event struct {
	//{{{
	Id          string        `json:"id" bson:"id"`
	Owner       string        `json:"owner" bson:"owner"`
	Invitations []*Invitation `json:"invitations" bson:"-"`
	Title       string        `json:"title" bson:"title"`
	Loc         string        `json:"loc" bson:"loc"`
	Detail      string        `json:"detail" bson:"detail"`
	MaxGuest    int           `json:"max_guest" bson:"max_guest"`
	MinGuest    int           `json:"min_guest" bson:"min_guest"`
	Cost        float64       `json:"cost" bson:"cost"`
	Img         string        `json:"img" bson:"img"`
	StartDate   time.Time     `json:"start_date" bson:"start_date"`
	EndDate     time.Time     `json:"end_date" bson:"end_date"`
	Votable     bool          `json:"votable" bson:"votable"`
	Defaults
	//}}}
}

func _event_fill_invitations(e *Event) error {
	//{{{
	invs, err := Invitation_get_by_event(e.Id)
	e.Invitations = invs
	return err //}}}
}
func (e *Event) GetGuestIds() []string {
	//{{{
	guest_ids := make([]string, len(e.Invitations))
	for i, inv := range e.Invitations {
		if nil != inv.Invitee {
			guest_ids[i] = inv.Invitee.Id
		}
	}
	return guest_ids //}}}
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
	err = _col_event.Remove(map[string]string{"id": id})
	if nil != err {
		return err
	}
	_, err = _col_invitation.RemoveAll(map[string]string{"event_id": id})
	return err
	//}}}
}
func Event_get_by_owner(owner_id string, page int) ([]*Event, error) {
	//{{{
	events := []*Event{}
	err := _col_event.Find(map[string]string{"owner": owner_id}).Skip(DATA_PER_PAGE * page).Limit(DATA_PER_PAGE).All(&events)
	if nil != err {
		for _, event := range events {
			_event_fill_invitations(event)
		}
	}
	return events, err
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
	_event_fill_invitations(_e)
	return _e, err
	//}}}
}
func Event_update(event_id, field string, value interface{}) error {
	//{{{
	fields := map[string]interface{}{
		"updated_at": time.Now(),
		field:        value,
	}
	return _col_event.Update(bson.M{"id": event_id}, bson.M{"$set": fields}) //}}}
}

func Event_get_by_guest(guest_id string, page int, filters map[string]interface{}) ([]*Event, error) {
	//{{{
	invs := []*Invitation{}
	var query map[string]interface{}
	if nil != filters {
		query = filters
	} else {
		query = map[string]interface{}{}
	}
	query["invitee_id"] = guest_id
	err := _col_invitation.Find(query).Skip(page * DATA_PER_PAGE).Limit(DATA_PER_PAGE).All(&invs)
	if nil != err {
		return nil, err
	}
	var query_2 = map[string]map[string][]string{"id": {"$in": []string{} /**/}}
	for _, inv := range invs {
		query_2["id"]["$in"] = append(query_2["id"]["$in"], inv.EventId)
	}
	var events = make([]*Event, 0, len(invs))
	err = _col_event.Find(query_2).All(&events)
	for _, e := range events {
		_event_fill_invitations(e)
	}
	return events, err
	//}}}
}
