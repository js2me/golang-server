package db

import (
	"../models"
	"gopkg.in/mgo.v2"
)

const (
	dbAddr     = "localhost:27017"
	dbName     = "shop"
	goodsCName = "goods"
	vapesCName = "vapes"
)

var session, err = mgo.Dial(dbAddr)

func FindGoods(query interface{}) ([]models.Good, error) {
	var results []models.Good
	if err := session.DB(dbName).C(goodsCName).Find(query).All(&results); err != nil {
		panic(err)
	}
	return results, err
}

func FindVapes(query interface{}) ([]models.Vape, error) {
	var results []models.Vape
	if err := session.DB(dbName).C(vapesCName).Find(query).All(&results); err != nil {
		panic(err)
	}
	return results, err
}

func FindVapeById(id interface{}) (models.Vape, error) {
	var result models.Vape
	if err := session.DB(dbName).C(vapesCName).FindId(id).One(&result); err != nil {
		panic(err)
	}
	return result, err
}
