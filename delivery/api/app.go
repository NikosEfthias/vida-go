package api

import (
	"net/http"
	"strings"

	"github.com/julienschmidt/httprouter"
	"gitlab.mugsoft.io/vida/go-api/services/app"
)

const PREFIX_APP = "/api/app"

func mount__app(mux *httprouter.Router) {
	mux.POST(PREFIX_APP+"/invite/:token",
		//{{{
		func(w http.ResponseWriter, r *http.Request, p httprouter.Params) {
			if nil != __parse__form(w, r) {
				return
			}
			var invitees interface{} = __fv(r, "invitees")
			invitees = strings.Split(invitees.(string), ":")
			resp, err := app.Service_invite_people(p.ByName("token"), invitees.([]string))
			__respond__from__service(resp, err, w, r)
			//}}}
		})
}
