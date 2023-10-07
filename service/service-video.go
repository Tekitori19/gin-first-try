package service

import "github.com/Tekitori19/gin-first-try/get_started/entity"

type VideoService interface {
	Save(entity.Video) entity.Video
	FindAll() []entity.Video
}

type videoService struct {
	videos []entity.Video
}

func (v *videoService) Save(newVideo entity.Video) entity.Video {
	v.videos = append(v.videos, newVideo)
	return newVideo
}

func (v *videoService) FindAll() []entity.Video {
	return v.videos
}

func New() VideoService {
	return &videoService{}
}
