package main

import (
	"github.com/gin-gonic/gin"
	"gitlab.com/hatemosphere/API-Exercise/solution/api/db"
	"gitlab.com/hatemosphere/API-Exercise/solution/api/handlers/passengers"
	"gitlab.com/hatemosphere/API-Exercise/solution/api/middlewares"
)

const (
	Port      = "3000"
	ApiPrefix = "/api"
)

func init() {
	db.Connect()
}

func main() {
	router := gin.Default()
	router.Use(middlewares.Connect)

	group := router.Group(ApiPrefix + "/people")
	group.GET("/:id", passengers.GetOne)
	group.GET("/", passengers.List)
	group.POST("/", passengers.Create)
	group.DELETE("/:id", passengers.Delete)
	group.PUT("/:id", passengers.Update)

	router.Run(":" + Port)
}
