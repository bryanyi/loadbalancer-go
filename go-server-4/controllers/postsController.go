package controllers

import "github.com/gin-gonic/gin"

func PostsCreate (c *gin.Context) {
	c.JSON(200, gin.H{
		"server4": "received",
	})
}
