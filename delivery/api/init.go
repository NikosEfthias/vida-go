package api

import (
	"encoding/json"
	"net/http"

	"github.com/julienschmidt/httprouter"
	"gitlab.mugsoft.io/vida/api/go-api/helpers"
)

//Mount root handler
func Mount() http.Handler {
	var mux = httprouter.New()
	//__mux is the interface so it can be used for middlewares
	var __mux http.Handler
	mount_user(mux)
	mount__event(mux)
	mount__app(mux)
	//middleware mounting
	__mux = __middleware_headers_set(mux)
	return __mux
}

func __fv(r *http.Request, field__name string) string {
	return r.FormValue(field__name)
}

func __middleware_headers_set(next http.Handler) http.Handler {
	__hdrs := map[string]string{
		"Content-Type":                 "application/json",
		"Access-Control-Allow-Origin":  "*",
		"Access-Control-Allow-Headers": "GET, POST, PUT, PATCH, DELETE",
	}

	fn := func(w http.ResponseWriter, r *http.Request) {
		for k, v := range __hdrs {
			w.Header().Set(k, v)
		}
		next.ServeHTTP(w, r)
	}
	return http.HandlerFunc(fn)
}

func __parse__form(w http.ResponseWriter, r *http.Request) error {
	err := r.ParseForm()
	if nil != err {
		helpers.Log(helpers.ERR, err.Error())
		w.WriteHeader(http.StatusInternalServerError)
		json.NewEncoder(w).Encode(map[string]string{
			"error": err.Error(),
		})
		return err
	}
	return nil
}

func __parse__multipart__form(w http.ResponseWriter, r *http.Request, size int64) error {
	err := r.ParseMultipartForm(size)
	if nil != err {
		helpers.Log(helpers.ERR, err.Error())
		w.WriteHeader(http.StatusInternalServerError)
		json.NewEncoder(w).Encode(map[string]string{
			"error": err.Error(),
		})
		return err
	}
	return nil
}

func __respond__from__service(msg interface{}, err error, w http.ResponseWriter, r *http.Request) {
	if nil != err {
		w.WriteHeader(http.StatusForbidden)
		json.NewEncoder(w).Encode(map[string]string{
			"error": err.Error(),
		})
		return
	}
	json.NewEncoder(w).Encode(map[string]interface{}{
		"data": msg,
	})
}
