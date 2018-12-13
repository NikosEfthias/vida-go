package static

import (
	"net/http"
	"strings"

	"gitlab.mugsoft.io/vida/go-api/services/storage"
)

func mount__public(mux *http.ServeMux) {
	mux.HandleFunc("/static/public/", func(w http.ResponseWriter, r *http.Request) {
		fpath := strings.Replace(r.URL.Path, "/static/public/", "", 1)
		d, m, err := storage.Service_public_files(fpath)
		__respond_from_service(d, m, err, w, r)
	})
}
