package api

import (
	"net/http"

	"github.com/julienschmidt/httprouter"
)

//Mount root handler
func Mount() http.Handler {
	var mux = httprouter.New()
	mount__user(mux)
	return mux
}
func __fv(r *http.Request, field__name string) string {
	return r.FormValue(field__name)
}
