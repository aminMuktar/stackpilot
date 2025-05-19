package main

import (
	"github.com/aminMuktar/stackpilot/config"
	"github.com/aminMuktar/stackpilot/internal/database"
	"github.com/aminMuktar/stackpilot/internal/logger"
	"github.com/gin-gonic/gin"
)

func main() {
	router := gin.Default()

	config.LoadEnv()
	logger.Init(false)
	defer logger.Log.Sync()
	logger.Log.Info("Starting application ...")

	database.Init()
	if err := router.Run(":8000"); err != nil {
		logger.Log.Fatal(`Failed to run the server`)
	}

}
