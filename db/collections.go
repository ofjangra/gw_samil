package db

import "go.mongodb.org/mongo-driver/mongo"

var SamilUsersCollection *mongo.Collection

var EmployeesCollection *mongo.Collection

var VehiclesCollection *mongo.Collection

var FleetUsersCollection *mongo.Collection

var SP_LoadingCollection *mongo.Collection

var SP_PartnerCollection *mongo.Collection

var SP_BusCollection *mongo.Collection

var SP_MotorCab *mongo.Collection

var SP_ConstructionEqp *mongo.Collection

var SP_FarmEqp *mongo.Collection

func init() {

	client := dbInstance()

	EmployeesCollection = client.Database("main").Collection("employees")

	SamilUsersCollection = client.Database("users").Collection("samil_users")

	FleetUsersCollection = client.Database("users").Collection("fleet_users")

	VehiclesCollection = client.Database("vehicles").Collection("fleet_vehicles")

	SP_LoadingCollection = client.Database("pricing_panel").Collection("loading")

	SP_BusCollection = client.Database("pricing_panel").Collection("bus")

	SP_MotorCab = client.Database("pricing_panel").Collection("motor_cab")

	SP_ConstructionEqp = client.Database("pricing_panel").Collection("cons_eqp")

	SP_FarmEqp = client.Database("pricing_panel").Collection("farm_eqp")

	SP_PartnerCollection = client.Database("pricing_panel").Collection("partner")

}
