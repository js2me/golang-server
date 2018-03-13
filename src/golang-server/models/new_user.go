package models

import "gopkg.in/mgo.v2/bson"

type NewUser struct {
	FName        string        `json:"fname" bson:"fname"`
	SName        string        `json:"sname" bson:"sname"`
	MName        string        `json:"mname" bson:"mname"`
	Born		 bson.MongoTimestamp `json:"born" bson:"born"`
	Sex			 string `json:"sex" bson:"sex"`
}
