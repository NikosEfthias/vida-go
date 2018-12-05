package api

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"

	"github.com/julienschmidt/httprouter"
	"github.com/mugsoft/tools/bytesize"
	"gitlab.mugsoft.io/vida/go-api/helpers"
	"gitlab.mugsoft.io/vida/go-api/services/event"
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
				return
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
	mux.POST(PREFIX_EVENT+"/update/:event_id/:field/:token",
		//{{{
		func(w http.ResponseWriter, r *http.Request, p httprouter.Params) {
			msg, err := event.Service_update(p.ByName("token"), p.ByName("event_id"), p.ByName("field"), __fv(r, "value"))
			__respond__from__service(msg, err, w, r)
		}) //}}}
	mux.PUT(PREFIX_EVENT+"/update/pp/:event_id/:token",
		//{{{
		func(w http.ResponseWriter, r *http.Request, p httprouter.Params) {
			if nil != __parse__multipart__form(w, r, int64(bytesize.MB*5)) {
				return
			}
			f, _, err := r.FormFile("file")
			if nil != err {
				helpers.Log(helpers.ERR, err.Error())
				json.NewEncoder(w).Encode(map[string]string{"error": "cannot parse file"})
			}
			defer f.Close()
			msg, err := event.Service_update_img(p.ByName("token"), p.ByName("event_id"), f)
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
	mux.GET(PREFIX_EVENT+"/byowner/:token/:page",
		//{{{
		func(w http.ResponseWriter, r *http.Request, p httprouter.Params) {
			msg, err := event.Service_get_by_owner(p.ByName("token"), p.ByName("page"), nil)
			__respond__from__service(msg, err, w, r)
		}) //}}}
	mux.POST(PREFIX_EVENT+"/byparticipant/:token/:page",
		//{{{
		func(w http.ResponseWriter, r *http.Request, p httprouter.Params) {
			var filters = map[string]interface{}{}
			//filter checks{{{
			if err := _fill_filter_int("status", __fv(r, "status"), filters); nil != err {
				__respond__from__service("", err, w, r)
			}
			//}}}
			msg, err := event.Service_get_by_participant(p.ByName("token"), p.ByName("page"), filters)
			__respond__from__service(msg, err, w, r)
		}) //}}}
	mux.POST(PREFIX_EVENT+"/invite/:event_id/:token",
		//{{{
		func(w http.ResponseWriter, r *http.Request, p httprouter.Params) {
			msg, err := event.Service_event_invite(p.ByName("token"), p.ByName("event_id"), __fv(r, "invitees"))
			__respond__from__service(msg, err, w, r)
		}) //}}}
	mux.GET(PREFIX_EVENT+"/accept/:event_id/:token",
		//{{{
		func(w http.ResponseWriter, r *http.Request, p httprouter.Params) {
			msg, err := event.Service_event_accept(p.ByName("token"), p.ByName("event_id"))
			__respond__from__service(msg, err, w, r)
		}) //}}}
	mux.GET(PREFIX_EVENT+"/decline/:event_id/:token",
		//{{{
		func(w http.ResponseWriter, r *http.Request, p httprouter.Params) {
			msg, err := event.Service_event_decline(p.ByName("token"), p.ByName("event_id"))
			__respond__from__service(msg, err, w, r)
		}) //}}}
}
func _fill_filter_int(name, value string, store map[string]interface{}) error {
	//{{{
	if nil == store {
		fmt.Errorf("nil store system error")
	}
	if value == "" {
		return nil
	}
	i, err := strconv.Atoi(value)
	if nil != err {
		return err
	}
	store[name] = i
	return nil //}}}
}
