package main

import (
	"./controllers"
	"github.com/gin-gonic/gin"
	"os"
)

var router *gin.Engine

func main() {
	port := os.Getenv("PORT")
	if port == "" {
		port = "8000"
	}

	router = gin.Default()
	router.Use(CORSMiddleware())

	// the jwt middleware


	router.POST("/login", controllers.AuthMiddleware.LoginHandler)
	//router.POST("/registration", controllers.AuthMiddleware.Re)

	v1 := router.Group("/api/v1")
	{
		/*** VAPES ***/
		vapes := new(controllers.VapesController)
		v1.GET("/vapes", vapes.FindAll)
		v1.GET("/vapes/:id", vapes.FindById)

		/*** GOODS ***/
		goods := new(controllers.GoodsController)
		v1.GET("/goods", goods.FindAll)

		/*** AUTH ***/
		auth := new(controllers.AuthController)
		v1.POST("/auth/registration", auth.Registration)
		authR := v1.Group("/auth")
		{
			authR.Use(controllers.AuthMiddleware.MiddlewareFunc())
			//auth.GET("/account", helloHandler)
			authR.GET("/refresh_token", controllers.AuthMiddleware.RefreshHandler)
		}

	}
	router.Run(":" + port)
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
