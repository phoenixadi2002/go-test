package services

import (
	"go-test/internal/models"
	"go-test/internal/repositories"
)

func GetUserByID(id uint) (*models.User, error) {
    var user models.User
    if err := repositories.DB.First(&user, id).Error; err != nil {
        return nil, err
    }
    return &user, nil
}

func CreateUser(user *models.User) error {
    return repositories.DB.Create(user).Error
}