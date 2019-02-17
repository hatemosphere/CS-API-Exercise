package passengers

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/globalsign/mgo"
	"github.com/globalsign/mgo/bson"
	"gitlab.com/hatemosphere/API-Exercise/solution/api/models"
)

func List(c *gin.Context) {
	db := c.MustGet("db").(*mgo.Database)
	passengers := []models.Passenger{}

	err := db.C(models.CollectionPassengers).Find(nil).All(&passengers)
	if err != nil {
		c.Error(err)
	}

	c.JSON(http.StatusOK, passengers)
}

func Create(c *gin.Context) {
	db := c.MustGet("db").(*mgo.Database)

	passenger := models.Passenger{}
	err := c.BindJSON(&passenger)
	if err != nil {
		c.Error(err)
		return
	}

	err = db.C(models.CollectionPassengers).Insert(passenger)
	if err != nil {
		c.Error(err)
	}

	c.JSON(http.StatusCreated, passenger)
}

func GetOne(c *gin.Context) {
	db := c.MustGet("db").(*mgo.Database)

	passenger := models.Passenger{}
	oID := bson.ObjectIdHex(c.Param("_id"))
	err := db.C(models.CollectionPassengers).FindId(oID).One(&passenger)
	if err != nil {
		c.Error(err)
	}

	c.JSON(http.StatusOK, passenger)
}

func Delete(c *gin.Context) {
	db := c.MustGet("db").(*mgo.Database)

	oID := bson.ObjectIdHex(c.Param("_id"))
	err := db.C(models.CollectionPassengers).RemoveId(oID)
	if err != nil {
		c.Error(err)
	}

	c.Data(204, "application/json", make([]byte, 0))
}

func Update(c *gin.Context) {
	db := c.MustGet("db").(*mgo.Database)

	passenger := models.Passenger{}
	err := c.Bind(&passenger)
	if err != nil {
		c.Error(err)
		return
	}

	query := bson.M{"_id": bson.ObjectIdHex(c.Param("_id"))}
	doc := bson.M{
		"survived":                passenger.Survived,
		"passengerClass":          passenger.PassengerClass,
		"name":                    passenger.Name,
		"sex":                     passenger.Sex,
		"age":                     passenger.Age,
		"siblingsOrSpousesAboard": passenger.SiblingsOrSpousesAboard,
		"parentsOrChildrenAboard": passenger.ParentsOrChildrenAboard,
		"fare":                    passenger.Fare,
	}
	err = db.C(models.CollectionPassengers).Update(query, doc)
	if err != nil {
		c.Error(err)
	}

	c.Data(http.StatusOK, "application/json", make([]byte, 0))
}
