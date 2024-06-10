package services

import (
	"github.com/markmumba/project-tracker/database"
	"github.com/markmumba/project-tracker/models"
)

func CreateSubmission(submission *models.Submission) error {
    result := database.DB.Create(submission)
    return result.Error
}

func GetSubmission(id uint) (*models.Submission, error) {
    var submission models.Submission
    result := database.DB.First(&submission, id)
    return &submission, result.Error
}
