package web

import "go.mongodb.org/mongo-driver/bson/primitive"

type UserResponse struct {
	Id           primitive.ObjectID `json:"id,omitempty"`
	Username     string             `json:"username,omitempty"`
	CreatedAt    int64              `json:"created_at,omitempty"`
	LastModified int64              `json:"last_modified,omitempty"`
}
