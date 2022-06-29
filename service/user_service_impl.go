package service

import (
	"context"
	"time"

	"github.com/go-playground/validator/v10"
	"github.com/vincen320/user-service-mongodb/exception"
	"github.com/vincen320/user-service-mongodb/helper"
	"github.com/vincen320/user-service-mongodb/model/authservice"
	"github.com/vincen320/user-service-mongodb/model/domain"
	"github.com/vincen320/user-service-mongodb/model/web"
	"github.com/vincen320/user-service-mongodb/repository"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

type UserServiceImpl struct {
	DB         *mongo.Database
	Validator  *validator.Validate
	Repository repository.UserRepository
}

func NewUserService(db *mongo.Database, validator *validator.Validate, repository repository.UserRepository) UserService {
	return &UserServiceImpl{
		DB:         db,
		Validator:  validator,
		Repository: repository,
	}
}

func (us *UserServiceImpl) Create(ctx context.Context, createRequest web.UserCreateRequest) (web.UserResponse, error) {
	var result web.UserResponse
	var savedUser domain.User
	err := us.Validator.Struct(createRequest)
	if err != nil {
		return result, err //Bad Request
	}

	session, err := us.DB.Client().StartSession()
	if err != nil {
		return result, err // Internal Server Error
	}
	defer session.EndSession(context.Background())

	err = mongo.WithSession(ctx, session, func(sc mongo.SessionContext) error {
		err := session.StartTransaction()
		if err != nil {
			return err // Abort
		}

		hashedPassword, err := helper.BcryptPassword(createRequest.Password)
		if err != nil {
			return err // Abort
		}

		timeNow := time.Now().UTC().UnixMilli()
		savedUser, err = us.Repository.Save(ctx, us.DB, domain.User{
			Username:     createRequest.Username,
			Password:     hashedPassword,
			CreatedAt:    timeNow,
			LastModified: timeNow,
		})
		if err != nil {
			return err //Abort
		}

		err = session.CommitTransaction(ctx)
		if err != nil {
			return err //Abort
		}
		return nil
	})

	if err != nil {
		errAbort := session.AbortTransaction(ctx)
		if errAbort != nil {
			return result, errAbort // Internal Server Error
		}
		return result, err // Err yang terdapat direturn di fungsi diatas (mongo.WithSession)
	}

	result.Id = savedUser.Id
	result.Username = savedUser.Username
	result.CreatedAt = savedUser.CreatedAt
	result.LastModified = savedUser.LastModified

	return result, nil
}

func (us *UserServiceImpl) Update(ctx context.Context, updateRequest web.UserUpdateRequest) (web.UserResponse, error) {
	var result web.UserResponse
	var updatedUser domain.User
	err := us.Validator.Struct(updateRequest)
	if err != nil {
		return result, err //Bad Request
	}

	session, err := us.DB.Client().StartSession()
	if err != nil {
		return result, err
	}
	defer session.EndSession(context.Background())

	err = mongo.WithSession(ctx, session, func(sc mongo.SessionContext) error {
		err := session.StartTransaction()
		if err != nil {
			return err // Abort
		}

		timeNow := time.Now().UTC().UnixMilli()
		updatedUser, err = us.Repository.Update(ctx, us.DB, domain.User{
			Id:           updateRequest.Id,
			Username:     updateRequest.Username,
			CreatedAt:    timeNow,
			LastModified: timeNow,
		})
		if err != nil {
			return err //Abort
		}

		err = session.CommitTransaction(ctx)
		if err != nil {
			return err //Abort
		}
		return nil
	})

	if err != nil {
		errAbort := session.AbortTransaction(ctx)
		if errAbort != nil {
			return result, errAbort // Internal Server Error
		}
		return result, err // Err yang terdapat direturn di fungsi diatas (mongo.WithSession)
	}

	result.Id = updatedUser.Id
	result.Username = updatedUser.Username
	result.CreatedAt = updatedUser.CreatedAt
	result.LastModified = updatedUser.LastModified

	return result, nil
}

func (us *UserServiceImpl) Delete(ctx context.Context, userId primitive.ObjectID) error {
	session, err := us.DB.Client().StartSession()
	if err != nil {
		return err
	}
	defer session.EndSession(context.Background())

	err = mongo.WithSession(ctx, session, func(sc mongo.SessionContext) error {
		err := session.StartTransaction()
		if err != nil {
			return err // Abort
		}

		err = us.Repository.Delete(ctx, us.DB, userId)
		if err != nil {
			return err //Abort
		}

		err = session.CommitTransaction(ctx)
		if err != nil {
			return err //Abort
		}
		return nil
	})

	if err != nil {
		errAbort := session.AbortTransaction(ctx)
		if errAbort != nil {
			return errAbort // Internal Server Error
		}
		return err // Err yang terdapat direturn di fungsi diatas (mongo.WithSession)
	}
	return nil
}

func (us *UserServiceImpl) FindById(ctx context.Context, userId primitive.ObjectID) (web.UserResponse, error) {
	var result web.UserResponse
	var user domain.User
	session, err := us.DB.Client().StartSession()
	if err != nil {
		return result, err
	}
	defer session.EndSession(context.Background())

	err = mongo.WithSession(ctx, session, func(sc mongo.SessionContext) error {
		err := session.StartTransaction()
		if err != nil {
			return err // Abort
		}

		user, err = us.Repository.FindById(ctx, us.DB, userId)
		if err != nil {
			return err //Abort
		}

		err = session.CommitTransaction(ctx)
		if err != nil {
			return err //Abort
		}
		return nil
	})

	if err != nil {
		errAbort := session.AbortTransaction(ctx)
		if errAbort != nil {
			return result, errAbort // Internal Server Error
		}
		return result, err // Err yang terdapat direturn di fungsi diatas (mongo.WithSession)
	}

	result.Id = user.Id
	result.Username = user.Username
	result.CreatedAt = user.CreatedAt
	result.LastModified = user.LastModified
	return result, nil
}

func (us *UserServiceImpl) FindAll(ctx context.Context) ([]web.UserResponse, error) {
	var results []web.UserResponse
	var users []domain.User

	session, err := us.DB.Client().StartSession()
	if err != nil {
		return results, err
	}
	defer session.EndSession(context.Background())

	err = mongo.WithSession(ctx, session, func(sc mongo.SessionContext) error {
		err := session.StartTransaction()
		if err != nil {
			return err // Abort
		}

		users, err = us.Repository.FindAll(ctx, us.DB)
		if err != nil {
			return err //Abort
		}

		err = session.CommitTransaction(ctx)
		if err != nil {
			return err //Abort
		}
		return nil
	})

	if err != nil {
		errAbort := session.AbortTransaction(ctx)
		if errAbort != nil {
			return results, errAbort // Internal Server Error
		}
		return results, err // Err yang terdapat direturn di fungsi diatas (mongo.WithSession)
	}

	for _, v := range users {
		results = append(results, web.UserResponse{
			Id:           v.Id,
			Username:     v.Username,
			CreatedAt:    v.CreatedAt,
			LastModified: v.LastModified,
		})
	}
	return results, nil
}

func (us *UserServiceImpl) ChangePassword(ctx context.Context, changePwRequest web.UserChangePwRequest) error {
	err := us.Validator.Struct(changePwRequest)
	if err != nil {
		return err //Bad Request
	}

	session, err := us.DB.Client().StartSession()
	if err != nil {
		return err
	}
	defer session.EndSession(context.Background())

	err = mongo.WithSession(ctx, session, func(sc mongo.SessionContext) error {
		err := session.StartTransaction()
		if err != nil {
			return err // Abort
		}
		user, err := us.Repository.FindById(ctx, us.DB, changePwRequest.Id)
		if err != nil {
			return err // Abort
		}

		isSame := helper.ComparePassword(user.Password, changePwRequest.OldPassword)
		if !isSame {
			return exception.NewBadRequestError("wrong old password")
		}

		hashNewPassword, err := helper.BcryptPassword(changePwRequest.NewPassword)
		if err != nil {
			return err // Abort
		}

		timeNow := time.Now().UTC().UnixMilli()
		_, err = us.Repository.Update(ctx, us.DB, domain.User{
			Id:           changePwRequest.Id,
			Password:     hashNewPassword,
			LastModified: timeNow,
		})

		if err != nil {
			return err //Abort
		}

		err = session.CommitTransaction(ctx)
		if err != nil {
			return err //Abort
		}
		return nil
	})

	if err != nil {
		errAbort := session.AbortTransaction(ctx)
		if errAbort != nil {
			return errAbort // Internal Server Error
		}
		return err // Err yang terdapat direturn di fungsi diatas (mongo.WithSession)
	}
	return nil
}

func (us *UserServiceImpl) FindByUsername(ctx context.Context, username string) (authservice.UserServiceLoginResponse, error) {
	var result authservice.UserServiceLoginResponse
	var user domain.User
	session, err := us.DB.Client().StartSession()
	if err != nil {
		return result, err
	}
	defer session.EndSession(context.Background())

	err = mongo.WithSession(ctx, session, func(sc mongo.SessionContext) error {
		err := session.StartTransaction()
		if err != nil {
			return err // Abort
		}

		user, err = us.Repository.FindByUsername(ctx, us.DB, username)
		if err != nil {
			return err //Abort
		}

		err = session.CommitTransaction(ctx)
		if err != nil {
			return err //Abort
		}
		return nil
	})

	if err != nil {
		errAbort := session.AbortTransaction(ctx)
		if errAbort != nil {
			return result, errAbort // Internal Server Error
		}
		return result, err // Err yang terdapat direturn di fungsi diatas (mongo.WithSession)
	}

	result.Id = user.Id
	result.Username = user.Username
	result.Password = user.Password
	result.CreatedAt = user.CreatedAt
	result.LastModified = user.LastModified
	return result, nil
}
