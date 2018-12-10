package models

import (
	"time"

	"gopkg.in/mgo.v2/bson"
)

const _COL_VOTE_STR = "votes"

var _col_vote = db_get().C(_COL_VOTE_STR)

type Vote struct {
	//{{{
	Voter_id string    `json:"voter_id" bson:"voter_id"`
	Event_id string    `json:"event_id" bson:"event_id"`
	Time     time.Time `json:"time" bson:"time"`
	//}}}
}

func Vote_event(event_id, user_id string, tm int64) error {
	//{{{{{{
	_, err := _col_vote.Upsert(bson.M{
		"event_id": event_id,
		"voter_id": user_id,
	}, bson.M{
		"$set": bson.M{"time": time.Unix(tm, 0)},
	})
	//}}}
	return err //}}}
}
func Votes_get_for_event(event_id string) ([]*Vote, error) {
	//{{{
	var votes = make([]*Vote, 0, 10)
	err := _col_vote.Find(bson.M{"event_id": event_id}).All(&votes)
	return votes, err //}}}
}
func Delete_vote(event_id, voter_id string) error {
	//{{{
	return _col_vote.Remove(bson.M{"event_id": event_id, "voter_id": voter_id})
	//}}}
}
