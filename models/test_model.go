package models

import "go.mongodb.org/mongo-driver/bson/primitive"

type Test struct {
	Id    primitive.ObjectID `json:"id,omitempty" bson:"_id,omitempty"`
	Title string `json:"title" bson:"title"`
	Description string `json:"description" bson:"description"`
}