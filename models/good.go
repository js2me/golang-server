package models

import (
	"gopkg.in/mgo.v2/bson"
)

type Good struct {
	ID           bson.ObjectId `json:"_id" bson:"_id"`
	Name         string        `json:"name" bson:"name"`
	Type         string        `json:"type" bson:"type"`
	Price        float32       `json:"price" bson:"price"`
	FullObjectId bson.ObjectId `json:"full_object_id" bson:"full_object_id"`
}
