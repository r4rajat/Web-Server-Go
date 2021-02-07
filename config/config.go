package config

import "github.com/gorilla/mux"

// Models
type Books struct {
	Name      string  `json:"name"`
	ISBN      int     `json:"isbn"`
	Publisher string  `json:"publisher"`
	Author    *Author `json:"author"`
}

type Author struct {
	FirstName string `json:"first_name"`
	LastName  string `json:"last_name"`
}

var Router = mux.NewRouter()
