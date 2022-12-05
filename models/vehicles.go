package models

import (
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type Document struct {
	IssueDate  string `json:"issue_date" bson:"issue_date"`
	ExpiryDate string `json:"expiry_date" bson:"expiry_date"`
	DocUrl     string `json:"doc_url"`
}

type VehicleDocuments struct {
	RCUrl            string   `json:"rc_url" bson:"rc_url"`
	Fitness          Document `json:"fitness"`
	Permit           Document `json:"permit"`
	Insurance        Document `json:"insurance"`
	SpeedLimitDevice Document `json:"speed_limit_device" bson:"speed_limit_device"`
	ReflectiveTape   Document `json:"reflective_tape" bson:"reflective_tape"`
	RoadTax          Document `json:"road_tax" bson:"road_tax"`
	Pollution        Document `json:"pollution"`
}

type Vehicle struct {
	ID                  primitive.ObjectID `json:"_id" bson:"_id"`
	User                string             `json:"user"`
	Company             string             `json:"company"`
	RegisterationNumber string             `json:"registeration_number" bson:"registeration_number"`
	Model               string             `json:"model"`
	ManufactureDate     string             `json:"manufacture_date"`
	RegisterationDate   string             `json:"registeration_date" bson:"registeration_date"`
	ManufacturedBy      string             `json:"manufactured_by"`
	VehicleDocuments    VehicleDocuments   `json:"vehicle_documents" bson:"vehicle_documents"`
	CreatedOn           primitive.DateTime `json:"created_on" bson:"created_on"`
	UpdatedOn           primitive.DateTime `json:"updated_on" bson:"updated_on"`
	CreatedBy           string             `json:"created_by" bson:"created_by"`
	UpdatedBy           string             `json:"updated_by" bson:"updated_by"`
}
