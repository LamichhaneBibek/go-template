package routes

import (
	"github.com/LamichhaneBibek/go-template/internal/api/handlers"
	"github.com/LamichhaneBibek/go-template/internal/config"
	"github.com/gin-gonic/gin"
)

func Auth(router *gin.RouterGroup, config *config.Config) {
	public := handlers.NewUsersHandler(config)
	router.POST("/login", public.Login)
	router.POST("/register", public.Register)
}
