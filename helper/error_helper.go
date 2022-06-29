package helper

import (
	"errors"
	"log"
	"net/http"
	"reflect"

	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
	"github.com/vincen320/user-service-mongodb/exception"
	"github.com/vincen320/user-service-mongodb/model/web"
)

var (
	badRequest      *exception.BadRequestError
	notFound        *exception.NotFoundError
	validationError *validator.ValidationErrors
)

func ReturnError(c *gin.Context, e error) {
	log.Println(reflect.TypeOf(e))
	if errors.As(e, &badRequest) {
		ErrorResponse(c, http.StatusBadRequest, badRequest.Error())
	} else if errors.As(e, &notFound) {
		ErrorResponse(c, http.StatusNotFound, notFound.Error())
	} else if errors.As(e, &validationError) {
		ErrorResponse(c, http.StatusBadRequest, validationError.Error())
	} else {
		ErrorResponse(c, http.StatusInternalServerError, e.Error())
	}
}

func ErrorResponse(c *gin.Context, code int, message string) {
	c.JSON(code, web.WebResponse{
		Code:    code,
		Message: message,
	})
}
