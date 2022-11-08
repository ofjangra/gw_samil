package db

import "go.mongodb.org/mongo-driver/mongo"

const DBNAME = "user_db"

var UsersCollection *mongo.Collection

func init() {

	client := dbInstance()

	UsersCollection = client.Database(DBNAME).Collection("users")

}
