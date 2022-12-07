package db

import (
	"context"
	"fmt"
	"log"
	"os"
	"time"

	// "github.com/joho/godotenv"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func dbInstance() *mongo.Client {

	// envLoadErr := godotenv.Load(".env")

	// if envLoadErr != nil {
	// 	log.Fatal("Failed to load environment variables")
	// }

	var DBURI string = os.Getenv("DBURI")

	clientOptions := options.Client().ApplyURI(DBURI)

	ctx, cancel := context.WithTimeout(context.Background(), 120*time.Second)

	defer cancel()

	client, clientErr := mongo.Connect(ctx, clientOptions)

	if clientErr != nil {
		log.Fatal(clientErr)
		return nil
	}
	fmt.Println("Connected to db")
	return client
}
