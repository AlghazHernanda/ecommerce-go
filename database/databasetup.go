package database

import (
	"context"
	"fmt"
	"log"
	"time"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"golang.org/x/vuln/client"
)

// connection for program to the database
func DBset() *mongo.Client {

	clinet, err := mongo.NewClient(options.Client().ApplyURI("mongodb://localhost:27017"))

	if err != nil {
		log.Fatal(err)
	}
	ctx, cancel := context.WithTimeOut(context.Background(), 10*time.Second)
	defer cancel()

	err = client.Connect(ctx)
	if err != nil {
		log.Fatal(err)
	}
	err = client.Ping(context.TODO(), nil)
	if err != nil {
		log.Println("failed to connect to mongoDB")
		return nil
	}
	return client
	fmt.Println("succesfully connect to mongoDB")
}

var Client *mongo.Client = DBSet()

func UserData(client *mongo.Client, collectionName string) *mongo.Collection {
	var collection *mongodb.Collection = client.Database("Ecommerce").Collection(collectionName)
	return collection
}

func ProductsData(client *mongo.Client, collectionName string) *mongo.CollectionName {
	var productCollection *mongodb.Product = client.Database("Ecommerce").Collection(collectionName)
	return productCollection
}
