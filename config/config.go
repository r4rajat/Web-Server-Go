package config

import (
	"context"
	"github.com/gorilla/mux"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"go.mongodb.org/mongo-driver/mongo/readpref"
	"log"
	"os"
)

// Models
type Book struct {
	ID        primitive.ObjectID `json:"_id,omitempty" bson:"_id,omitempty"`
	Name      string             `json:"name,omitempty" bson:"name,omitempty"`
	ISBN      int                `json:"isbn,omitempty" bson:"isbn,omitempty"`
	Publisher string             `json:"publisher,omitempty" bson:"publisher,omitempty"`
	Author    *Author            `json:"author,omitempty" bson:"author,omitempty"`
}
type Author struct {
	FirstName string `json:"first_name"`
	LastName  string `json:"last_name"`
}

// Environmental Variables
var DbHost = os.Getenv("DB_HOST")
var DbPort = os.Getenv("DB_PORT")
var AppHost = os.Getenv("APP_HOST")
var AppPort = os.Getenv("APP_PORT")
var DbName = os.Getenv("DB_NAME")
var DbCollection = os.Getenv("DB_COLLECTION")

// Gorilla Mux Router
var Router = mux.NewRouter()

// Get Mongo DB Connection
var Ctx = context.TODO()

func GetMongoDBDriver() *mongo.Client {
	DbAddress := ""
	if DbHost == "" || DbPort == "" {
		if DbHost == "" && DbPort == "" {
			log.Println("[DB_HOST]: Value Not Set. Setting it to Default (localhost)")
			log.Println("[DB_PORT]: Value Not Set. Setting it to Default (27017)")
			DbAddress = "mongodb://localhost:27017"
		}
		if DbHost == "" && DbPort != "" {
			log.Println("[DB_HOST]: Value Not Set. Setting it to Default (0.0.0.0)")
			DbAddress = "mongodb://localhost:" + DbPort
		}
		if DbPort == "" && DbHost != "" {
			log.Println("[DB_PORT]: Value Not Set. Setting it to Default (27017)")
			DbAddress = "mongodb://" + DbHost + ":" + "27017"
		}
	} else {
		DbAddress = "mongodb://" + DbHost + ":" + DbPort
	}
	clientOptions := options.Client().ApplyURI(DbAddress)
	client, err := mongo.NewClient(clientOptions)
	if err != nil {
		log.Fatal(err)
	}
	err = client.Connect(Ctx)
	if err != nil {
		log.Fatal(err)
	}
	err = client.Ping(context.Background(), readpref.Primary())
	if err != nil {
		log.Fatal("Couldn't connect to the database", err)
	}
	return client
}
