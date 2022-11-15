package controller

import (
	"github.com/gin-gonic/gin"
)

func FindUserById(c *gin.Context) {
	c.JSON(200, gin.H{
		"message": "pong",
	})
}

func FindUserByEmail(c *gin.Context) {
}
