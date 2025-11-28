package main

import (
	"fmt"

	"go-test/internal/config"
	"go-test/internal/handlers"
	"go-test/internal/repositories"

	"go-test/internal/models"

	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
)

func main() {
    cfg := config.LoadConfig()

    logger, _ := zap.NewProduction()
    defer logger.Sync()

    repositories.ConnectDB(cfg)
    repositories.DB.AutoMigrate(&models.User{}) // Add models here

    r := gin.Default()
    handlers.SetupRoutes(r, logger)

    r.Run(fmt.Sprintf(":%s", cfg.Port))
}