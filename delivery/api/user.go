package api

import (
	"encoding/json"
	"net/http"

	"github.com/julienschmidt/httprouter"
	"github.com/mugsoft/vida/helpers"
	"github.com/mugsoft/vida/services/user"
)

const PREFIX__USER = "/api/user"

func mount__user(mux *httprouter.Router) {

	mux.POST(PREFIX__USER+"/register",
		func(w http.ResponseWriter, r *http.Request, p httprouter.Params) {
			err := r.ParseForm()
			if nil != err {
				helpers.Log(helpers.ERR, err.Error())
				w.WriteHeader(http.StatusInternalServerError)
				json.NewEncoder(w).Encode(map[string]string{
					"error": err.Error(),
				})
				return
			}
			var name, lastname, email, phone, password string
			name = __fv(r, "name")
			lastname = __fv(r, "lastname")
			email = __fv(r, "email")
			phone = __fv(r, "phone")
			password = __fv(r, "password")
			msg, err := user.Service_register(name, lastname, email, phone, password)
			if nil != err {
				w.WriteHeader(http.StatusForbidden)
				json.NewEncoder(w).Encode(map[string]string{
					"error": err.Error(),
				})
				return
			}
			json.NewEncoder(w).Encode(map[string]string{
				"data": msg,
			})
		})

}
