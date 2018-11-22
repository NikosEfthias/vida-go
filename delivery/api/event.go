package api

import (
	"encoding/json"
	"net/http"

	"github.com/julienschmidt/httprouter"
	"github.com/mugsoft/tools/bytesize"
	"gitlab.mugsoft.io/vida/api/go-api/helpers"
	"gitlab.mugsoft.io/vida/api/go-api/services/event"
)

const PREFIX_EVENT = "/api/event"

func mount__event(mux *httprouter.Router) {
	mux.POST(PREFIX_EVENT+"/create/:token",
		//{{{
		func(w http.ResponseWriter, r *http.Request, p httprouter.Params) {
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
		}) //}}}
	mux.GET(PREFIX_EVENT+"/delete/:id/:token",
		//{{{
		func(w http.ResponseWriter, r *http.Request, p httprouter.Params) {
			msg, err := event.Service_delete(p.ByName("token"), p.ByName("id"))
			__respond__from__service(msg, err, w, r)

		}) //}}}
	mux.GET(PREFIX_EVENT+"/byid/:id/:token",
		//{{{
		func(w http.ResponseWriter, r *http.Request, p httprouter.Params) {
			msg, err := event.Service_get_by_id(p.ByName("token"), p.ByName("id"), nil)
			__respond__from__service(msg, err, w, r)

		}) //}}}
	mux.GET(PREFIX_EVENT+"/byowner/:token/:start/:end",
		//{{{
		func(w http.ResponseWriter, r *http.Request, p httprouter.Params) {
			msg, err := event.Service_get_by_owner(p.ByName("token"), p.ByName("start"), p.ByName("end"), nil)
			__respond__from__service(msg, err, w, r)
		}) //}}}
	mux.GET(PREFIX_EVENT+"/byparticipant/:token/:start/:end",
		//{{{
		func(w http.ResponseWriter, r *http.Request, p httprouter.Params) {
			msg, err := event.Service_get_by_participant(p.ByName("token"), p.ByName("start"), p.ByName("end"), nil)
			__respond__from__service(msg, err, w, r)
		}) //}}}
}
