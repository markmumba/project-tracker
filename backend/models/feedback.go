package models

import (
	"gorm.io/gorm"
)

type Feedback struct {
	gorm.Model
	SubmissionID uint       `json:"submission_id"`
	LecturerID   uint       `json:"lecturer_id"`
	FeedbackDate string     `json:"feedback_date"`
	Comments     string     `json:"comments"`
	Submission   Submission `gorm:"foreignKey:SubmissionID" json:"submission" `
	Lecturer     User       `gorm:"foreignKey:LecturerID"  `
}

type FeedbackDTO struct {
	SubmissionID uint   `json:"submission_id"`
	LecturerID   uint   `json:"lecturer_id"`
	FeedbackDate string `json:"feedback_date"`
	Comments     string `json:"comments"`
}

func FeedbackToDTO(f *Feedback) FeedbackDTO {
	return FeedbackDTO{
		SubmissionID: f.SubmissionID,
		LecturerID:   f.LecturerID,
		FeedbackDate: f.FeedbackDate,
		Comments:     f.Comments,
	}
}

func FeedbackToDTOs(feedbacks []Feedback) []FeedbackDTO {
	feedbackDTOs := make([]FeedbackDTO, len(feedbacks))
	for i, feedback := range feedbacks {
		feedbackDTOs[i] = FeedbackToDTO(&feedback)
	}
	return feedbackDTOs
}
