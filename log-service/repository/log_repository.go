package repository

import (
	"context"
	"github.com/aligoren/go_ecommerce_microservice/log-service/database"
	"github.com/aligoren/go_ecommerce_microservice/log-service/models"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"log"
	"os"
	"time"
)

func GetAllLogs() ([]*models.LogModel, error) {

	databaseName, _ := os.LookupEnv("MONGO_COLLECTION_NAME")
	collectionName, _ := os.LookupEnv("MONGO_COLLECTION_NAME")

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	collection := database.MongoClient.Database(databaseName).Collection(collectionName)

	logOptions := options.Find().SetSort(bson.D{{"createdAt", -1}})

	cursor, err := collection.Find(context.Background(), bson.D{}, logOptions)
	if err != nil {
		log.Fatalf("Error while getting logs %v", err)

		return nil, err
	}

	defer cursor.Close(ctx)

	var logs []*models.LogModel

	for cursor.Next(ctx) {
		var item models.LogModel

		err := cursor.Decode(&item)
		if err != nil {
			log.Fatalf("Error while decoding log data %v", err)

			return nil, err
		}

		logs = append(logs, &item)
	}

	return logs, nil
}

func GetLogByID(id string) (*models.LogModel, error) {
	databaseName, _ := os.LookupEnv("MONGO_COLLECTION_NAME")
	collectionName, _ := os.LookupEnv("MONGO_COLLECTION_NAME")

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	collection := database.MongoClient.Database(databaseName).Collection(collectionName)

	docID, err := primitive.ObjectIDFromHex(id)

	if err != nil {
		return nil, err
	}

	var logModel models.LogModel
	err = collection.FindOne(ctx, bson.M{"_id": docID}).Decode(&logModel)
	if err != nil {
		return nil, err
	}

	return &logModel, nil
}

func InsertLog(logModel models.LogModel) error {

	databaseName, _ := os.LookupEnv("MONGO_COLLECTION_NAME")
	collectionName, _ := os.LookupEnv("MONGO_COLLECTION_NAME")

	collection := database.MongoClient.Database(databaseName).Collection(collectionName)

	_, err := collection.InsertOne(context.Background(), logModel)

	if err != nil {
		log.Fatalf("Error while inserting data to mongoDB %v, db: %s", err, databaseName)
		return err
	}

	return nil
}

func UpdateLog(logModel models.LogModel) (*mongo.UpdateResult, error) {
	databaseName, _ := os.LookupEnv("MONGO_COLLECTION_NAME")
	collectionName, _ := os.LookupEnv("MONGO_COLLECTION_NAME")

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	collection := database.MongoClient.Database(databaseName).Collection(collectionName)

	docID, err := primitive.ObjectIDFromHex(logModel.ID)

	if err != nil {
		return nil, err
	}

	result, err := collection.UpdateOne(
		ctx,
		bson.M{"_id": docID},
		bson.D{{
			"$set",
			bson.D{
				{"name", logModel.Name},
				{"data", logModel.Data},
				{"updatedAt", time.Now()},
			},
		}})

	if err != nil {
		return nil, err
	}

	return result, nil
}

func DropCollection() error {
	databaseName, _ := os.LookupEnv("MONGO_COLLECTION_NAME")
	collectionName, _ := os.LookupEnv("MONGO_COLLECTION_NAME")

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	collection := database.MongoClient.Database(databaseName).Collection(collectionName)

	if err := collection.Drop(ctx); err != nil {
		log.Fatalf("Error while dropping collection %v", err)
		return err
	}

	return nil
}
