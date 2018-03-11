package models

import (
	"gopkg.in/mgo.v2/bson"
)

type Vape struct {
	ID    bson.ObjectId `json:"_id" bson:"_id"`
	Name  string        `json:"name" bson:"name"`
	Price float32       `json:"price" bson:"price"`
}
