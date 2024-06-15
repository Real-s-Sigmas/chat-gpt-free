package handler

import (
	"github.com/gin-gonic/gin"
	"gpt/pkg/service"
)

type Handler struct {
	service *service.Service
}

func NewHandler(service *service.Service) *Handler {
	return &Handler{service: service}
}

func (h *Handler) InitRoutes() *gin.Engine {
	router := gin.New()

	auth := router.Group("/auth")
	{
		auth.POST("/sign-up")
		auth.POST("/sign-in")
		auth.GET("/sign-out")
	}
	user := router.Group("/user")
	{
		user.GET("/user-info")
		user.PUT("/update-info")
		user.PUT("/update-password")
		user.GET("/avatar")
	}
	ai := router.Group("/ai")
	{
		ai.GET("/show-chats")
		ai.GET("/show-chat")
		ai.DELETE("/delete-chat")
	}
	chat := router.Group("/chat")
	{
		chat.POST("/post-message")
		chat.DELETE("/clear-chat")
	}

	return router
}
