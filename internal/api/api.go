package api

import (
	"fmt"

	"github.com/LamichhaneBibek/go-template/internal/api/routes"
	"github.com/LamichhaneBibek/go-template/internal/config"
	"github.com/gin-gonic/gin"
)

func InitServer() {
	config := config.GetConfig()
	r := gin.New()
	r.Use(gin.Logger(), gin.Recovery())
	v1 := r.Group("api/v1")
	{
		auth := v1.Group("auth")
		routes.Auth(auth, config)

	}

	r.Run(fmt.Sprintf(":%s", config.Server.InternalPort))
}
