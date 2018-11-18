package models

import (
	//{{{
	"fmt"
	"time"

	"gitlab.mugsoft.io/vida/api/go-api/helpers"
	//}}}
)

const __COL_INVITATION_STR = "invitations"

var __col_invitation = db_get().C(__COL_INVITATION_STR)

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
	EventId string            `json:"-" bson:"event_id,omitempty"`
	Status  Invitation_status `json:"status,omitempty" bson:"status,omitempty"`
	Message string            `json:"message,omitempty" bson:"message,omitempty"`
	//Type can be app invitation or an invitation to a particular event
	Type      Invitation_type `json:"type,omitempty" bson:"type,omitempty"`
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
	var events = make([]*Invitation, 0)
	err := Invitation_get_by_invitee(typ, invitee, events, string(event_id))
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
	return i, __col_invitation.Insert(i)
	//}}}
}

//Invitation_get_by_invitee fetches the invitations based on type and invitee_id and optionally by event id
//and fills the result with that on unsuccessful fetch it will return an error otherwise nil
func Invitation_get_by_invitee(invitationtype Invitation_type, invitee string, result []*Invitation, event_id ...string) error {
	//{{{
	q := map[string]interface{}{
		"type":       invitationtype,
		"invitee_id": invitee,
	}
	if len(event_id) > 0 {
		q["event_id"] = event_id[0]
	}
	return __col_invitation.Find(q).All(&result)
	//}}}
}
