package service

import (
	"context"

	"github.com/vincen320/user-service-mongodb/model/authservice"
	"github.com/vincen320/user-service-mongodb/model/web"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type UserService interface {
	Create(ctx context.Context, createRequest web.UserCreateRequest) (web.UserResponse, error)
	Update(ctx context.Context, updateRequest web.UserUpdateRequest) (web.UserResponse, error)
	Delete(ctx context.Context, userId primitive.ObjectID) error
	FindById(ctx context.Context, userId primitive.ObjectID) (web.UserResponse, error)
	FindAll(ctx context.Context) ([]web.UserResponse, error)
	ChangePassword(ctx context.Context, changePwRequest web.UserChangePwRequest) error
	FindByUsername(ctx context.Context, username string) (authservice.UserServiceLoginResponse, error)
}
