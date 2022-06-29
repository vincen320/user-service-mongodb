package repository

import (
	"context"

	"github.com/vincen320/user-service-mongodb/model/domain"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

type UserRepository interface {
	Save(ctx context.Context, db *mongo.Database, user domain.User) (domain.User, error)
	Update(ctx context.Context, db *mongo.Database, user domain.User) (domain.User, error)
	Delete(ctx context.Context, db *mongo.Database, userId primitive.ObjectID) error
	FindById(ctx context.Context, db *mongo.Database, userId primitive.ObjectID) (domain.User, error)
	FindAll(ctx context.Context, db *mongo.Database) ([]domain.User, error)
	FindByUsername(ctx context.Context, db *mongo.Database, userName string) (domain.User, error)
}
