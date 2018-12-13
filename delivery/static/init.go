package static

import (
	"encoding/json"
	"fmt"
	"net/http"
	"time"
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
	w.Header().Set("Cache-Control", "public, max-age=78840000")
	w.Header().Set("Expires", time.Now().Add(time.Hour*1314000).Format(time.RFC1123)) //5 years
	//TODO:  Add gzip support
	w.WriteHeader(http.StatusOK)
	w.Write(data)

} //}}}
