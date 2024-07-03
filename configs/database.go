package configs

import (
	"context"
	"fmt"
	"log"
	"time"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func ConnectDb() *mongo.Client {
	client, err := mongo.Connect(
		context.TODO(),
		options.Client().ApplyURI(EnvMongoURI()),
	)

	if err != nil {
		log.Fatal(err)
	}

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	err = client.Ping(ctx, nil)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("Connected to MongoDB")
	return client
}

var DB *mongo.Client = ConnectDb()

func GetCollection(client *mongo.Client, collectionName string) *mongo.Collection {
	collection := client.Database("xAPI").Collection(collectionName)
	return collection
}
