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
	router.GET(ApiPrefix+"/people/:_id", passengers.GetOne)
	router.GET(ApiPrefix+"/people", passengers.List)
	router.POST(ApiPrefix+"/people", passengers.Create)
	router.DELETE(ApiPrefix+"/people/:_id", passengers.Delete)
	router.PUT(ApiPrefix+"/people/:_id", passengers.Update)

	router.Run(":" + Port)
}
