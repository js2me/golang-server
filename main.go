package main

import (
	"./controllers"
	"github.com/appleboy/gin-jwt"
	"github.com/gin-gonic/gin"
	"os"
	"time"
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
	authMiddleware := &jwt.GinJWTMiddleware{
		Realm:      "test zone",
		Key:        []byte("secret key"),
		Timeout:    time.Hour,
		MaxRefresh: time.Hour,
		Authenticator: func(userId string, password string, c *gin.Context) (string, bool) {
			return userId, (userId == "admin" && password == "admin") || (userId == "test" && password == "test")
		},
		Authorizator: func(userId string, c *gin.Context) bool {
			return userId == "admin"
		},
		Unauthorized: func(c *gin.Context, code int, message string) {
			c.JSON(code, gin.H{
				"code":    code,
				"message": message,
			})
		},
		// TokenLookup is a string in the form of "<source>:<name>" that is used
		// to extract token from the request.
		// Optional. Default value "header:Authorization".
		// Possible values:
		// - "header:<name>"
		// - "query:<name>"
		// - "cookie:<name>"
		TokenLookup: "header:Authorization",
		// TokenLookup: "query:token",
		// TokenLookup: "cookie:token",

		// TokenHeadName is a string in the header. Default value is "Bearer"
		TokenHeadName: "Bearer",

		// TimeFunc provides the current time. You can override it to use another time value. This is useful for testing or if your server uses a different time zone than your tokens.
		TimeFunc: time.Now,
	}

	router.POST("/login", authMiddleware.LoginHandler)
	//router.POST("/registration", authMiddleware.Re)

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
		auth := v1.Group("/auth")
		{
			auth.Use(authMiddleware.MiddlewareFunc())
			//auth.GET("/account", helloHandler)
			auth.GET("/refresh_token", authMiddleware.RefreshHandler)
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
