package static

import (
	"encoding/json"
	"fmt"
	"net/http"
)

func Mount() http.Handler {
	mux := http.NewServeMux()
	mount__public(mux)
	return mux
}

func __respond_from_service(data []byte, mime string, err error, w http.ResponseWriter, r *http.Request) { //{{{
	if nil != err {
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusForbidden)
		json.NewEncoder(w).Encode(map[string]string{
			"error": err.Error(),
		})
		return
	}
	w.Header().Set("Content-Type", mime)
	w.Header().Set("Content-Length", fmt.Sprint(len(data)))
	//TODO:  Add gzip support
	w.WriteHeader(http.StatusOK)
	w.Write(data)

} //}}}
