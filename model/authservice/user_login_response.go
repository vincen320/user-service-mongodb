package authservice

import "go.mongodb.org/mongo-driver/bson/primitive"

type UserServiceLoginResponse struct {
	Id           primitive.ObjectID `bson:"_id,omitempty" json:"id,omitempty"`
	Username     string             `bson:"username,omitempty" json:"username,omitempty"`
	Password     string             `bson:"password,omitempty" json:"password,omitempty"`
	CreatedAt    int64              `bson:"created_At,omitempty" json:"createdAt,omitempty"`
	LastModified int64              `bson:"last_modified,omitempty" json:"lastOnline,omitempty"` //json = last_online karena auth_service terimanya last_online
}
