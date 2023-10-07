package main

import (
	"io"
	"net/http"
	"os"

	"github.com/Tekitori19/gin-first-try/get_started/controllers"
	"github.com/Tekitori19/gin-first-try/get_started/middlewares"
	"github.com/Tekitori19/gin-first-try/get_started/service"
	"github.com/gin-gonic/gin"
)

var (
	videoService service.VideoService = service.New()
	videoController controllers.VideoController = controllers.New(videoService)
)

func setupLogOutput() {
	f, _ := os.Create("gin.log")
	defer f.Close()
	gin.DefaultWriter = io.MultiWriter(f, os.Stdout)
}

func main() {

	setupLogOutput()

	server := gin.New()

	server.Use(
		gin.Recovery(),
		middlewares.Logger(),
		middlewares.BasicAuth(),
	)

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