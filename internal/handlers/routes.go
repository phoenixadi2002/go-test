package handlers

import (
    "github.com/gin-gonic/gin"
    "go.uber.org/zap"
    "go/internal/middleware"
)

func SetupRoutes(r *gin.Engine, logger *zap.Logger) {
    r.Use(middleware.Logger(logger))
    // r.Use(gin.Recovery()) // Built-in recovery

    // API version group
    api := r.Group("/api/v1")

    // User routes
    users := api.Group("/users")
    {
        users.GET("/:id", GetUser)
        users.POST("/", CreateUser)
    }

    // Add more groups for other features (e.g., /products, /orders)
}