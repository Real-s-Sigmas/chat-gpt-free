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
		auth.POST("/sign-up", h.SignUp)
		auth.POST("/sign-in", h.SignIn)
		auth.GET("/sign-out", h.SignOut)
	}
	user := router.Group("/user")
	{
		user.GET("/user-info", h.GetUserInfo)
		user.PUT("/update-info", h.UpdateUserInfo)
		user.PUT("/update-password", h.UpdatePassword)
		user.GET("/avatar", h.ShowAvatar)
	}
	ai := router.Group("/ai")
	{
		ai.GET("/show-chats", h.ShowChats)
		ai.GET("/show-chat", h.ShowChat)
		ai.DELETE("/delete-chat", h.DeleteChat)
	}
	chat := router.Group("/chat")
	{
		chat.POST("/post-message", h.PostMessage)
		chat.DELETE("/clear-chat", h.ClearChat)
	}

	return router
}
