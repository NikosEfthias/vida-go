package models

import (
	"testing"

	"gopkg.in/mgo.v2/bson"
)

func _test_vote_cleanup(owners ...string) {
	if len(owners) == 0 {
		return
	}
	_col_vote.RemoveAll(bson.M{"voter_id": bson.M{"$in": owners}})
}
func TestVote(t *testing.T) {
	//cases setup{{{
	const __voter_id = "__testing__"
	defer _test_vote_cleanup(__voter_id)

	type _CASE struct {
		error_expected bool
		desciption     string
		event_id       string
	}
	cases := []_CASE{
		_CASE{
			false,
			"voting an event for the first time must succeed",
			"1",
		}}
	//}}}
	//execute{{{
	for _, c := range cases {
		err := Vote_event(c.event_id, __voter_id, 0)
		errored := nil != err
		if c.error_expected != errored {
			t.Fatalf("expected errored status to be %v found error= (%v)"+
				"\n Description: %s",
				c.error_expected, err, c.desciption)
		}
	} //}}}
}
func TestVotes_get_for_event(t *testing.T) {
	//{{{
	defer _test_vote_cleanup("__test__")
	Vote_event("1", "__test__", 0)
	events, err := Votes_get_for_event("1")
	if nil != err {
		t.Fatalf("expected no error found ( %v )", err)
	}
	if len(events) != 1 {
		t.Fatalf("expected event count to be one but found %d", len(events))
	}
	if events[0].Event_id != "1" || events[0].Voter_id != "__test__" {
		t.Fatalf("function doesnt properly work expected event owner and event id to be __test__ , 1 found ( %s,%s )", events[0].Voter_id, events[0].Event_id)
	}
	//}}}
}
func TestDelete_vote(t *testing.T) {
	//{{{
	defer _test_vote_cleanup("__test__")
	err := Vote_event("1", "__test__", 0)
	if nil != err {
		t.Fatalf("expected error to be nil found %v", err)
	}
	err = Delete_vote("1", "__test__")
	if nil != err {
		t.Fatalf("expected error to be nil while deleting an existing vote bot found error (%v)", err)
	}
	err = Delete_vote("1", "__test__")
	if nil == err {
		t.Fatalf("deleting a non existing event must fail yet theres no error")
	}
	//}}}
}
