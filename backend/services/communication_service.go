package services

import (
	"github.com/markmumba/project-tracker/database"
	"github.com/markmumba/project-tracker/models"
)

func CreateMessage(message *models.CommunicationHistory) error {
    result := database.DB.Create(message)
    return result.Error
}

func GetCommunicationHistory(projectID uint) ([]models.CommunicationHistory, error) {
    var messages []models.CommunicationHistory
    result := database.DB.Where("project_id = ?", projectID).Find(&messages)
    return messages, result.Error
}
