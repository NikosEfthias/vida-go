package helpers

import (
	"fmt"
	"regexp"
)

//Check_missing_fields gets string field names and values and checks if they are empty
func Check_missing_fields(names, values []string) error {
	//{{{
	if len(names) != len(values) {
		return fmt.Errorf("missing fields")
	}
	for i := range names {
		if values[i] == "" {
			return fmt.Errorf("missing field: %s", names[i])
		}
	}
	return nil //}}}
}

//Check_id_format checks the id format retrieved from user
func Check_id_format(id string) error {
	//{{{
	matched, err := regexp.MatchString("^[a-zA-Z0-9]{32}$", id)
	if nil != err {
		Log(ERR, "wtf this error is never supposed to be non-nil")
		return err
	}
	if !matched {
		return fmt.Errorf("invalid id format")
	}
	return nil //}}}
}

//Can_user_see_event decides whether the user can see the event
func Can_user_see_event(id string, event_participants []string, event_owner_id string) bool {
	//{{{
	if id == event_owner_id || Index_of_str(event_participants, id) > -1 {
		return true
	}
	return false
	//}}}
}

func Index_of_str(s []string, v string) int {
	//{{{
	for i, vv := range s {
		if vv == v {
			return i
		}
	}
	return -1 //}}}
}
