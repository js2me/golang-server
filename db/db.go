package db

import (
	"fmt"
	"gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"
)

var session, err = mgo.Dial("localhost:27017")

func FindAll(database string, collection string, query interface{}) ([]bson.M, error) {
	var results = []bson.M{}
	fmt.Println("Search from ", database, " in ", collection, " using query ", query)
	err := session.DB(database).C(collection).Find(query).All(&results)
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println("Results All: ", results)
	}
	return results, err
}

func FindOne(database string, collection string, query interface{}) (bson.M, error) {
	var result = bson.M{}
	fmt.Println("Search from ", database, " in ", collection, " using query ", query)
	err := session.DB(database).C(collection).Find(query).One(&result)
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println("Result One: ", result)
	}
	return result, err
}
