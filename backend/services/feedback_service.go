package services

import (
	"github.com/markmumba/project-tracker/database"
	"github.com/markmumba/project-tracker/models"
)

func CreateFeedback(feedback *models.Feedback) error {
	result := database.DB.Create(feedback)
	if result.Error != nil {
		return result.Error
	}
	return nil
}
func GetFeedback(id uint) (*models.Feedback, error) {
	var feedback models.Feedback
	result := database.DB.First(&feedback, id)
	return &feedback, result.Error
}

func GetAllFeedback() ([]models.Feedback, error) {
	var feedbacks []models.Feedback
	result := database.DB.Find(&feedbacks)
	return feedbacks, result.Error
}
func GetFeedbackBySubmissionId(submissionId uint) ([]models.Feedback, error) {
	var feedbacks []models.Feedback
	result := database.DB.Where("submission_id = ?", submissionId).Find(&feedbacks)
	return feedbacks, result.Error
}

func UpdateFeedback(feedback *models.Feedback) error {
	result := database.DB.Save(feedback)
	return result.Error
}

func DeleteFeedback(id uint) error {
	result := database.DB.Delete(&models.Feedback{}, id)
	return result.Error
}
