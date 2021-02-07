package apis

import (
	"Web-Server-Go/config"
	"encoding/json"
	"github.com/bitly/go-simplejson"
	"log"
	"net/http"
)

func CreateBookEndpoint(response http.ResponseWriter, request *http.Request) {
	response.Header().Add("Content-Type", "application/json")
	var book config.Book
	_ = json.NewDecoder(request.Body).Decode(&book)
	client := config.GetMongoDBDriver()
	collection := client.Database(config.DbName).Collection(config.DbCollection)
	result, _ := collection.InsertOne(config.Ctx, book)
	j := simplejson.New()
	j.Set("_id", result.InsertedID)
	j.Set("message", "Book Created")
	payload, err := j.MarshalJSON()
	if err != nil {
		log.Println(err)
	}
	_, _ = response.Write(payload)

}
