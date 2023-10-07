package main

import (
	"net/http"

	"github.com/Tekitori19/gin-first-try/get_started/controllers"
	"github.com/Tekitori19/gin-first-try/get_started/service"
	"github.com/gin-gonic/gin"
)

var (
	videoService service.VideoService = service.New()
	videoController controllers.VideoController = controllers.New(videoService)
)

func main() {
	server := gin.Default()

	server.GET("/hello", func(ctx *gin.Context) {
		ctx.JSON(http.StatusAccepted, gin.H{
			"Phan": "Duong Dinh",
		})
	})

	server.GET("/videos", func(ctx *gin.Context) {
		ctx.JSON(http.StatusAccepted, videoController.FindAll())
	})

	server.POST("/videos", func(ctx *gin.Context) {
		ctx.JSON(http.StatusAccepted, videoController.Save(ctx))
	})

	server.Run(":8080")
}