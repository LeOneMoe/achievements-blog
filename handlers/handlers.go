package handlers

import "github.com/gin-gonic/gin"

func PostPingPong(c *gin.Context) {
	c.JSON(200, gin.H{
		"message": "pong",
	})
}
