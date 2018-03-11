package controllers

import "github.com/gin-gonic/gin"

func IsNoError(err error, c *gin.Context) bool {
	noErr := err == nil
	if !noErr {
		panic(err)
		c.Error(err)
	}
	return noErr
}
