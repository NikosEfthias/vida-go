package models

type Invitation struct {
	Id        string
	EventId   string
	Status    int //0:pending,1:accepted,2:declined
	Message   string
	Type      int //0:event,1:app invitation
	InviterId string
	InviteeId string //if invitee is already member
	Defaults
}
