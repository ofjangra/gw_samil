package models

import (
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type FleetUser struct {
	ID        primitive.ObjectID `json:"_id" bson:"_id"`
	Firstname string             `json:"firstname"`
	Lastname  string             `json:"lastname"`
	Image_url string             `json:"image_url" bson:"image_url"`
	Email     string             `json:"email"`
	Phone     string             `json:"phone"`
	Password  string             `json:"password"`
	DOB       string             `json:"DOB"`
	Gender    string             `json:"gender"`
	Activated bool               `json:"activated"`
	Location  string             `json:"location"`
	Address   string             `json:"address"`
	Vehicles  []string           `json:"vehicles"`
	CreatedOn primitive.DateTime `json:"created_on" bson:"created_on"`
	UpdatedOn primitive.DateTime `json:"updated_on" bson:"updated_on"`
	CreatedBy string             `json:"created_by" bson:"created_by"`
	UpdatedBy string             `json:"updated_by" bson:"updated_by"`
}

type SamilUser struct {
	ID        primitive.ObjectID `json:"_id" bson:"_id"`
	Email     string             `json:"email"`
	Location  string             `json:"location"`
	UserLink  string             `json:"userlink"`
	Password  string             `json:"password"`
	Active    bool               `json:"active"`
	CreatedOn primitive.DateTime `json:"created_on" bson:"created_on"`
	UpdatedOn primitive.DateTime `json:"updated_on" bson:"updated_on"`
	CreatedBy string             `json:"created_by" bson:"created_by"`
	UpdatedBy string             `json:"updated_by" bson:"updated_by"`
}
