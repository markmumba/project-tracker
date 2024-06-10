package services

import (
	"github.com/markmumba/project-tracker/database"
	"github.com/markmumba/project-tracker/models"
)

func CreateUser(user *models.User) error {
    result := database.DB.Create(user)
    return result.Error
}

func GetUser(id uint) (*models.User, error) {
    var user models.User
    result := database.DB.First(&user, id)
    return &user, result.Error
}
