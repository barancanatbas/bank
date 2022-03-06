package models

import "go.mongodb.org/mongo-driver/bson/primitive"

type Account struct {
	ID      primitive.ObjectID `json:"_id,omitempty" bson:"_id,omitempty"`
	Name    string             `json:"Name,omitempty"`
	Status  bool               `json:"status,omitempty"`
	Balance int                `json:"balance,omitempty"`
}
