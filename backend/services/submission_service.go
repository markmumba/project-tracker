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

func GetAllSubmissionByStudentId(studentId uint) ([]models.Submission, error) {
	user ,err := GetUser(studentId)
	if err != nil {
		return nil, err
	}
	if user.Role.Name != "student" {
		return nil, nil
	}
	var submissions []models.Submission
	result := database.DB.Where("student_id = ?", studentId).Find(&submissions)
	return submissions, result.Error
}

func DeleteSubmission(id uint) error {
	var submission models.Submission
	result := database.DB.Delete(&submission, id)
	return result.Error
}
