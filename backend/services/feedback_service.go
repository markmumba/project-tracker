package services

import (
	"github.com/markmumba/project-tracker/database"
	"github.com/markmumba/project-tracker/models"
)

func CreateFeedback(feedback *models.Feedback) error {
    result := database.DB.Create(feedback)
    return result.Error
}

func GetFeedback(id uint) (*models.Feedback, error) {
    var feedback models.Feedback
    result := database.DB.First(&feedback, id)
    return &feedback, result.Error
}
