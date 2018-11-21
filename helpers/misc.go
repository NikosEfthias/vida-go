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

//Check_id_format checks the id format retreived from user
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
