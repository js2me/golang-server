package controllers

import (
	"github.com/gin-gonic/gin"
	"db"
	"utils"
	"gopkg.in/mgo.v2/bson"
	"strconv"
)

//UserController ...
type VapesController struct{}

func (ctrl *VapesController) getFilter(c *gin.Context) map[string]interface{} {
	query := bson.M{}
	if c, ok := c.GetQuery("company"); ok {
		query["company"] = c
	}
	if k, ok := c.GetQuery("kit"); ok {
		query["is_kit"], _ = strconv.ParseBool(k)
	}

	return query
}

func (ctrl *VapesController) FindAll(c *gin.Context) {
	var result, err = db.FindVapes(ctrl.getFilter(c))
	if utils.IsNoError(err, c) {
		c.JSON(200, result)
	}
}

func (ctrl *VapesController) FindById(c *gin.Context) {
	var result, err = db.FindVapeById(bson.ObjectIdHex(c.Param("id")))
	if utils.IsNoError(err, c) {
		c.JSON(200, result)
	}
}
