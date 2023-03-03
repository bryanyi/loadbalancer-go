package main

import (
	"bryanyi.com/controllers"
	"bryanyi.com/initializers"
	"github.com/gin-gonic/gin"
)

func init() {
	initializers.LoadEnvVariables()
	initializers.ConnectDatabase()
}

func main() {
	r := gin.Default()
	r.GET("/", controllers.PostsCreate)
	r.Run()
}
