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
	} else {
		AppAddress = config.AppHost + ":" + config.AppPort
	}
}

func checkDBNameCollection() {
	if config.DbName == "" || config.DbCollection == "" {
		if config.DbName == "" {
			log.Println("[DB_NAME]: Value Not Set")
		}
		if config.DbCollection == "" {
			log.Println("[DB_COLLECTION]: Value Not Set")
		}
		log.Fatal("Set Above Environmental Variable's Value First")
	}
}

func main() {
	setAppAddress()
	checkDBNameCollection()

	log.Println("Server Running on : ", AppAddress)

	// API Endpoints
	config.Router.HandleFunc("/", apis.HomeHandler).Methods("GET")
	config.Router.HandleFunc("/api/book/create", apis.CreateBookEndpoint).Methods("POST")
	config.Router.HandleFunc("/api/books/view", apis.GetBooksEndpoint).Methods("GET")
	config.Router.HandleFunc("/api/book/view", apis.GetBookDetailsEndpoint).Methods("GET")
	config.Router.HandleFunc("/api/book/update", apis.UpdateBookDetailsEndpoint).Methods("PATCH")
	config.Router.HandleFunc("/api/book/delete", apis.DeleteBookEndpoint).Methods("DELETE")
	_ = http.ListenAndServe(AppAddress, config.Router)
}
