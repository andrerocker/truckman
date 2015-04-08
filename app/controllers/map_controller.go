package controller

import (
	"../models"
	"../services"
	"github.com/gin-gonic/gin"
	"github.com/gin-gonic/gin/binding"
)

func CreateNewMap(context *gin.Context) {
	context.Request.ParseForm()
	content := context.Request.Form.Get("content")
	mapName := context.Params.ByName("name")

	service, _ := services.NewMapService(mapName)
	service.InitializeDatabase(content)

	context.String(200, mapName+" map created!\n\n"+content)
}

func TraceRoute(context *gin.Context) {
	var routeForm models.RouteForm
	if !context.BindWith(&routeForm, binding.Form) {
		context.JSON(400, gin.H{"error": "bad request: pending params"})
		return
	}

	mapName := context.Params.ByName("name")
	service, err := services.NewMapService(mapName)
	if err != nil {
		context.JSON(500, gin.H{"error": err})
		return
	}

	response, err := service.TraceBetterRoute(routeForm)
	if err != nil {
		context.JSON(500, gin.H{"error": err})
		return
	}

	context.JSON(200, response)
}
