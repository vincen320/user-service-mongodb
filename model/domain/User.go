package domain

import "go.mongodb.org/mongo-driver/bson/primitive"

type User struct {
	Id           primitive.ObjectID `bson:"_id,omitempty"`
	Username     string             `bson:"username,omitempty"`
	Password     string             `bson:"password,omitempty"`
	CreatedAt    int64              `bson:"created_at,omitempty"`
	LastModified int64              `bson:"last_modified,omitempty"`
}
