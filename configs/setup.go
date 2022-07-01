package configs

import (
	"context"
	"fmt"
	"log"
	"time"

	"flaq.club/backend/configs/vars"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func ConnectDB() *mongo.Client {
	client, err := mongo.NewClient(options.Client().ApplyURI(LoadEnv(vars.MONGOURI)))

	if err != nil {
		log.Fatal("Could not initialize client: ", err)
	}

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	if err := client.Connect(ctx); err != nil {
		log.Fatal("Could not connect to client: ", err)
	}

	if err := client.Ping(ctx, nil); err != nil {
		log.Fatal("Could not ping the database: ", err)
	}

	fmt.Println("Connected to MongoDB")

	return client
}

var DB *mongo.Client = ConnectDB()

func GetCollection(client *mongo.Client, collectionName string) *mongo.Collection {
	collection := client.Database("flaq").Collection(collectionName)
	return collection
}
