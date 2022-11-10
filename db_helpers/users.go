package db_helpers

import (
	"context"
	"errors"
	"log"
	"time"

	"github.com/ofjangra/gwonline/db"
	"github.com/ofjangra/gwonline/models"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

func Insertuser(user *models.User) error {
	ctx, cancel := context.WithTimeout(context.Background(), 60*time.Second)

	defer cancel()

	docCountWithThisUserID, userIDDocCountErr := db.UsersCollection.CountDocuments(ctx, bson.M{"userid": user.UserID})

	if userIDDocCountErr != nil {
		return errors.New("failed to create account")
	}

	if docCountWithThisUserID > 0 {
		return errors.New("id already exists")
	}

	docCountWithThisUserLink, userLinkDocCountError := db.UsersCollection.CountDocuments(ctx, bson.M{"userlink": user.UserLink})

	if userLinkDocCountError != nil {
		return errors.New("failed to create account")
	}

	if docCountWithThisUserLink > 0 {
		return errors.New("link already exists")
	}

	_, insertionErr := db.UsersCollection.InsertOne(ctx, user)

	if insertionErr != nil {
		log.Fatal(insertionErr)
		return errors.New("failed to create user")
	}

	return nil
}

func GetUserByUserID(id string) *mongo.SingleResult {

	ctx, cancel := context.WithTimeout(context.Background(), 60*time.Second)

	defer cancel()

	result := db.UsersCollection.FindOne(ctx, bson.M{"userid": id})

	return result
}
