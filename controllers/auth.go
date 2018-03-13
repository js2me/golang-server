package controllers

import (
	"github.com/gin-gonic/gin"
	"../models"
	"time"
	"github.com/appleboy/gin-jwt"
	"fmt"
)

var AuthMiddleware = &jwt.GinJWTMiddleware{
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


type AuthController struct {
}

func (ctrl *AuthController) Registration(c *gin.Context) {
	var regForm models.NewUser
	if err := c.BindJSON(&regForm);err != nil{
		panic(err)
		c.Error(err)
		return
	}
	fmt.Println(regForm)
}
