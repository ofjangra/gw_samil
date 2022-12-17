package db

import "go.mongodb.org/mongo-driver/mongo"

var SamilUsersCollection *mongo.Collection

var EmployeesCollection *mongo.Collection

var VehiclesCollection *mongo.Collection

var FleetUsersCollection *mongo.Collection

var SP_AutomobileDealerCollection *mongo.Collection

func init() {

	client := dbInstance()

	EmployeesCollection = client.Database("main").Collection("employees")

	SamilUsersCollection = client.Database("users").Collection("samil_users")

	FleetUsersCollection = client.Database("users").Collection("fleet_users")

	VehiclesCollection = client.Database("vehicles").Collection("fleet_vehicles")

	SP_AutomobileDealerCollection = client.Database("main").Collection("services_pricing")

}
