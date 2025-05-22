package auth

import (
	"github.com/aminMuktar/stackpilot/internal/auth"
	"github.com/gin-gonic/gin"
)

func RegisterAuthRoutes(router *gin.Engine) {
	authGroup := router.Group("/auth")
	{
		authGroup.POST("/register", auth.Register)
		authGroup.POST("/login", auth.Login)
	}
}
