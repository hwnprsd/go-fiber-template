package models

import "go.mongodb.org/mongo-driver/bson/primitive"

type User struct {
	Id        primitive.ObjectID `json:"_id,omitempty"`
	FirstName string             `json:"firstName"`
	LastName  string             `json:"lastName"`
	Username  string             `json:"username"`
}
