package db

import "go.mongodb.org/mongo-driver/mongo"

var Employee_samilCollection *mongo.Collection

var EmployeesCollection *mongo.Collection

func init() {

	client := dbInstance()

	EmployeesCollection = client.Database("main").Collection("employees")

	Employee_samilCollection = client.Database("samil_users").Collection("users")

}
