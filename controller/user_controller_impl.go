package controller

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/vincen320/user-service-mongodb/exception"
	"github.com/vincen320/user-service-mongodb/helper"
	"github.com/vincen320/user-service-mongodb/model/web"
	"github.com/vincen320/user-service-mongodb/service"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type UserControllerImpl struct {
	Service service.UserService
}

func NewUserController(service service.UserService) UserController {
	return &UserControllerImpl{
		Service: service,
	}
}

func (uc *UserControllerImpl) Create(c *gin.Context) {
	var createRequest web.UserCreateRequest
	err := c.ShouldBind(&createRequest)
	if err != nil {
		helper.ReturnError(c, err) //Validation error
		return
	}

	response, err := uc.Service.Create(c, createRequest)
	if err != nil {
		helper.ReturnError(c, err)
		return
	}

	c.JSON(http.StatusOK, web.WebResponse{
		Code:    http.StatusOK,
		Message: "Success Create New User",
		Data:    response,
	})
}

func (uc *UserControllerImpl) Update(c *gin.Context) {
	idString := c.Param("userId")
	var userId primitive.ObjectID

	userId, err := primitive.ObjectIDFromHex(idString)
	if err != nil {
		notFoundErr := exception.NewNotFoundError("user not found with error: " + err.Error())
		helper.ReturnError(c, notFoundErr)
		return
	}

	var updateRequest web.UserUpdateRequest
	err = c.ShouldBind(&updateRequest)
	if err != nil {
		helper.ReturnError(c, err) //Validation Error
		return
	}
	updateRequest.Id = userId
	response, err := uc.Service.Update(c, updateRequest)
	if err != nil {
		helper.ReturnError(c, err) //any error
		return
	}

	c.JSON(http.StatusOK, web.WebResponse{
		Code:    http.StatusOK,
		Message: "Success Update Username",
		Data:    response,
	})
}

func (uc *UserControllerImpl) Delete(c *gin.Context) {
	idString := c.Param("userId")
	var userId primitive.ObjectID

	userId, err := primitive.ObjectIDFromHex(idString)
	if err != nil {
		notFoundErr := exception.NewNotFoundError("user not found with error: " + err.Error())
		helper.ReturnError(c, notFoundErr)
		return
	}

	err = uc.Service.Delete(c, userId)
	if err != nil {
		helper.ReturnError(c, err) //any error
		return
	}

	c.JSON(http.StatusOK, web.WebResponse{
		Code:    http.StatusOK,
		Message: "Success Delete User",
	})
}

func (uc *UserControllerImpl) FindById(c *gin.Context) {
	idString := c.Param("userId")
	var userId primitive.ObjectID

	userId, err := primitive.ObjectIDFromHex(idString)
	if err != nil {
		notFoundErr := exception.NewNotFoundError("user not found with error: " + err.Error())
		helper.ReturnError(c, notFoundErr)
		return
	}

	response, err := uc.Service.FindById(c, userId)
	if err != nil {
		helper.ReturnError(c, err) //any error
		return
	}

	c.JSON(http.StatusOK, web.WebResponse{
		Code:    http.StatusOK,
		Message: "Success Find By Id",
		Data:    response,
	})
}

func (uc *UserControllerImpl) FindAll(c *gin.Context) {
	responses, err := uc.Service.FindAll(c)
	if err != nil {
		helper.ReturnError(c, err) //any error
		return
	}

	c.JSON(http.StatusOK, web.WebResponse{
		Code:    http.StatusOK,
		Message: "Success Find All",
		Data:    responses,
	})
}

func (uc *UserControllerImpl) ChangePassword(c *gin.Context) {
	idString := c.Param("userId")
	var userId primitive.ObjectID

	userId, err := primitive.ObjectIDFromHex(idString)
	if err != nil {
		notFoundErr := exception.NewNotFoundError("user not found with error: " + err.Error())
		helper.ReturnError(c, notFoundErr)
		return
	}

	var changePwRequest web.UserChangePwRequest
	err = c.ShouldBind(&changePwRequest)
	if err != nil {
		helper.ReturnError(c, err) //Validation Error
		return
	}
	changePwRequest.Id = userId
	err = uc.Service.ChangePassword(c, changePwRequest)
	if err != nil {
		helper.ReturnError(c, err) //any error
		return
	}

	c.JSON(http.StatusOK, web.WebResponse{
		Code:    http.StatusOK,
		Message: "Success Change Password",
	})
}

func (uc *UserControllerImpl) FindByUsername(c *gin.Context) {
	username := c.Param("username")

	response, err := uc.Service.FindByUsername(c, username)
	if err != nil {
		helper.ReturnError(c, err) //any error
		return
	}

	c.JSON(http.StatusOK, web.WebResponse{
		Code:    http.StatusOK,
		Message: "Success Find By Username",
		Data:    response,
	})
}
