package services

import (
	"github.com/markmumba/project-tracker/database"
	"github.com/markmumba/project-tracker/models"
	"gorm.io/gorm"
)

func CreateSubmission(submission *models.Submission) error {
	result := database.DB.Create(submission)
	if result.Error != nil {
		return result.Error
	}
	err := database.DB.Preload("Project", func(db *gorm.DB) *gorm.DB {
		return db.Select("ID", "StudentID", "LecturerID", "Title", "Description", "StartDate", "EndDate")
	}).Preload("Student", func(db *gorm.DB) *gorm.DB {
		return db.Select("ID", "Name", "Email", "Role")
	}).First(submission, submission.ID).Error

	if err != nil {
		return err
	}

	return nil
}

func GetSubmission(id uint) (*models.Submission, error) {
	var submission models.Submission
	result := database.DB.First(&submission, id)
	return &submission, result.Error
}
