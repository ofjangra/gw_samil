package db_helpers

import (
	"context"
	"errors"
	"time"

	"github.com/ofjangra/gwonline/db"
	"github.com/ofjangra/gwonline/models"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func InsertAVehicle(vehicle *models.Vehicle) error {
	ctx, cancel := context.WithTimeout(context.Background(), time.Second*60)

	defer cancel()

	docCountWithThisRegNumber, RegNumberDocCountErr := db.VehiclesCollection.CountDocuments(ctx, bson.M{"registeration_number": vehicle.RegisterationNumber})

	if RegNumberDocCountErr != nil {
		return errors.New("failed to insert vehicle")
	}

	if docCountWithThisRegNumber > 0 {
		return errors.New("registeration number already exists")
	}

	_, insertionErr := db.VehiclesCollection.InsertOne(ctx, vehicle)

	if insertionErr != nil {
		return errors.New("failed to create user")
	}

	return nil

}

func GetAVehicleById(id string) (*mongo.SingleResult, error) {

	ctx, cancel := context.WithTimeout(context.Background(), 60*time.Second)

	defer cancel()

	vehicleId, idErr := primitive.ObjectIDFromHex(id)

	if idErr != nil {
		return nil, errors.New("failed to fetch vehicle details")
	}

	result := db.VehiclesCollection.FindOne(ctx, bson.M{"_id": vehicleId})

	return result, nil
}

func GetAllVehicles() ([]primitive.M, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 100*time.Second)

	defer cancel()

	vehicles := []primitive.M{}

	cur, curErr := db.VehiclesCollection.Find(ctx, bson.M{}, options.Find().SetSort(bson.M{"created_on": -1}))

	if curErr != nil {
		return nil, curErr
	}

	for cur.Next(context.Background()) {
		vehicle := bson.M{}

		err := cur.Decode(&vehicle)

		if err != nil {
			return nil, err
		}

		vehicles = append(vehicles, vehicle)
	}

	defer cur.Close(context.Background())

	return vehicles, nil
}
