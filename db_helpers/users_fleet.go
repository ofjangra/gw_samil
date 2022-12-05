package db_helpers

import (
	"context"
	"errors"
	"fmt"
	"time"

	"github.com/ofjangra/gwonline/db"
	"github.com/ofjangra/gwonline/models"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func CreateFleetUser(user *models.FleetUser) error {
	ctx, cancel := context.WithTimeout(context.Background(), 100*time.Second)

	defer cancel()

	docCountWithThisEmail, docCountWithEmailErr := db.FleetUsersCollection.CountDocuments(ctx, bson.M{"email": user.Email})

	if docCountWithEmailErr != nil {
		return errors.New("failed to create account")
	}

	if docCountWithThisEmail > 0 {
		return errors.New("email already in use")
	}

	docCountWithThisPhone, docCountWithThisPhoneErr := db.FleetUsersCollection.CountDocuments(ctx, bson.M{"contact": user.Phone})

	if docCountWithThisPhoneErr != nil {
		return errors.New("failed to create account")
	}

	if docCountWithThisPhone > 0 {
		return errors.New("phone number already in use")
	}

	_, insertError := db.FleetUsersCollection.InsertOne(ctx, user)

	if insertError != nil {
		return errors.New("failed to create account")
	}

	return nil
}

func GetFleetUserByPhone(phone string) *mongo.SingleResult {
	ctx, cancel := context.WithTimeout(context.Background(), 60*time.Second)

	defer cancel()

	result := db.FleetUsersCollection.FindOne(ctx, bson.M{"phone": phone})

	return result
}

func GetFleetUserById(id string) (*mongo.SingleResult, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 60*time.Second)

	defer cancel()

	userId, idErr := primitive.ObjectIDFromHex(id)

	if idErr != nil {
		return nil, idErr
	}

	result := db.FleetUsersCollection.FindOne(ctx, bson.M{"_id": userId},
		options.FindOne().SetProjection(bson.M{
			"password": 0}))

	if result.Err() != nil {
		return nil, result.Err()
	}

	return result, nil
}

func GetAllFleetUsers() ([]primitive.M, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 100*time.Second)

	defer cancel()

	users := []primitive.M{}

	cur, curErr := db.FleetUsersCollection.Find(ctx, bson.M{}, options.Find().SetSort(bson.M{"created_on": -1}))

	if curErr != nil {
		return nil, errors.New("failed to fetch users")
	}

	for cur.Next(context.Background()) {
		user := bson.M{}

		err := cur.Decode(&user)

		if err != nil {
			return nil, errors.New("failed to fetch users")
		}

		users = append(users, user)
	}

	defer cur.Close(context.Background())

	return users, nil

}

func UpdateFleetUserProfile(id string, update bson.M) error {
	ctx, cancel := context.WithTimeout(context.Background(), 60*time.Second)

	defer cancel()
	fmt.Println(id)
	userId, idErr := primitive.ObjectIDFromHex(id)

	if idErr != nil {
		return errors.New("failed to update profile 1")
	}

	_, updateErr := db.FleetUsersCollection.UpdateByID(ctx, userId, bson.M{"$set": update})

	if updateErr != nil {
		fmt.Println(updateErr)
		return errors.New("failed to update profile 2")
	}
	return nil

}

func DeleteFleetUser(id string) error {
	ctx, cancel := context.WithTimeout(context.Background(), 60*time.Second)

	defer cancel()
	fmt.Println(id)
	userId, idErr := primitive.ObjectIDFromHex(id)

	if idErr != nil {
		return errors.New("failed to update profile 1")
	}

	_, deleteErr := db.FleetUsersCollection.DeleteOne(ctx, bson.M{"_id": userId})

	if deleteErr != nil {
		fmt.Println(deleteErr)
		return errors.New("failed to delete profile")
	}

	return nil
}
