package main

import (
	"fmt"
	"net/http"

	"gitlab.mugsoft.io/vida/go-api/config"
	"gitlab.mugsoft.io/vida/go-api/delivery"
)

func main() {
	fmt.Println("listening on", config.Get("LISTEN_ADDR"))
	http.ListenAndServe(config.Get("LISTEN_ADDR"), delivery.Mount())
}
