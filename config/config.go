package config

import (
	"context"
	"github.com/gorilla/mux"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"log"
	"os"
	"time"
)

// Models
type Books struct {
	ID        primitive.ObjectID `json:"_id,omitempty" bson:"_id,omitempty"`
	Name      string             `json:"name"`
	ISBN      int                `json:"isbn"`
	Publisher string             `json:"publisher"`
	Author    *Author            `json:"author"`
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

// Gorilla Mux Router
var Router = mux.NewRouter()

// Get Mongo DB Connection
var ctx, cancel = context.WithTimeout(context.Background(), 10*time.Second)

func getMongoDBDriver() *mongo.Client {
	DbAddress := ""
	if DbHost == "" || DbPort == "" {
		if DbHost == "" && DbPort == "" {
			log.Println("[DB_HOST]: Value Not Set. Setting it to Default (0.0.0.0)")
			log.Println("[DB_PORT]: Value Not Set. Setting it to Default (27017)")
			DbAddress = "mongodb://0.0.0.0:27017"
		}
		if DbHost == "" && DbPort != "" {
			log.Println("[DB_HOST]: Value Not Set. Setting it to Default (0.0.0.0)")
			DbAddress = "mongodb://0.0.0.0:" + DbPort
		}
		if DbPort == "" && DbHost != "" {
			log.Println("[DB_PORT]: Value Not Set. Setting it to Default (27017)")
			DbAddress = "mongodb://" + DbHost + ":" + "27017"
		}
	}
	var client *mongo.Client
	defer cancel()
	client, _ = mongo.Connect(ctx, options.Client().ApplyURI(DbAddress))
	return client
}
