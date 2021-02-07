package apis

import (
	"github.com/bitly/go-simplejson"
	"log"
	"net/http"
)

func HomeHandler(w http.ResponseWriter, _ *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	json := simplejson.New()
	json.Set("Hello", "World")
	payload, err := json.MarshalJSON()
	if err != nil {
		log.Println(err)
	}
	_, _ = w.Write(payload)
}
