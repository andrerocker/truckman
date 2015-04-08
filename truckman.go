package main

import (
	"./app/controllers"
	"github.com/gin-gonic/gin"
)

func main() {
	service := gin.Default()
	v1 := service.Group("/v1")
	v1.POST("/map/:name", controller.CreateNewMap)
	v1.GET("/map/:name/trace", controller.TraceRoute)

	service.Run(":8080")
}
