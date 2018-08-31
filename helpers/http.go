package helpers

import (
	"encoding/json"
	"net/http"
)

func Respond_json(w http.ResponseWriter, data interface{}) error {
	return json.NewEncoder(w).Encode(data)
}
