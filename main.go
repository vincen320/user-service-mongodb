package main

import (
	"log"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
	"github.com/vincen320/user-service-mongodb/app"
	"github.com/vincen320/user-service-mongodb/controller"
	"github.com/vincen320/user-service-mongodb/middleware"
	"github.com/vincen320/user-service-mongodb/repository"
	"github.com/vincen320/user-service-mongodb/service"
)

func main() {
	mongoDB, err := app.ConnectMongo()
	if err != nil {
		panic(err)
	}
	validator := validator.New()
	userRepository := repository.NewUserRepository()
	userService := service.NewUserService(mongoDB, validator, userRepository)
	userController := controller.NewUserController(userService)

	router := gin.New()
	router.Use(middleware.PanicHandling())

	router.POST("/users", userController.Create)
	router.PUT("/users/:userId", userController.Update)
	router.DELETE("/users/:userId", userController.Delete)
	router.GET("/users/:userId", userController.FindById)
	router.GET("/users", userController.FindAll)
	router.PATCH("/users/changepassword/:userId", userController.ChangePassword)
	router.GET("/users/username/:username", userController.FindByUsername)

	server := http.Server{
		Addr:           ":8080",
		Handler:        router,
		ReadTimeout:    10 * time.Second,
		WriteTimeout:   10 * time.Second,
		MaxHeaderBytes: 1 << 20,
	}

	log.Println("User Service Start in 8080 port")
	err = server.ListenAndServe()
	if err != nil {
		panic("Cannot Start Server " + err.Error()) //500 Internal Server Error
	}
}
