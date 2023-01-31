package db_helpers

import (
	"context"
	"errors"
	"fmt"
	"strings"
	"time"

	"github.com/ofjangra/gwonline/db"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

func CreateEntry_SP_Partner(vc, rto string, entry interface{}) error {

	ctx, cancel := context.WithTimeout(context.Background(), 100*time.Second)

	defer cancel()

	docCountWithThisEntry, docCountErr := db.SP_PartnerCollection.CountDocuments(ctx,
		bson.M{
			"vehicle_category": strings.ToUpper(vc),
			"rto":              strings.ToUpper(rto)})

	if docCountErr != nil {
		return errors.New("failed to create entry")
	}

	if docCountWithThisEntry > 0 {
		return errors.New("an entry with same vehicle category and rto already exists")
	}

	_, insertionErr := db.SP_PartnerCollection.InsertOne(ctx, entry)

	if insertionErr != nil {
		return errors.New("failed to create entry")
	}

	return nil
}

func GetAnEntry_SP_Partner(vehicleCategory string, rto string) *mongo.SingleResult {

	fmt.Println(vehicleCategory, rto)
	ctx, cancel := context.WithTimeout(context.Background(), 100*time.Second)

	defer cancel()

	result := db.SP_PartnerCollection.FindOne(ctx, bson.M{"vehicle_category": vehicleCategory, "rto": rto})

	return result

}

func UpdateEntry_SP_Partner(id string, update bson.M) error {
	ctx, cancel := context.WithTimeout(context.Background(), 100*time.Second)

	defer cancel()

	entryId, idErr := primitive.ObjectIDFromHex(id)

	if idErr != nil {
		return errors.New("failed to update entry")
	}

	_, updateErr := db.SP_PartnerCollection.UpdateByID(ctx, entryId, bson.M{"$set": update})

	if updateErr != nil {
		fmt.Println(updateErr)
		return errors.New("failed to update entry")
	}
	return nil

}
