package config

import (
	"encoding/json"
	"os"
)

var conf = map[string]string{
	"DB_ADDR":          "mongodb://localhost:27017",
	"DB":               "vida",
	"LISTEN_ADDR":      ":8080",
	"APP_EMAIL_ADDR":   "info@vidaevents.org",
	"SMTP_ADDR":        "smtp.gmail.com:587",
	"APP_EMAIL_PASSWD": "<secret>",
	"APP_BASE_URL":     "https://devo.vidavidavida.com",
}

func init() {
	f, err := os.Open("conf.json")
	if nil != err {
		switch {
		case os.IsNotExist(err):
			d, _ := json.MarshalIndent(conf, "", "	")
			f, err = os.OpenFile("conf.json", os.O_CREATE|os.O_WRONLY, 0644)
			if nil != err {
				panic(err)
			}
			_, _ = f.Write(d)
			_ = f.Close()
		default:
			panic(err)
		}
		return
	}
	defer f.Close()
	err = json.NewDecoder(f).Decode(&conf)
	if nil != err {
		panic(err)
	}
}

//Get config
func Get(k string) string {
	return conf[k]
}
