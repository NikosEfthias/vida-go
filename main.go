package main

import (
	"fmt"
	"net/http"

	"github.com/mugsoft/vida/config"
	"github.com/mugsoft/vida/delivery/api"
)

func main() {
	fmt.Println("listening on", config.Get("LISTEN_ADDR"))
	http.ListenAndServe(config.Get("LISTEN_ADDR"), api.Mount())
}
