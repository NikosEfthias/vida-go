package delivery

import (
	"net/http"

	"gitlab.mugsoft.io/vida/go-api/delivery/api"
	"gitlab.mugsoft.io/vida/go-api/delivery/static"
)

func Mount() http.Handler {
	mux := http.NewServeMux()
	mux.Handle("/api/", api.Mount())
	mux.Handle("/static/", static.Mount())
	return mux
}
