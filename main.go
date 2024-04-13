package main

import (
	"example.com/gocrud/Initializers"
	"example.com/gocrud/controllers"
	"github.com/gin-gonic/gin"
)

func init() {
	initializers.LoadEnvVariables()
	initializers.ConnectToDB()
}

func main() {

	r := gin.Default()

	r.POST("/posts", controllers.PostCreate)
	r.GET("/posts", controllers.PostsIndexs)
	r.DELETE("posts/:id", controllers.PostDelete)
	r.PUT("/posts/:id", controllers.PostsUpdate)
	r.GET("/posts/:id", controllers.PostShow)

	r.Run() // listen and s
}
