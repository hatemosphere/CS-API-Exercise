package middlewares

import (
	"github.com/gin-gonic/gin"
	"gitlab.com/hatemosphere/API-Exercise/solution/api/db"
)

func Connect(c *gin.Context) {
	s := db.Session.Clone()

	defer s.Close()

	c.Set("db", s.DB(db.Mongo.Database))
	c.Next()
}
