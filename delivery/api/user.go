package api

import (
	//{{{
	"encoding/json"
	"net/http"

	"github.com/julienschmidt/httprouter"
	"github.com/mugsoft/tools/bytesize"
	"gitlab.mugsoft.io/vida/go-api/helpers"
	"gitlab.mugsoft.io/vida/go-api/services/user"
	//}}}
)

const PREFIX_USER = "/api/user"

func mount_user(mux *httprouter.Router) {

	mux.POST(PREFIX_USER+"/register",
		//{{{
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
		},
		//}}}
	)

	mux.POST(PREFIX_USER+"/login",
		//{{{
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
		},
	//}}}
	)
	mux.POST(PREFIX_USER+"/forgot",
		//{{{
		func(w http.ResponseWriter, r *http.Request, p httprouter.Params) {
			if nil != __parse__form(w, r) {
				return
			}
			var email string
			email = __fv(r, "email")
			msg, err := user.Service_forgot_password(email)
			__respond__from__service(msg, err, w, r)
		},
	//}}}
	)
	mux.GET(PREFIX_USER+"/:token",
		//{{{
		func(w http.ResponseWriter, r *http.Request, p httprouter.Params) {
			msg, err := user.Service_get(p.ByName("token"))
			__respond__from__service(msg, err, w, r)
		},
	//}}}
	)

	mux.POST(PREFIX_USER+"/update/:token/:field",
		//{{{
		func(w http.ResponseWriter, r *http.Request, p httprouter.Params) {
			if nil != __parse__form(w, r) {
				return
			}
			msg, err := user.Service_update(p.ByName("field"), p.ByName("token"), __fv(r, "value"))
			__respond__from__service(msg, err, w, r)
		},
	//}}}
	)

	mux.POST(PREFIX_USER+"/pp/:token",
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
			msg, err := user.Service_profile_pic(p.ByName("token"), f)
			__respond__from__service(msg, err, w, r)
		},
	//}}}
	)
	mux.DELETE(PREFIX_USER+"/logout/:token",
		//{{{
		func(w http.ResponseWriter, r *http.Request, p httprouter.Params) {
			msg, err := user.Service_logout(p.ByName("token"))
			__respond__from__service(msg, err, w, r)
		},
	//}}}
	)

}
