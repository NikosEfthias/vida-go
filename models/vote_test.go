package models

import (
	"testing"

	"gopkg.in/mgo.v2/bson"
)

func _test_vote_cleanup(owners ...string) {
	if len(owners) == 0 {
		return
	}
	_col_vote.RemoveAll(bson.M{"voter": bson.M{"$in": owners}})
}
func TestVote(t *testing.T) {
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
		},
	}
	for _, c := range cases {
		err := Vote_event(c.event_id, __voter_id, 0)
		errored := nil != err
		if c.error_expected != errored {
			t.Fatalf("expected errored status to be %v found error= (%v)"+
				"\n Description: %s",
				c.error_expected, err, c.desciption)
		}
	}
}
