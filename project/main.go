package main

import (
	"learning_gin/project/controllers"

	"github.com/gin-gonic/gin"
)

func main() {
	controllers.InitDatabase()
	r := gin.Default()
	r.Run()
}
