package db_helpers

import (
	"context"
	"errors"
	"fmt"
	"time"

	"github.com/ofjangra/gwonline/db"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

func CreateEntry_SP_Loading(passing float32, rto string, pricingEntry interface{}, collection *mongo.Collection) error {
	ctx, cancel := context.WithTimeout(context.Background(), 100*time.Second)

	defer cancel()

	docCountWithThisEntry, docCountErr := collection.CountDocuments(ctx, bson.M{"passing": passing, "rto": rto})

	if docCountErr != nil {
		return errors.New("failed to ceate entry")
	}

	if docCountWithThisEntry > 0 {
		return errors.New("an entry with same passing and rto alreday exists")
	}

	_, insertionErr := collection.InsertOne(ctx, pricingEntry)

	if insertionErr != nil {
		return errors.New("failed to create entry")
	}

	return nil
}

func CreateEntry_SP_Seating(seats int32, rto string, pricingEntry interface{}, collection *mongo.Collection) error {
	ctx, cancel := context.WithTimeout(context.Background(), 100*time.Second)

	defer cancel()

	docCountWithThisEntry, docCountErr := collection.CountDocuments(ctx, bson.M{"seating": seats, "rto": rto})

	if docCountErr != nil {
		return errors.New("failed to ceate entry")
	}

	if docCountWithThisEntry > 0 {
		return errors.New("an entry with same seating and rto alreday exists")
	}

	_, insertionErr := collection.InsertOne(ctx, pricingEntry)

	if insertionErr != nil {
		return errors.New("failed to create entry")
	}

	return nil
}

func GetAnEntry_SP_Loading(passing float32, rto string, collection *mongo.Collection) *mongo.SingleResult {
	ctx, cancel := context.WithTimeout(context.Background(), 100*time.Second)

	defer cancel()

	result := collection.FindOne(ctx, bson.M{"passing": passing, "rto": rto})

	return result

}

func GetAnEntry_SP_Seating(seating int32, rto string, collection *mongo.Collection) *mongo.SingleResult {
	ctx, cancel := context.WithTimeout(context.Background(), 100*time.Second)

	defer cancel()

	result := collection.FindOne(ctx, bson.M{"seating": seating, "rto": rto})

	return result

}

func UpdateEntry_SP_Loading(id string, update bson.M) error {
	ctx, cancel := context.WithTimeout(context.Background(), 100*time.Second)

	defer cancel()

	entryId, idErr := primitive.ObjectIDFromHex(id)

	if idErr != nil {
		return errors.New("failed to update entry")
	}

	_, updateErr := db.SP_LoadingCollection.UpdateByID(ctx, entryId, bson.M{"$set": update})

	if updateErr != nil {
		fmt.Println(updateErr)
		return errors.New("failed to update entry")
	}
	return nil

}
