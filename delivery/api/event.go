package api

import (
	"encoding/json"
	"net/http"

	"github.com/julienschmidt/httprouter"
	"github.com/mugsoft/tools/bytesize"
	"github.com/mugsoft/vida/helpers"
	"github.com/mugsoft/vida/services/event"
)

const PREFIX_EVENT = "/api/event"

func mount__event(mux *httprouter.Router) {

	mux.POST(PREFIX_EVENT+"/create/:token", func(w http.ResponseWriter, r *http.Request, p httprouter.Params) {
		if nil != __parse__multipart__form(w, r, int64(bytesize.MB*5)) {
			return
		}
		f, _, err := r.FormFile("image")
		if nil != err {
			helpers.Log(helpers.ERR, err.Error())
			json.NewEncoder(w).Encode(map[string]string{"error": "cannot parse image"})
		}
		defer f.Close()
		msg, err := event.Service_create(p.ByName("token"),
			__fv(r, "title"),
			__fv(r, "location"),
			__fv(r, "start_date"),
			__fv(r, "end_date"),
			__fv(r, "details"),
			__fv(r, "max_num_guest"),
			__fv(r, "min_num_guest"),
			__fv(r, "cost"),
			__fv(r, "votable"),
			f)
		__respond__from__service(msg, err, w, r)
	})

}
