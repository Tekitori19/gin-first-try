package controllers

import (
	"github.com/Tekitori19/gin-first-try/get_started/entity"
	"github.com/Tekitori19/gin-first-try/get_started/service"
	"github.com/gin-gonic/gin"
)

type VideoController interface {
	Save(*gin.Context) entity.Video
	FindAll() []entity.Video
}

type controller struct {
	service service.VideoService
}

func (c *controller) Save(ctx *gin.Context) entity.Video {
	var video entity.Video
	ctx.BindJSON(&video)
	c.service.Save(video)
	return video
}

func (c *controller) FindAll() []entity.Video {
	return c.service.FindAll()
}

func New(service service.VideoService) VideoController {
	return &controller{
		service: service,
	}
}
