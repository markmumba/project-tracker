package services

import (
	"github.com/markmumba/project-tracker/database"
	"github.com/markmumba/project-tracker/models"
)

func CreateProject(project *models.Project) error {
    result := database.DB.Create(project)
    return result.Error
}

func GetProject(id uint) (*models.Project, error) {
    var project models.Project
    result := database.DB.First(&project, id)
    return &project, result.Error
}
