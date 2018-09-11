package models

type Invitation struct {
	Id        string `json:"-" bson:"id,omitempty"`
	EventId   string `json:"-" bson:"event_id,omitempty"`
	Status    int    `json:"status,omitempty" bson:"status,omitempty"` //0:pending,1:accepted,2:declined
	Message   string `json:"message,omitempty" bson:"message,omitempty"`
	Type      int    `json:"type,omitempty" bson:"type,omitempty"` //0:event,1:app invitation
	InviterId string `json:"inviter_id,omitempty" bson:"inviter_id,omitempty"`
	InviteeId string `json:"invitee_id,omitempty" bson:"invitee_id,omitempty"` //if invitee is already member
	Invitee   *User  `json:"invitee,omitempty" bson:"-"`
	Defaults  `json:"defaults,omitempty" bson:"defaults,omitempty"`
}
