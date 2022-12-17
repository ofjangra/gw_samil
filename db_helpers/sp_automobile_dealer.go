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
)

func CreateEntry_SP_AutomobileDealer(pricingEntry *models.SP_AutomobileDealer) error {
	ctx, cancel := context.WithTimeout(context.Background(), 100*time.Second)

	defer cancel()

	docCountWithThisEntry, docCountErr := db.SP_AutomobileDealerCollection.CountDocuments(ctx, bson.M{"passing": pricingEntry.Passing, "rto": pricingEntry.RTO})

	if docCountErr != nil {
		return errors.New("failed to ceate entry")
	}

	if docCountWithThisEntry > 0 {
		return errors.New("an entry with same passing and rto alreday exists")
	}

	_, insertionErr := db.SP_AutomobileDealerCollection.InsertOne(ctx, pricingEntry)

	if insertionErr != nil {
		return errors.New("failed to create entry")
	}

	return nil
}

func GetAnEntry_SP_AutomobileDealer(passing float64, rto string) *mongo.SingleResult {
	ctx, cancel := context.WithTimeout(context.Background(), 100*time.Second)

	defer cancel()

	result := db.SP_AutomobileDealerCollection.FindOne(ctx, bson.M{"passing": passing, "rto": rto})

	return result

}

func UpdateEntry_SP_AutomobileDealer(id string, update bson.M) error {
	ctx, cancel := context.WithTimeout(context.Background(), 100*time.Second)

	defer cancel()

	entryId, idErr := primitive.ObjectIDFromHex(id)

	if idErr != nil {
		return errors.New("failed to update entry")
	}

	_, updateErr := db.SP_AutomobileDealerCollection.UpdateByID(ctx, entryId, bson.M{"$set": update})

	if updateErr != nil {
		fmt.Println(updateErr)
		return errors.New("failed to update entry")
	}
	return nil

}
