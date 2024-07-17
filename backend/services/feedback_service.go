package services

import (
	"github.com/markmumba/project-tracker/models"
	"github.com/markmumba/project-tracker/repository"
)

type FeedbackService struct {
	FeedbackRepository repository.FeedbackRepository
}

func NewFeedbackService(feedbackRepo repository.FeedbackRepository) *FeedbackService {
	return &FeedbackService{
		FeedbackRepository: feedbackRepo,
	}
}

func (s *FeedbackService) CreateFeedback(feedback *models.Feedback) error {
	return s.FeedbackRepository.CreateFeedback(feedback)
}

func (s *FeedbackService) GetFeedback(id uint) (*models.Feedback, error) {
	return s.FeedbackRepository.GetFeedback(id)
}

func (s *FeedbackService) GetFeedbackByStudent(studentID uint) (*[]models.Feedback, error) {
	return s.FeedbackRepository.GetFeedbackByStudent(studentID)
}

func (s *FeedbackService) GetAllFeedback() ([]models.Feedback, error) {
	return s.FeedbackRepository.GetAllFeedback()
}

func (s *FeedbackService) GetFeedbackBySubmissionId(submissionId uint) ([]models.Feedback, error) {
	return s.FeedbackRepository.GetFeedbackBySubmissionId(submissionId)
}

func (s *FeedbackService) UpdateFeedback(feedback *models.Feedback) error {
	return s.FeedbackRepository.UpdateFeedback(feedback)
}

func (s *FeedbackService) DeleteFeedback(id uint) error {
	return s.FeedbackRepository.DeleteFeedback(id)
}
