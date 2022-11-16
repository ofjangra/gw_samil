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

func InsertEmployee_Samil(employee *models.SamilEmployee) error {
	ctx, cancel := context.WithTimeout(context.Background(), 60*time.Second)

	defer cancel()

	docCountWithThisEmail, emailDocCountErr := db.Employee_samilCollection.CountDocuments(ctx, bson.M{"email": employee.Email})

	if emailDocCountErr != nil {
		return errors.New("failed to create account")
	}

	if docCountWithThisEmail > 0 {
		return errors.New("email already in use")
	}

	docCountWithThisUserLink, userLinkDocCountError := db.Employee_samilCollection.CountDocuments(ctx, bson.M{"userlink": employee.UserLink})

	if userLinkDocCountError != nil {
		return errors.New("failed to create account")
	}

	if docCountWithThisUserLink > 0 {
		return errors.New("link already exists")
	}

	docCountWithThisPhone, docCountWithThisPhoneErr := db.Employee_samilCollection.CountDocuments(ctx, bson.M{"contact": employee.Contact})

	if docCountWithThisPhoneErr != nil {
		return errors.New("failed to create employee p1")
	}

	if docCountWithThisPhone > 0 {
		return errors.New("phone number already in use")
	}

	_, insertionErr := db.Employee_samilCollection.InsertOne(ctx, employee)

	if insertionErr != nil {
		log.Fatal(insertionErr)
		return errors.New("failed to create employee")
	}

	return nil
}

func GetSamilEmployeeByEmail(email string) *mongo.SingleResult {

	ctx, cancel := context.WithTimeout(context.Background(), 60*time.Second)

	defer cancel()

	result := db.Employee_samilCollection.FindOne(ctx, bson.M{"email": email})

	return result
}
