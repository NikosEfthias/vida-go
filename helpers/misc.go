package helpers

import "fmt"

func Check_missing_fields(names, values []string) error {
	if len(names) != len(values) {
		return fmt.Errorf("missing fields")
	}
	for i := range names {
		if values[i] == "" {
			return fmt.Errorf("missing field: %s", names[i])
		}
	}
	return nil
}
