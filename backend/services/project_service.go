package services

import (
	"github.com/markmumba/project-tracker/database"
	"github.com/markmumba/project-tracker/models"
	"gorm.io/gorm"
)

func CreateProject(project *models.Project) error {
	result := database.DB.Create(project)
	if result.Error != nil {
		return result.Error
	}

	err := database.DB.Preload("Student", func(db *gorm.DB) *gorm.DB {
		return db.Select("ID", "Name", "Email", "Role")
	}).Preload("Lecturer", func(db *gorm.DB) *gorm.DB {
		return db.Select("ID", "Name", "Email", "Role")
	}).First(project, project.ID).Error

	if err != nil {
		return err
	}


	return nil
}
func GetProject(id uint) (*models.Project, error) {
	var project models.Project
	result := database.DB.First(&project, id)
	return &project, result.Error
}
