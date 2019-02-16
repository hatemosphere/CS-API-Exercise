package middlewares

import (
	"../db"
	"github.com/gin-gonic/gin"
)

func Connect(c *gin.Context) {
	s := db.Session.Clone()

	defer s.Close()

	c.Set("db", s.DB(db.Mongo.Database))
	c.Next()
}
