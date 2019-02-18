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
	var passengers []models.Passenger

	err := db.C(models.CollectionPassengers).Find(nil).All(&passengers)
	if err != nil {
		c.Error(err)
		if err == mgo.ErrNotFound {
			c.JSON(http.StatusNotFound, "Passengers not found")
			return
		}
	}

	c.JSON(http.StatusOK, passengers)
}

func Create(c *gin.Context) {
	db := c.MustGet("db").(*mgo.Database)

	var passenger models.Passenger
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

	var passenger models.Passenger
	if !bson.IsObjectIdHex(c.Param("_id")) {
		c.JSON(http.StatusUnprocessableEntity, "_id field is incorrect")
		return
	}
	oID := bson.ObjectIdHex(c.Param("_id"))
	err := db.C(models.CollectionPassengers).FindId(oID).One(&passenger)
	if err != nil {
		c.Error(err)
		if err == mgo.ErrNotFound {
			c.JSON(http.StatusNotFound, "Passenger not found")
			return
		}
	}

	c.JSON(http.StatusOK, passenger)
}

func Delete(c *gin.Context) {
	db := c.MustGet("db").(*mgo.Database)

	if !bson.IsObjectIdHex(c.Param("_id")) {
		c.JSON(http.StatusUnprocessableEntity, "_id field is incorrect")
		return
	}
	oID := bson.ObjectIdHex(c.Param("_id"))
	err := db.C(models.CollectionPassengers).RemoveId(oID)
	if err != nil {
		c.Error(err)
	}

	c.Data(http.StatusNoContent, "application/json", make([]byte, 0))
}

func Update(c *gin.Context) {
	db := c.MustGet("db").(*mgo.Database)

	var passenger models.Passenger
	err := c.Bind(&passenger)
	if err != nil {
		c.Error(err)
		return
	}

	if !bson.IsObjectIdHex(c.Param("_id")) {
		c.JSON(http.StatusUnprocessableEntity, "_id field is incorrect")
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
