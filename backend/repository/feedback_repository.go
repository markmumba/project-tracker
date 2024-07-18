package repository

import (
	"github.com/markmumba/project-tracker/database"
	"github.com/markmumba/project-tracker/models"
)

type FeedbackRepositoryImpl struct{}

func NewFeedbackRepository() FeedbackRepository {
	return &FeedbackRepositoryImpl{}
}

func (repo *FeedbackRepositoryImpl) CreateFeedback(feedback *models.Feedback) error {
	result := database.DB.Create(feedback)
	return result.Error
}

func (repo *FeedbackRepositoryImpl) GetFeedback(id uint) (*models.Feedback, error) {
	var feedback models.Feedback
	result := database.DB.First(&feedback, id)
	if result.Error != nil {
		return nil, result.Error
	}
	return &feedback, nil
}

func (repo *FeedbackRepositoryImpl) GetFeedbackByStudent(studentID uint) (*[]models.Feedback, error) {
    var feedbacks []models.Feedback
    
    err := database.DB.
        Preload("Submission.Project").
        Preload("Submission.Student").
        Preload("Lecturer").
        Joins("JOIN submissions ON feedbacks.submission_id = submissions.id").
        Where("submissions.student_id = ?", studentID).
        Find(&feedbacks).Error

    if err != nil {
        return nil, err
    }

    return &feedbacks, nil
}

func (repo *FeedbackRepositoryImpl) GetAllFeedback() ([]models.Feedback, error) {
	var feedbacks []models.Feedback
	result := database.DB.Find(&feedbacks)
	return feedbacks, result.Error
}

func (repo *FeedbackRepositoryImpl) GetFeedbackBySubmissionId(submissionId uint) ([]models.Feedback, error) {
	var feedbacks []models.Feedback
	result := database.DB.Where("submission_id = ?", submissionId).Find(&feedbacks)
	return feedbacks, result.Error
}

func (repo *FeedbackRepositoryImpl) UpdateFeedback(feedback *models.Feedback) error {
	result := database.DB.Save(feedback)
	return result.Error
}

func (repo *FeedbackRepositoryImpl) DeleteFeedback(id uint) error {
	result := database.DB.Delete(&models.Feedback{}, id)
	return result.Error
}
