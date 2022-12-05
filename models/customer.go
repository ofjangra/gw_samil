package models

import "go.mongodb.org/mongo-driver/bson/primitive"

type Customer struct {
	ID                   primitive.ObjectID `json:"_id" bson:"_id"`
	Name                 string             `json:"name"`
	Image_url            string             `json:"image_url" bson:"image_url"`
	Email                string             `json:"email"`
	Address              string             `json:"address"`
	OwnerName            string             `json:"owner_name" bson:"owner_name"`
	OwnerEmail           string             `json:"owner_email" bson:"owner_email"`
	OwnerContact         string             `json:"owner_contact" bson:"owner_contact"`
	ContactPersonName    string             `json:"contact_person_name" bson:"contact_person_name"`
	ContactPersonEmail   string             `json:"contact_person_email" bson:"contact_person_email"`
	ContactPersonContact string             `json:"contact_person_contact" bson:"contact_person_contact"`
	GSTIN                string             `json:"gstin"`
	CIN                  string             `json:"cin"`
	PanNumber            string             `json:"pan_number"`
	CreatedOn            primitive.DateTime `json:"created_on" bson:"updated_on"`
	UpdatedOn            primitive.DateTime `json:"updated_on" bson:"updated_on"`
	CreatedBy            string             `json:"created_by" bson:"created_by"`
	UpdatedBy            string             `json:"updated_by" bson:"updated_by"`
}
