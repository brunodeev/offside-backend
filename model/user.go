package model

import "go.mongodb.org/mongo-driver/bson/primitive"

type User struct {
	ID       primitive.ObjectID `json:"id,omitempty" bson:"_id"`
	Name     string             `json:"name,omitempty" bson:"name"`
	Email    string             `json:"email,omitempty" bson:"email"`
	Password string             `json:"password,omitempty" bson:"password"`
}
