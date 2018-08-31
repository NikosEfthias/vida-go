package api

import (
	"encoding/json"
	"net/http"

	"github.com/julienschmidt/httprouter"
	"github.com/mugsoft/vida/helpers"
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
