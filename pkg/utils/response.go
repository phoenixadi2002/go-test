package utils

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

// ErrorResponse sends a JSON error response
func ErrorResponse(c *gin.Context, status int, msg string) {
    c.JSON(status, gin.H{
        "error": msg,
    })
}

// SuccessResponse sends a JSON success response
func SuccessResponse(c *gin.Context, data any) {
    c.JSON(http.StatusOK, gin.H{
        "data": data,
    })
}
