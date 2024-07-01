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
func GetProjectsByLecturerId(lecturerId uint) ([]models.Project, error) {
	user, err := GetUser(lecturerId)
	if err != nil {
		return nil, err
	}
	if user.Role.Name != "lecturer" {
		return nil, nil
	}
	var projects []models.Project
	result := database.DB.Where("lecturer_id = ?", lecturerId).Find(&projects)
	return projects, result.Error
}
func UpdateProject(project *models.Project) error {
	result := database.DB.Save(project)
	return result.Error
}

func DeleteProject(id uint) error {
	var project models.Project
	result := database.DB.Delete(&project, id)
	return result.Error
}
