package main

import (
	"mymodule/config"
	"mymodule/modules/auth"

	"github.com/gin-gonic/gin"
	cors "github.com/rs/cors/wrapper/gin"
)

func main() {
	db := config.Connect()

	router := gin.Default()
	router.Use(cors.AllowAll())

	firstPath := router.Group("api")
	auth.NewAuthHandler(firstPath, auth.AuthRegistry(db))
	

	router.Run()
}
