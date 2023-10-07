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
	gin.DefaultWriter = io.MultiWriter(f, os.Stdout)
}

func main() {

	setupLogOutput()

	server := gin.New()

	server.Static("/css", "./templates/css")
	server.LoadHTMLGlob("templates/*.html")

	server.Use(
		gin.Recovery(),
		middlewares.Logger(),
		middlewares.BasicAuth(),
	)

	apiRoutes := server.Group("/api")
	{
		apiRoutes.GET("/videos", func(ctx *gin.Context) {
			ctx.JSON(http.StatusAccepted, videoController.FindAll())
		})

		apiRoutes.POST("/videos", func(ctx *gin.Context) {
			if err := videoController.Save(ctx); err != nil {
				ctx.JSON(http.StatusBadRequest, gin.H{
					"err": err.Error(),
				})
			} else {
				ctx.JSON(http.StatusAccepted, gin.H{
					"message":"valid",
				})
			}
		})
	}

	viewRoutes := server.Group("/views")
	{
		viewRoutes.GET("/videos", videoController.ShowAll)
	}

	server.Run(":8080")
}