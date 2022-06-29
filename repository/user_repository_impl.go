package repository

import (
	"context"

	"github.com/vincen320/user-service-mongodb/exception"
	"github.com/vincen320/user-service-mongodb/model/domain"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

type UserRepositoryImpl struct {
}

func NewUserRepository() UserRepository {
	return &UserRepositoryImpl{}
}

func (us *UserRepositoryImpl) Save(ctx context.Context, db *mongo.Database, user domain.User) (domain.User, error) {
	result, err := db.Collection("user").InsertOne(ctx, user)
	if err != nil {
		return user, err //500 internal server error
	}
	id := result.InsertedID
	idObject, ok := id.(primitive.ObjectID)
	if !ok {
		//user.Id = [12]byte{88, 88, 88, 88, 88, 88, 88, 88, 88, 88, 88, 88} //ID: XXXXXXXXXXX
		return user, exception.ErrParseId // 500 internal server error
	}
	user.Id = idObject
	return user, nil //success insert user
}

func (us *UserRepositoryImpl) Update(ctx context.Context, db *mongo.Database, user domain.User) (domain.User, error) {
	filter := bson.M{
		"_id": user.Id,
	}
	result, err := db.Collection("user").UpdateOne(ctx, filter, bson.M{
		"$set": user,
	})

	if err != nil {
		return user, err // 500 internal server error
	}
	if result.MatchedCount == 0 {
		return user, exception.NewNotFoundError("user not found") //404 not found, return error trs smpe service ttp return error
	}
	return user, nil //berhasil update
}

func (us *UserRepositoryImpl) Delete(ctx context.Context, db *mongo.Database, userId primitive.ObjectID) error {
	filter := bson.M{
		"_id": userId,
	}

	result, err := db.Collection("user").DeleteOne(ctx, filter)
	if err != nil {
		return err
	}
	if result.DeletedCount == 0 {
		return exception.NewNotFoundError("user not found") //404 not found, return error trs smpe service ttp return error
	}
	return nil //berhasil delete
}

func (us *UserRepositoryImpl) FindById(ctx context.Context, db *mongo.Database, userId primitive.ObjectID) (domain.User, error) {
	filter := bson.M{
		"_id": userId,
	}

	result := db.Collection("user").FindOne(ctx, filter)
	var user domain.User

	err := result.Decode(&user)

	if err != nil {
		if err == mongo.ErrNoDocuments {
			return user, exception.NewNotFoundError("user not found") //404 not found, return error trs smpe service ttp return error
		}
		return user, err
	}

	return user, nil
}

func (us *UserRepositoryImpl) FindAll(ctx context.Context, db *mongo.Database) ([]domain.User, error) {
	cursor, err := db.Collection("user").Find(ctx, bson.M{})
	var users []domain.User
	if err != nil {
		if err == mongo.ErrNoDocuments {
			return users, exception.NewNotFoundError("user not found") //404 not found, return error trs smpe service ttp return error
		}
		return users, err
	}
	defer cursor.Close(ctx)

	err = cursor.All(ctx, &users)
	if err != nil {
		return users, exception.ErrDecode //panic
	}
	return users, nil
}

// filter := bson.M{
// 	"username": userName,
// } terlihat bisa ganti key nya dengan parameter/sebuah variable, berarti bisa buat Fungsi FindBy(bebas), dengan memasukkan ...parameter(Variadic)

func (uc *UserRepositoryImpl) FindByUsername(ctx context.Context, db *mongo.Database, userName string) (domain.User, error) {
	filter := bson.M{
		"username": userName,
	}

	result := db.Collection("user").FindOne(ctx, filter)
	var user domain.User

	err := result.Decode(&user)

	if err != nil {
		if err == mongo.ErrNoDocuments {
			return user, exception.NewNotFoundError("user not found") //404 not found, return error trs smpe service ttp return error
		}
		return user, err
	}

	return user, nil
}
