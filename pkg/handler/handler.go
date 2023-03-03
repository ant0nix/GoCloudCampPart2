package handler

import (
	"github.com/ant0nix/GoCloudCampPart2/pkg/service"
	"github.com/gin-gonic/gin"
)

type Handler struct {
	services *service.Service
}

func NewHandler(service *service.Service) *Handler {
	return &Handler{services: service}
}

func (h *Handler) InitRoutes() *gin.Engine {
	router := gin.New()
	adder := router.Group("/playlist")
	{
		adder.PATCH("/change-song/:id", h.ChangeSong)
		adder.POST("/add-song", h.AddSong)
		adder.DELETE("/delete-song/:id", h.DeleteSong)

	}
	play := router.Group("/play")
	{
		play.GET("/", h.Start)
		play.GET("/next", h.Next)
		play.GET("/prev", h.Prev)
		play.GET("/paused", h.Pause)
		play.GET("/stop", h.StopPlaylist)
	}
	return router
}
