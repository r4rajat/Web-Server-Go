package main

import (
	"Web-Server-Go/apis"
	"Web-Server-Go/config"
	"log"
	"net/http"
)

var AppAddress = ""

func setAppAddress() {
	if config.AppHost == "" || config.AppPort == "" {
		if config.AppHost == "" && config.AppPort == "" {
			log.Println("[APP_HOST]: Value Not Set. Setting it to Default (0.0.0.0)")
			log.Println("[APP_PORT]: Value Not Set. Setting it to Default (5000)")
			AppAddress = "0.0.0.0:5000"
		}
		if config.AppHost == "" && config.AppPort != "" {
			log.Println("[APP_HOST]: Value Not Set. Setting it to Default (0.0.0.0)")
			AppAddress = "0.0.0.0:" + config.AppPort
		}
		if config.AppPort == "" && config.AppHost != "" {
			log.Println("[APP_PORT]: Value Not Set. Setting it to Default (5000)")
			AppAddress = config.AppHost + ":5000"
		}
	}
}

func main() {
	setAppAddress()
	log.Println("Server Running on : ", AppAddress)
	config.Router.HandleFunc("/", apis.HomeHandler)
	_ = http.ListenAndServe(AppAddress, config.Router)
}
