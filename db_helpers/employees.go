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

func GetEmployeeById(id string) (*mongo.SingleResult, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 60*time.Second)

	defer cancel()

	employeeId, idErr := primitive.ObjectIDFromHex(id)

	if idErr != nil {
		return nil, idErr
	}

	result := db.EmployeesCollection.FindOne(ctx, bson.M{"_id": employeeId},
		options.FindOne().SetProjection(bson.M{
			"password": 0}))

	if result.Err() != nil {
		return nil, result.Err()
	}

	return result, nil
}

func GetAllEmployees() ([]primitive.M, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 100*time.Second)

	defer cancel()

	employees := []primitive.M{}

	cur, curErr := db.EmployeesCollection.Find(ctx, bson.M{}, options.Find().SetSort(bson.M{"created_on": -1}))

	if curErr != nil {
		return nil, errors.New("failed to fetch employees")
	}

	for cur.Next(context.Background()) {
		employee := bson.M{}

		err := cur.Decode(&employee)

		if err != nil {
			return nil, errors.New("failed to fetch employees")
		}

		employees = append(employees, employee)
	}

	defer cur.Close(context.Background())

	return employees, nil

}

func UpdateEmployeeProfile(id string, update bson.M) error {
	ctx, cancel := context.WithTimeout(context.Background(), 60*time.Second)

	defer cancel()
	fmt.Println(id)
	employeeId, idErr := primitive.ObjectIDFromHex(id)

	if idErr != nil {
		return errors.New("failed to update profile")
	}

	_, updateErr := db.EmployeesCollection.UpdateByID(ctx, employeeId, bson.M{"$set": update})

	if updateErr != nil {
		fmt.Println(updateErr)
		return errors.New("failed to update profile")
	}
	return nil

}

func DeleteEmployee(id string) error {
	ctx, cancel := context.WithTimeout(context.Background(), 60*time.Second)

	defer cancel()
	fmt.Println(id)
	employeeId, idErr := primitive.ObjectIDFromHex(id)

	if idErr != nil {
		return errors.New("failed to update profile 1")
	}

	_, deleteErr := db.EmployeesCollection.DeleteOne(ctx, bson.M{"_id": employeeId})

	if deleteErr != nil {
		fmt.Println(deleteErr)
		return errors.New("failed to delete profile")
	}

	return nil
}
