package models

import "fmt"

const _COL_VOTE_STR = "votes"

var _col_vote = db_get().C(_COL_VOTE_STR)

type Vote struct {
	//{{{
	Id        string
	Voter_id  string
	Event_id  string
	Vote_time int64
	//}}}
}

func Vote_event(event_id, user_id string, time int64) error {
	return fmt.Errorf("not implemented yet")
}
