package database

import (
	"context"
	"log"
	"parties-app/backend/config"
	"time"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

var (
	Client           *mongo.Client
	UserCollection   *mongo.Collection
	PostCollection   *mongo.Collection
	EventsCollection *mongo.Collection
	Context          context.Context
)

func init() {
	Context, err := context.WithTimeout(context.Background(), 10*time.Second)
	if err != nil {
		panic(err)
	}

	Client, connectErr := mongo.Connect(Context, options.Client().ApplyURI(config.MONGO_URL))
	if connectErr != nil {
		panic(err)
	}

	log.Println("Connected to database!")

	UserCollection = Client.Database("PlusOne").Collection("users")
	PostCollection = Client.Database("PlusOne").Collection("posts")
	EventsCollection = Client.Database("PlusOne").Collection("events")
}
