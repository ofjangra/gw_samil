package models

import "go.mongodb.org/mongo-driver/bson/primitive"

type FleetUser struct {
	ID        primitive.ObjectID `json:"_id" bson:"_id"`
	Firstname string             `json:"firstname"`
	Lastname  string             `json:"lastname"`
	Email     string             `json:"email"`
	Phone     string             `json:"phone"`
	DOB       string             `json:"DOB"`
	Gender    string             `json:"gender"`
	Activated bool               `json:"activated"`
	Location  string             `json:"location"`
	Address   string             `json:"address"`
}
