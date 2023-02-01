package http

import (
	"github.com/gin-gonic/gin"
	"practice/internal/service"
)

type Handler struct {
	services *service.Services
}

func NewHandler(services *service.Services) *Handler {
	return &Handler{
		services: services,
	}
}

func (h *Handler) InitRoutes() *gin.Engine {
	router := gin.Default()

	api := router.Group("/api")
	{
		books := api.Group("/books")
		{
			books.POST("/")
			books.GET("/")
			books.GET("/:id")
			books.PUT("/:id")
			books.DELETE("/:id")
		}
	}

	return router
}
