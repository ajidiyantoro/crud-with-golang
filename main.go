package main

import (
	"log"

	"crud-with-golang/handler"
	"crud-with-golang/user"

	"github.com/gin-gonic/gin"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func main() {
	dsn := "root:Password12345@tcp(127.0.0.1:3306)/golang?charset=utf8mb4&parseTime=True&loc=Local"

	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatal(err.Error())
	}

	userRepository := user.NewRepository(db)
	userService := user.NewService(userRepository)
	userHandler := handler.NewHandler(userService)

	router := gin.Default()
	api := router.Group("/api/v1")
	api.GET("/users", userHandler.GetUsers)
	api.GET("/users/:id", userHandler.GetUserByID)
	api.POST("/users", userHandler.CreateUser)

	router.Run("localhost:8080")
}
