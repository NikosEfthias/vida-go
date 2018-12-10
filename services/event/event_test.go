package event

import (
	"fmt"
	"testing"
	"time"

	"gitlab.mugsoft.io/vida/go-api/helpers"
	"gitlab.mugsoft.io/vida/go-api/models"
	"gitlab.mugsoft.io/vida/go-api/services"
	"gitlab.mugsoft.io/vida/go-api/services/storage"
)

func TestService_vote(t *testing.T) { //{{{
	type _case struct { //case struct{{{
		description string
		expected    interface{}
		token       string
		event_id    string
		time        string
	} //}}}
	//test variables{{{
	var valid_date string = fmt.Sprint(time.Now().Add(time.Hour * 25).Unix())
	var early_date string = fmt.Sprint(time.Now().Unix())
	var late_date string = fmt.Sprint(time.Now().Add(time.Hour * 50).Unix())
	var valid_event = &models.Event{
		Owner:     "123",
		StartDate: time.Now().Add(time.Hour * 24),
		EndDate:   time.Now().Add(time.Hour * 26),
		Votable:   true,
	}
	var valid_event_accepted = &models.Event{
		Owner:     "123",
		StartDate: time.Now().Add(time.Hour * 24),
		EndDate:   time.Now().Add(time.Hour * 26),
		Votable:   true,
	} //}}}
	//setup{{{
	storage.Add_or_update_user(&models.User{
		Id:    "123",
		Token: "123",
	})

	err := models.Event_new(valid_event)
	_f_on_err(t, err)
	err = models.Event_new(valid_event_accepted)
	_f_on_err(t, err)
	valid_inv, err := models.Invitation_create(models.INV_EVENT, []rune(valid_event.Id), "123", "123", "hello")
	_f_on_err(t, err)
	accepted_valid_inv, err := models.Invitation_create(models.INV_EVENT, []rune(valid_event_accepted.Id), "1234", "123", "hello")
	_f_on_err(t, err)
	err = models.Invitation_accept(valid_event_accepted.Id, "123")
	_f_on_err(t, err)
	//}}}
	//teardown{{{
	defer storage.Remove_user_by_token("123")
	defer models.Event_delete(valid_event.Id)
	defer models.Event_delete(valid_event_accepted.Id)
	defer models.Invitation_delete(valid_inv.Id)
	defer models.Invitation_delete(accepted_valid_inv.Id)
	// }}}
	//cases{{{

	var cases = []_case{
		{
			description: "invalid token must fail with login error",
			expected:    services.ERR_N_LOGIN,
			token:       "dummy",
			event_id:    "1234",
			time:        valid_date,
		},
		{"visually invalid time parameter must fail", ERR_INVALID_TIME, "123", valid_event.Id, "abc"},
		{"visually invalid event id must fail", ERR_INVALID_EVENT_ID, "123", "test123", valid_date},
		{"non existing event id must fail", ERR_EVENT_NOT_FOUND, "123", helpers.Unique_id(), valid_date},
		{"only if the person has invitation on event can vote", ERR_NOT_INVITED, "123", valid_event.Id, valid_date},
		{"time cannot be smaller than the start date", ERR_INVALID_TIME_RANGE, "123", valid_event_accepted.Id, early_date},
		{"time cannot be greater than the end date", ERR_INVALID_TIME_RANGE, "123", valid_event_accepted.Id, late_date},
		{"with valid params error must be nil", nil, "123", valid_event_accepted.Id, valid_date},
	}
	//}}}
	//run cases{{{
	for _, c := range cases {
		fmt.Println(c.description)
		_, err := Service_vote(c.token, c.event_id, c.time)
		if c.expected != err {
			t.Fatalf("expected error to be %v\nfound %v", c.expected, err.Error())
		}
	} //}}}
	//check db lastly if data is there{{{
	votes, err := models.Votes_get_for_event(valid_event_accepted.Id)
	_f_on_err(t, err)
	var found = false
	for _, v := range votes {
		if v.Voter_id != "123" || v.Event_id != valid_event_accepted.Id {
			continue
		}
		found = true
		err := models.Delete_vote(v.Event_id, v.Voter_id)
		_f_on_err(t, err)
	}
	if !found {
		t.Fatal("vote cannot be found")
	} //}}}
} //}}}

func _f_on_err(t *testing.T, err error) {
	if nil != err {
		t.Fatalf("setup error %v", err)
	}
}
