package controllers

import (
	"../db"
	"github.com/gin-gonic/gin"
	"gopkg.in/mgo.v2/bson"
	"strconv"
)

type GoodsController struct{}

func getFilterFromContext(c *gin.Context) map[string]interface{} {
	query := bson.M{}
	if goodsType, ok := c.GetQuery("type"); ok {
		query["type"] = goodsType
	}
	priceExtrems := bson.M{}
	if minPrice, ok := c.GetQuery("min_price"); ok {
		priceExtrems["$gt"], _ = strconv.ParseFloat(minPrice, 32)
	}
	if maxPrice, ok := c.GetQuery("max_price"); ok {
		priceExtrems["$lt"], _ = strconv.ParseFloat(maxPrice, 32)
	}
	if len(priceExtrems) > 0 {
		query["price"] = priceExtrems
	}
	return query
}

func (ctrl *GoodsController) FindAll(c *gin.Context) {
	var result, err = db.FindAll("shop", "goods", getFilterFromContext(c))
	if IsNoError(err, c) {
		c.JSON(200, result)
	}

}
