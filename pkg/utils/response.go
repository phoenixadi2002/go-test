package utils

import (
    "net/http"

    "github.com/gin-gonic/gin"
)

func ErrorResponse(c *gin.Context, status int, msg string) {
    c.JSON(status, gin.H{"error": msg})
}