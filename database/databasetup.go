package database

import (
	// "context"
	// "fmt"
	// "log"
	// "time"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

// connection for program to the database
func DBset() *mongo.Client {

	clinet, err := mongo.NewClient(options.Client().ApplyURI("mongodb://localhost:27017"))
}

func UserData(client *mongo.Client, collectionName string) *mongo.Collection {

}

func ProductsData(client *mongo.Client, collectionName string) *mongo.CollectionName {

}
