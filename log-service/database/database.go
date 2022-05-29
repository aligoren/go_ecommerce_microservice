package database

import (
	"context"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"log"
	"os"
	"time"
)

var MongoClient *mongo.Client

type MongoConfig struct {
}

func ConnectoToMongoDb() {

	mongoUrl := os.Getenv("MONGO_URL")
	mongoUsername := os.Getenv("MONGO_USERNAME")
	mongoPassword := os.Getenv("MONGO_PASSWORD")

	mongoClientConfig := options.Client().ApplyURI(mongoUrl)

	mongoCredential := options.Credential{
		Username: mongoUsername,
		Password: mongoPassword,
	}

	mongoClientConfig.SetAuth(mongoCredential)

	connect, err := mongo.Connect(context.Background(), mongoClientConfig)

	if err != nil {
		log.Fatalf("Couldn't connect to the MongoDB server: %v", err)
	}

	MongoClient = connect

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	defer func() {
		if err = MongoClient.Disconnect(ctx); err != nil {
			log.Fatalf("Error while disconnecting  mongoDB: %v", err)
		}
	}()
}
