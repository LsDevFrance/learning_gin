package main

import (
	"learning_gin/project/controllers"

	"github.com/gin-gonic/gin"
)

func main() {
	controllers.InitDatabase()
	r := gin.Default()
	r.GET("/books", controllers.FindBooks)
	r.POST("/books", controllers.CreateBooks)
	r.DELETE("/books/:id", controllers.DeleteBook)
	r.Run()
}
