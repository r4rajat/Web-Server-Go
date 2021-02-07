package main

import (
	"Web-Server-Go/apis"
	"Web-Server-Go/config"
	"net/http"
)

func main() {
	config.Router.HandleFunc("/", apis.HomeHandler)
	_ = http.ListenAndServe(":8080", config.Router)
}
