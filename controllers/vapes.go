package controllers

import (
	"../db"
	"github.com/gin-gonic/gin"
	"gopkg.in/mgo.v2/bson"
)

//UserController ...
type VapesController struct{}

func (ctrl *VapesController) FindAll(c *gin.Context) {
	var result, err = db.FindAll("shop", "vapes", nil)
	if IsNoError(err, c) {
		c.JSON(200, result)
	}
}

func (ctrl *VapesController) FindById(c *gin.Context) {
	var result, err = db.FindOne("shop", "vapes",
		bson.M{
			"_id": bson.ObjectIdHex(c.Param("id")),
		})
	if IsNoError(err, c) {
		c.JSON(200, result)
	}
}
