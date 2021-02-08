package apis

import (
	"Web-Server-Go/config"
	"encoding/json"
	"github.com/bitly/go-simplejson"
	"go.mongodb.org/mongo-driver/bson"
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
		response.WriteHeader(http.StatusInternalServerError)
	}
	_, _ = response.Write(payload)

}

func GetBooksEndpoint(response http.ResponseWriter, _ *http.Request) {
	response.Header().Add("Content-Type", "application/json")
	var books []config.Book
	j := simplejson.New()
	client := config.GetMongoDBDriver()
	collection := client.Database(config.DbName).Collection(config.DbCollection)
	cursor, err := collection.Find(config.Ctx, bson.M{})
	if err != nil {
		response.WriteHeader(http.StatusInternalServerError)
		j.Set("error_message", err.Error())
		payload, err := j.MarshalJSON()
		if err != nil {
			_, _ = response.Write(payload)
		}
		return
	}
	defer cursor.Close(config.Ctx)
	for cursor.Next(config.Ctx) {
		var book config.Book
		cursor.Decode(&book)
		books = append(books, book)
	}
	if err := cursor.Err(); err != nil {
		response.WriteHeader(http.StatusInternalServerError)
		j.Set("error_message", err.Error())
		payload, err := j.MarshalJSON()
		if err != nil {
			_, _ = response.Write(payload)
		}
		return
	}
	j.Set("data", books)
	payload, err := j.MarshalJSON()
	if err != nil {
		response.WriteHeader(http.StatusInternalServerError)
		j.Set("error_message", err.Error())
	}
	_, _ = response.Write(payload)
}
