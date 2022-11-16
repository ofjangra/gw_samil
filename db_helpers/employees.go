package db_helpers

import (
	"context"
	"errors"
	"time"

	"github.com/ofjangra/gwonline/db"
	"github.com/ofjangra/gwonline/models"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

func InsertEmployee(employee *models.Employee) error {
	ctx, cancel := context.WithTimeout(context.Background(), 30*time.Second)

	defer cancel()

	docCountWithThisEmail, docCountWithEmailErr := db.EmployeesCollection.CountDocuments(ctx, bson.M{"email": employee.Email})

	if docCountWithEmailErr != nil {
		return errors.New("failed to create employee e1")
	}

	if docCountWithThisEmail > 0 {
		return errors.New("email already in use")
	}

	docCountWithThisPhone, docCountWithThisPhoneErr := db.EmployeesCollection.CountDocuments(ctx, bson.M{"contact": employee.Contact})

	if docCountWithThisPhoneErr != nil {
		return errors.New("failed to create employee p1")
	}

	if docCountWithThisPhone > 0 {
		return errors.New("phone number already in use")
	}

	_, insertError := db.EmployeesCollection.InsertOne(ctx, employee)

	if insertError != nil {
		return errors.New("failed to create employee d1")
	}

	return nil
}

func GetEmployeeByEmail(email string) *mongo.SingleResult {
	ctx, cancel := context.WithTimeout(context.Background(), 60*time.Second)

	defer cancel()

	result := db.EmployeesCollection.FindOne(ctx, bson.M{"email": email})

	return result
}
