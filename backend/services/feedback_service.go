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

func GetFeedbackByStudent(id uint) (*[]models.Feedback, error) {
    var feedback []models.Feedback
    err := database.DB.Table("feedbacks").
        Select("feedbacks.id, feedbacks.feedback_date, feedbacks.comments, submissions.id as submission_id, submissions.submission_date, submissions.document_path, submissions.description, submissions.project_name, submissions.student_name").
        Joins("JOIN submissions ON feedbacks.submission_id = submissions.id").
        Where("submissions.student_id = ?", id).
        Find(&feedback).Error

    if err != nil {
        return nil, err
    }

    return &feedback, nil
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
