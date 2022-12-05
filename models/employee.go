package models

import (
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type AllowedOps struct {
	View   bool `json:"view"`
	Add    bool `json:"add"`
	Edit   bool `json:"edit"`
	Delete bool `json:"delete"`
}

type Privillages struct {
	Dashboard          AllowedOps `json:"dashboard"`
	CustomerMaster     AllowedOps `json:"customer_master" bson:"customer_master"`
	EmployeeMaster     AllowedOps `json:"employee_master" bson:"employee_master"`
	UserMaster         AllowedOps `json:"user_master" bson:"user_master"`
	VehicleMaster      AllowedOps `json:"vehicle_master" bson:"vehicle_master"`
	AdminMaster        AllowedOps `json:"admin_master" bson:"admin_master"`
	BulkDataMaster     AllowedOps `json:"bulk_data_master" bson:"bulk_data_master"`
	OtherOptionsMaster AllowedOps `json:"other_options_master" bson:"other_options_master"`
}

type Employee struct {
	ID            primitive.ObjectID `json:"_id" bson:"_id"`
	Name          string             `json:"name"`
	Email         string             `json:"email"`
	Contact       string             `json:"contact"`
	Password      string             `json:"password"`
	ProfilePhoto  string             `json:"profile_photo" bson:"profile_photo"`
	Gender        string             `json:"gender"`
	DOB           string             `json:"date_of_birth" bson:"date_of_birth"`
	DOJ           string             `json:"date_of_joining" bson:"date_of_joining"`
	Designation   string             `json:"designation"`
	Role          string             `json:"role"`
	ReportTo      string             `json:"report_to" bson:"report_to"`
	AltContact    string             `json:"alt_contact" bson:"alt_contact"`
	Address       string             `json:"address"`
	Location      string             `json:"location"`
	Active        bool               `json:"active"`
	EmployeeType  string             `json:"employee_type" bson:"employee_type"`
	AssignedUsers []string           `json:"assigned_users" bson:"assigned_users"`
	Privillages   Privillages        `json:"privillages"`
	CreatedOn     primitive.DateTime `json:"created_on" bson:"created_on"`
	UpdatedOn     primitive.DateTime `json:"updated_on" bson:"updated_on"`
	CreatedBy     string             `json:"created_by" bson:"created_by"`
	UpdatedBy     string             `json:"updated_by" bson:"updated_by"`
}
