package models

import (
	"gopkg.in/mgo.v2/bson"
)

type Vape struct {
	ID          bson.ObjectId `json:"id" bson:"_id"`
	Name        string        `json:"name" bson:"name"`
	Company     string        `json:"company" bson:"company"`
	Description string        `json:"description" bson:"description"`
	IsKit       bool          `json:"isKit" bson:"is_kit"`
}
