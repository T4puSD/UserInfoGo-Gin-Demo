package model

import "go.mongodb.org/mongo-driver/bson/primitive"

type User struct {
	Id    primitive.ObjectID `json:"id,omitempty"`
	Name  string             `json:"name" validate:"required"`
	Email string             `json:"email" validate:"required"`
}