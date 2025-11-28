package main

import (
    "fmt"

    "github.com/gin-gonic/gin"
    "go.uber.org/zap"
    "go/internal/config"
    "go/internal/handlers"
    "go/internal/repositories"
)

func main() {
    cfg := config.LoadConfig()

    logger, _ := zap.NewProduction()
    defer logger.Sync()

    repositories.ConnectDB(cfg)
    repositories.DB.AutoMigrate() // Add models here

    r := gin.Default()
    handlers.SetupRoutes(r, logger)

    r.Run(fmt.Sprintf(":%s", cfg.Port))
}