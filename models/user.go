package models

import "go.mongodb.org/mongo-driver/bson/primitive"

type User struct {
	ID       primitive.ObjectID `json:"_id" bson:"_id"`
	UserID   string             `json:"userid"`
	Location string             `json:"location"`
	UserLink string             `json:"userlink"`
	Password string             `json:"password"`
	Active   bool               `json:"active"` // active || non-active
}
