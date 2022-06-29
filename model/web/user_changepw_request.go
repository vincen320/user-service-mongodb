package web

import "go.mongodb.org/mongo-driver/bson/primitive"

type UserChangePwRequest struct {
	Id          primitive.ObjectID `bson:"_id,omitempty" json:"id,omitempty"`
	OldPassword string             `validate:"required,min=6,max=20" bson:"old_password,omitempty" json:"old_password,omitempty"`
	NewPassword string             `validate:"required,min=6,max=20" bson:"new_password,omitempty" json:"new_password,omitempty"`
}
