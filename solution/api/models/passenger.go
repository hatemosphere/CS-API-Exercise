package models

import (
	"github.com/globalsign/mgo/bson"
)

const (
	CollectionPassengers = "passengers"
)

type Passenger struct {
	ID                      bson.ObjectId `json:"_id,omitempty" bson:"_id,omitempty"`
	Survived                bool          `json:"survived" bson:"survived" binding:"exists"`
	PassengerClass          int32         `json:"passengerClass" bson:"passengerClass" binding:"required"`
	Name                    string        `json:"name" bson:"name" binding:"required"`
	Sex                     string        `json:"sex" bson:"sex,omitempty" binding:"required"`
	Age                     float64       `json:"age" bson:"age" binding:"required"`
	SiblingsOrSpousesAboard int32         `json:"siblingsOrSpousesAboard" bson:"siblingsOrSpousesAboard" binding:"exists"`
	ParentsOrChildrenAboard int32         `json:"parentsOrChildrenAboard" bson:"parentsOrChildrenAboard" binding:"exists"`
	Fare                    float64       `json:"fare" bson:"fare" binding:"required"`
}
