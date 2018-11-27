package models

import (
	//{{{
	"fmt"
	"time"

	"gitlab.mugsoft.io/vida/api/go-api/helpers"
	//}}}
)

const __COL_INVITATION_STR = "invitations"

var _col_invitation = db_get().C(__COL_INVITATION_STR)

type Invitation_type int
type Invitation_status int

const (
	INV_APP Invitation_type = iota
	INV_EVENT
)
const (
	INV_STATUS_PENDING Invitation_status = iota
	INV_STATUS_ACCEPTED
	INV_STATUS_DECLINED
)

//Invitation type
type Invitation struct {
	//{{{
	//Id of the Invitation
	Id string `json:"-" bson:"id,omitempty"`
	//EventId of the event the Invitation issued for
	EventId string            `json:"event_id" bson:"event_id,omitempty"`
	Status  Invitation_status `json:"status" bson:"status"`
	Message string            `json:"message,omitempty" bson:"message,omitempty"`
	//Type can be app invitation or an invitation to a particular event
	Type      Invitation_type `json:"type" bson:"type"`
	InviterId string          `json:"inviter_id,omitempty" bson:"inviter_id,omitempty"`
	InviteeId string          `json:"invitee_id,omitempty" bson:"invitee_id,omitempty"`
	//Invitee user details. This is ignored during db insertion
	Invitee  *User `json:"invitee,omitempty" bson:"-"`
	Defaults `json:"defaults,omitempty" bson:"defaults,omitempty"`
	//}}}
}

//Invitation_create creates an invitation the last parameter will be there only if the event type is Invitation_event
// event id is a rune slice rather than a string so it is nillable.
func Invitation_create(typ Invitation_type, event_id []rune, inviter string, invitee string, message string) (*Invitation, error) {
	//{{{
	switch {
	case typ == INV_EVENT && len(event_id) == 0:
		return nil, fmt.Errorf("missing event id on event invitation")
	case inviter == "":
		return nil, fmt.Errorf("missing inviter")
	case invitee == "":
		return nil, fmt.Errorf("missing invitee")
	case message == "":
		return nil, fmt.Errorf("missing invitation body")
	}
	events, err := Invitation_get_by_invitee(typ, invitee, string(event_id))
	if len(events) > 0 && nil == err {
		return nil, fmt.Errorf("invitee has already been invited to this event")
	}
	i := new(Invitation)
	i.Id = helpers.Unique_id()
	i.Type = typ
	i.EventId = string(event_id)
	i.InviterId = inviter
	i.InviteeId = invitee
	i.Message = message
	i.CreatedAt = time.Now()
	i.UpdatedAt = time.Now()
	i.Status = INV_STATUS_PENDING
	return i, _col_invitation.Insert(i)
	//}}}
}

//_inv_user_get fills the invitee field of the invitation
func _inv_user_get(inv *Invitation) error {
	//{{{
	invitee, err := User_get_by_id(inv.InviteeId)
	inv.Invitee = invitee
	return err //}}}
}

//Invitation_get_by_event gets all the invitations by event and adds user data to the invitations as well if possible
func Invitation_get_by_event(event_id string) ([]*Invitation, error) {
	//{{{
	var invs = make([]*Invitation, 0, 100)
	err := _col_invitation.Find(map[string]string{"event_id": event_id}).All(&invs)
	if nil != err {
		return nil, err
	}
	for _, inv := range invs {
		_inv_user_get(inv)
	}
	return invs, nil //}}}
}

//Invitation_get_by_invitee fetches the invitations based on type and invitee_id and optionally by event id
//and fills the result with that on unsuccessful fetch it will return an error otherwise nil
func Invitation_get_by_invitee(invitationtype Invitation_type, invitee string, event_id string) ([]*Invitation, error) {
	//{{{
	q := map[string]interface{}{
		"type":       invitationtype,
		"invitee_id": invitee,
	}
	var result []*Invitation
	if len(event_id) > 0 {
		q["event_id"] = event_id
		result = make([]*Invitation, 0, 1)
	} else {
		result = make([]*Invitation, 0, 10)
	}
	err := _col_invitation.Find(q).All(&result)
	return result, err
	//}}}
}

//Invitation_accept sets the status of invitation to INV_STATUS_ACCEPTED
func Invitation_accept(inv_id string) error {
	//{{{
	return _col_invitation.Update(map[string]string{"id": inv_id}, map[string]Invitation_status{"status": INV_STATUS_ACCEPTED}) //}}}
}

//Invitation_decline sets the status of invitation to INV_STATUS_ACCEPTED
func Invitation_decline(inv_id string) error {
	//{{{
	return _col_invitation.Update(map[string]string{"id": inv_id}, map[string]Invitation_status{"status": INV_STATUS_DECLINED}) //}}}
}
