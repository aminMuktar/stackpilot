package auth

import (
	"net/http"

	"github.com/aminMuktar/stackpilot/internal/database"
	"github.com/aminMuktar/stackpilot/internal/logger"
	"github.com/aminMuktar/stackpilot/internal/models"
	"github.com/aminMuktar/stackpilot/pkg/utils"
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
)

type LoginInput struct {
	Email    string `json:"email" binding:"required,email"`
	Password string `json:"password" binding:"required"`
}

func Login(c *gin.Context) {
	var input LoginInput
	if err := c.ShouldBindJSON(&input); err != nil {
		logger.Log.Error("Error binding JSON", zap.Error(err))
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	var user models.User

	if err := database.DB.Where("email = ?", input.Email).First(&user).Error; err != nil {
		logger.Log.Warn("User not found", zap.String("email", input.Email), zap.Error(err))
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid email or password"})
		return
	}

	// Check password directly
	if !utils.CheckPasswordHash(input.Password, user.Password) {
		logger.Log.Warn("Password mismatch",
			zap.String("email", input.Email),
			zap.String("stored_hash", user.Password))

		c.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid email or password"})
		return
	}

	accessToken, err := utils.GenerateToken(user.ID, 15, []byte("access-secret-key"))
	if err != nil {
		logger.Log.Error("Failed to generate access token", zap.Error(err))
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to generate access token"})
		return
	}

	refreshToken, err := utils.GenerateToken(user.ID, 7*24*60, []byte("refresh-secret-key"))
	if err != nil {
		logger.Log.Error("Failed to generate refresh token", zap.Error(err))
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to generate refresh token"})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"access_token":  accessToken,
		"refresh_token": refreshToken,
	})
}
