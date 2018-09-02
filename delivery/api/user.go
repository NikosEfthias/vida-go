package api

import (
	"net/http"

	"github.com/julienschmidt/httprouter"
	"github.com/mugsoft/vida/services/user"
)

const PREFIX__USER = "/api/user"

func mount__user(mux *httprouter.Router) {

	mux.POST(PREFIX__USER+"/register",
		func(w http.ResponseWriter, r *http.Request, p httprouter.Params) {
			if nil != __parse__form(w, r) {
				return
			}
			var name, lastname, email, phone, password string
			name = __fv(r, "name")
			lastname = __fv(r, "lastname")
			email = __fv(r, "email")
			phone = __fv(r, "phone")
			password = __fv(r, "password")
			msg, err := user.Service_register(name, lastname, email, phone, password)
			__respond__from__service(msg, err, w, r)
		})

	mux.POST(PREFIX__USER+"/login",
		func(w http.ResponseWriter, r *http.Request, p httprouter.Params) {
			if nil != __parse__form(w, r) {
				return
			}
			var email, phone, password string
			email = __fv(r, "email")
			phone = __fv(r, "phone")
			password = __fv(r, "password")
			msg, err := user.Service_login(email, phone, password)
			__respond__from__service(msg, err, w, r)
		})

	mux.GET(PREFIX__USER+"/:token",
		func(w http.ResponseWriter, r *http.Request, p httprouter.Params) {
			msg, err := user.Service_get(p.ByName("token"))
			__respond__from__service(msg, err, w, r)
		})

	mux.POST(PREFIX__USER+"/update/:token/:field",
		func(w http.ResponseWriter, r *http.Request, p httprouter.Params) {
			if nil != __parse__form(w, r) {
				return
			}
			msg, err := user.Service_update(p.ByName("field"), p.ByName("token"), __fv(r, "value"))
			__respond__from__service(msg, err, w, r)
		})
	mux.POST(PREFIX__USER+"/pp", func(w http.ResponseWriter, r *http.Request, p httprouter.Params) {

	})

}
