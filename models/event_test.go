package models

import "testing"

func Test_Event_delete(t *testing.T) {
	//before hook{{{
	event := &Event{Owner: "me"}
	err := Event_new(event)
	if nil != err {
		t.Fatal(err)
	}
	// }}}
	//invalid{{{
	invalidIds := []string{
		"",
		"/testing/",
		"*asterix",
		",comma",
		"b161c228ccdfc2f45ac5ba33a0964f3/",
	}
	_ = invalidIds
	for _, id := range invalidIds {
		err := Event_delete(id)
		if nil == err {
			t.Fatal("expected an error got nil")
		}
	}
	//}}}
	//valid {{{

	err = Event_delete(event.Id)
	if nil != err {
		t.Fatalf("expected nil got: '%s'", err.Error())
	}
	//}}}

}
