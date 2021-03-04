package entity

import "go.mongodb.org/mongo-driver/bson/primitive"

type User struct {
	ObjectID primitive.ObjectID `json:"-" bson:"_id, omitempty"`
	Id       string             `json:"id" bson:"-"`
	Name     string             `json:"name" bson:"name" `
	IsActive bool               `json:"isActive" bson:"isActive"`
}
