package db_helpers

import (
	"context"
	"errors"
	"log"
	"time"

	"github.com/ofjangra/gwonline/db"
	"github.com/ofjangra/gwonline/models"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

func InsertUser_Samil(user *models.SamilUser) error {
	ctx, cancel := context.WithTimeout(context.Background(), 60*time.Second)

	defer cancel()

	docCountWithThisEmail, emailDocCountErr := db.SamilUsersCollection.CountDocuments(ctx, bson.M{"email": user.Email})

	if emailDocCountErr != nil {
		return errors.New("failed to create account")
	}

	if docCountWithThisEmail > 0 {
		return errors.New("email already in use")
	}

	// docCountWithThisUserLink, userLinkDocCountError := db.Employee_samilCollection.CountDocuments(ctx, bson.M{"userlink": employee.UserLink})

	// if userLinkDocCountError != nil {
	// 	return errors.New("failed to create account")
	// }

	// if docCountWithThisUserLink > 0 {
	// 	return errors.New("link already exists")
	// }

	_, insertionErr := db.SamilUsersCollection.InsertOne(ctx, user)

	if insertionErr != nil {
		log.Fatal(insertionErr)
		return errors.New("failed to create user")
	}

	return nil
}

func GetSamilUserByEmail(email string) *mongo.SingleResult {

	ctx, cancel := context.WithTimeout(context.Background(), 60*time.Second)

	defer cancel()

	result := db.SamilUsersCollection.FindOne(ctx, bson.M{"email": email})

	return result
}

func GetSamilUserById(id string) *mongo.SingleResult {

	ctx, cancel := context.WithTimeout(context.Background(), 60*time.Second)

	defer cancel()

	userId, idErr := primitive.ObjectIDFromHex(id)

	if idErr != nil {
		return nil
	}

	result := db.SamilUsersCollection.FindOne(ctx, bson.M{"_id": userId})

	return result
}

func GetAllSamilUsers() ([]primitive.M, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 60*time.Second)

	defer cancel()

	users := []primitive.M{}

	cur, curErr := db.SamilUsersCollection.Find(ctx, bson.M{})

	if curErr != nil {
		return nil, curErr
	}

	for cur.Next(context.Background()) {
		user := bson.M{}

		err := cur.Decode(&user)

		if err != nil {
			return nil, err
		}

		users = append(users, user)
	}

	defer cur.Close(context.Background())

	return users, nil
}
