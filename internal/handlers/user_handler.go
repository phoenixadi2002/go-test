package handlers

import (
	"net/http"
	"strconv"

	"go-test/internal/models"
	"go-test/internal/services"

	"github.com/gin-gonic/gin"
)

func GetUser(c *gin.Context) {
    idStr := c.Param("id")
    id, _ := strconv.ParseUint(idStr, 10, 32)
    user, err := services.GetUserByID(uint(id))
    if err != nil {
        c.JSON(http.StatusNotFound, gin.H{"error": "User not found"})
        return
    }
    c.JSON(http.StatusOK, user)
}

func CreateUser(c *gin.Context) {
    var user models.User
    if err := c.ShouldBindJSON(&user); err != nil {
        c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
        return
    }
    if err := services.CreateUser(&user); err != nil {
        c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to create user"})
        return
    }
    c.JSON(http.StatusCreated, user)
}