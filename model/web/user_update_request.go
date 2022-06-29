package web

import "go.mongodb.org/mongo-driver/bson/primitive"

type UserUpdateRequest struct {
	Id       primitive.ObjectID `bson:"_id,omitempty" json:"id,omitempty"`
	Username string             `validate:"required,min=6,max=15" bson:"username,omitempty" json:"username,omitempty"`
}
