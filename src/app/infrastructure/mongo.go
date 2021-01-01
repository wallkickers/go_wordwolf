package infrastructure

import (
	"fmt"
	"os"
	"github.com/sirupsen/logrus"

	// mongodb
	"context" // manage multiple requests
	"reflect" // get an object type
	// "time"
	// "go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

// Connect DB接続
func Connect() (db *mongo.Client, err error) {
	// Declare host and port options to pass to the Connect() method
	clientOptions := options.Client().ApplyURI("mongodb://"+os.Getenv("DB_USERNAME")+":"+os.Getenv("DB_USERPASS")+"@mongo")
	fmt.Println("clientOptions TYPE:", reflect.TypeOf(clientOptions), "\n")
	
	// Connect to the MongoDB and return Client instance
	db, err = mongo.Connect(context.TODO(), clientOptions)
	if err != nil {
		fmt.Println("mongo.Connect() ERROR:", err)
		logrus.Fatalf("Error connect DB: %v", err)
		os.Exit(1)
	}
	return db, err
}
