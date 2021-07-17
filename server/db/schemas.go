package db

import "go.mongodb.org/mongo-driver/bson/primitive"

type User struct {
	ID       primitive.ObjectID `json:"_id,omitempty" bson:"_id,omitempty"`
	Email    string             `json:"Email,omitempty" bson:"Email,omitempty"`
	Username string             `json:"Username,omitempty" bson:"Username,omitempty"`
	Type     string             `json:"Type,omitempty" bson:"Type,omitempty"`
	Password string             `json:"Password,omitempty" bson:"Password,omitempty"`
}
