package main

import (
	"./controllers"
	"github.com/gin-gonic/gin"
)

var router *gin.Engine

func main() {

	router = gin.Default()
	router.Use(CORSMiddleware())
	v1 := router.Group("/api/v1")
	{
		/*** VAPES ***/
		vapes := new(controllers.VapesController)
		v1.GET("/vapes", vapes.FindAll)
		v1.GET("/vapes/:id", vapes.FindById)

		/*** GOODS ***/
		goods := new(controllers.GoodsController)
		v1.GET("/goods", goods.FindAll)
	}
	router.Run("127.0.0.1:8080")
}

func CORSMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		c.Header("Allow", "POST, GET, OPTIONS, PUT, DELETE, UPDATE")
		c.Header("Access-Control-Allow-Methods", "POST, GET, OPTIONS, PUT, DELETE, UPDATE")
		c.Header("Access-Control-Allow-Origin", "*")
		c.Header("Access-Control-Max-Age", "86400")
		c.Header("Access-Control-Allow-Headers", "X-Requested-With, Content-Type, Origin, Authorization, Accept, Client-Security-Token, Accept-Encoding, x-access-token")
		c.Header("Access-Control-Allow-Credentials", "true")
	}
}
