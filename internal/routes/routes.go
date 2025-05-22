package routes

import (
	"github.com/aminMuktar/stackpilot/internal/routes/auth"
	"github.com/gin-gonic/gin"
)

func RegisterRoutes(router *gin.Engine) {
	auth.RegisterAuthRoutes(router)
}
