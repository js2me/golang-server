package controllers

import (
	"../db"
	"../utils"
	"github.com/gin-gonic/gin"
	"gopkg.in/mgo.v2/bson"
	"strconv"
)

type GoodsController struct{}

/**
returns goods filter from url's query parameters
*/
func (ctrl *GoodsController) getFilter(c *gin.Context) map[string]interface{} {
	query := bson.M{}
	if goodsType, ok := c.GetQuery("type"); ok {
		query["type"] = goodsType
	}
	priceExtrems := bson.M{}
	if minPrice, ok := c.GetQuery("pmin"); ok {
		priceExtrems["$gt"], _ = strconv.ParseFloat(minPrice, 32)
	}
	if maxPrice, ok := c.GetQuery("pmax"); ok {
		priceExtrems["$lt"], _ = strconv.ParseFloat(maxPrice, 32)
	}
	if len(priceExtrems) > 0 {
		query["price"] = priceExtrems
	}
	return query
}

func (ctrl *GoodsController) FindAll(c *gin.Context) {
	var result, err = db.FindGoods(ctrl.getFilter(c))
	if utils.IsNoError(err, c) {
		c.JSON(200, result)
	}

}
