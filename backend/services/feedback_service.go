package services

import (
	"github.com/markmumba/project-tracker/database"
	"github.com/markmumba/project-tracker/models"
	"gorm.io/gorm"
)

func CreateFeedback(feedback *models.Feedback) error {
    result := database.DB.Create(feedback)
    if result.Error != nil {
        return result.Error
    }

    err := database.DB.Preload("Submission").Preload("Lecturer", func(db *gorm.DB) *gorm.DB {
        return db.Select("ID", "Name", "Email", "Role")
    }).First(feedback, feedback.ID).Error

    if err != nil {
        return err
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

func DeleteFeedback(id uint) error {
    var feedback models.Feedback
    result := database.DB.Delete(&feedback, id)
    return result.Error
}