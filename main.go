package main

import (
	"fmt"
	"net/http"

	"gitlab.mugsoft.io/vida/api/go-api/config"
	"gitlab.mugsoft.io/vida/api/go-api/delivery/api"
)

func main() {
	fmt.Println("listening on", config.Get("LISTEN_ADDR"))
	http.ListenAndServe(config.Get("LISTEN_ADDR"), api.Mount())
}
