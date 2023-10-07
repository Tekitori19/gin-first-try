package controllers

import (
	"net/http"

	"github.com/Tekitori19/gin-first-try/get_started/entity"
	"github.com/Tekitori19/gin-first-try/get_started/service"
	"github.com/Tekitori19/gin-first-try/get_started/validators"
	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
)

type VideoController interface {
	Save(*gin.Context) error
	FindAll() []entity.Video
	ShowAll(*gin.Context)
}

type controller struct {
	service service.VideoService
}

func (c *controller) Save(ctx *gin.Context) error {
	var video entity.Video
	if err := ctx.ShouldBindJSON(&video); err != nil {
		return err
	}
	if err := validate.Struct(video); err != nil {
		return err
	}
	c.service.Save(video)
	return nil
}

// tự validate dữ liệu
func (c *controller) FindAll() []entity.Video {
	return c.service.FindAll()
}

var validate *validator.Validate

func New(service service.VideoService) VideoController {
	validate = validator.New()
	validate.RegisterValidation("is-cool", validators.ValidateCoolTitle)
	return &controller{
		service: service,
	}
}

func (c *controller) ShowAll(ctx *gin.Context)  {
	videos := c.service.FindAll()
	data := gin.H {
		"title": "Gin",
		"videos": videos,
	}
	ctx.HTML(http.StatusOK, "index.html", data)
}