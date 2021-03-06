package apis

import (
	"Web-Server-Go/config"
	"encoding/json"
	"github.com/bitly/go-simplejson"
	"github.com/gorilla/mux"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
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
			return
		}
		_, _ = response.Write(payload)
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
			return
		}
		_, _ = response.Write(payload)
		return
	}
	cursor.Close(config.Ctx)
	j.Set("data", books)
	payload, err := j.MarshalJSON()
	if err != nil {
		response.WriteHeader(http.StatusInternalServerError)
		j.Set("error_message", err.Error())
	}
	_, _ = response.Write(payload)
}

func GetBookDetailsEndpoint(response http.ResponseWriter, request *http.Request) {
	response.Header().Add("Content-Type", "application/json")
	params := mux.Vars(request)
	id, _ := primitive.ObjectIDFromHex(params["_id"])
	var book config.Book
	j := simplejson.New()
	client := config.GetMongoDBDriver()
	collection := client.Database(config.DbName).Collection(config.DbCollection)
	err := collection.FindOne(config.Ctx, config.Book{ID: id}).Decode(&book)
	if err != nil {
		response.WriteHeader(http.StatusInternalServerError)
		j.Set("error_message", err.Error())
		payload, err := j.MarshalJSON()
		if err != nil {
			return
		}
		_, _ = response.Write(payload)
		return
	}
	j.Set("data", book)
	payload, err := j.MarshalJSON()
	if err != nil {
		response.WriteHeader(http.StatusInternalServerError)
		j.Set("error_message", err.Error())
	}
	_, _ = response.Write(payload)
}

func UpdateBookDetailsEndpoint(response http.ResponseWriter, request *http.Request) {
	response.Header().Add("Content-Type", "application/json")
	var book config.Book
	_ = json.NewDecoder(request.Body).Decode(&book)
	newData := bson.M{
		"$set": book,
	}
	params := mux.Vars(request)
	id, _ := primitive.ObjectIDFromHex(params["_id"])
	j := simplejson.New()
	client := config.GetMongoDBDriver()
	collection := client.Database(config.DbName).Collection(config.DbCollection)
	_, err := collection.UpdateOne(config.Ctx, config.Book{ID: id}, newData)
	if err != nil {
		response.WriteHeader(http.StatusInternalServerError)
		j.Set("error_message", err.Error())
		payload, err := j.MarshalJSON()
		if err != nil {
			return
		}
		_, _ = response.Write(payload)
		return
	}
	j.Set("message", "Values Updated")
	payload, err := j.MarshalJSON()
	if err != nil {
		response.WriteHeader(http.StatusInternalServerError)
		j.Set("error_message", err.Error())
	}
	_, _ = response.Write(payload)
}

func DeleteBookEndpoint(response http.ResponseWriter, request *http.Request) {
	response.Header().Add("Content-Type", "application/json")
	params := mux.Vars(request)
	id, _ := primitive.ObjectIDFromHex(params["_id"])
	j := simplejson.New()
	client := config.GetMongoDBDriver()
	collection := client.Database(config.DbName).Collection(config.DbCollection)
	_, err := collection.DeleteOne(config.Ctx, config.Book{ID: id})
	if err != nil {
		response.WriteHeader(http.StatusInternalServerError)
		j.Set("error_message", err.Error())
		payload, err := j.MarshalJSON()
		if err != nil {
			return
		}
		_, _ = response.Write(payload)
		return
	}
	j.Set("message", "Book with ID "+params["_id"]+"deleted")
	payload, err := j.MarshalJSON()
	if err != nil {
		response.WriteHeader(http.StatusInternalServerError)
		j.Set("error_message", err.Error())
	}
	_, _ = response.Write(payload)
}
