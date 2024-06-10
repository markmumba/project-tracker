package models

import (
	"gorm.io/gorm"
)

type Feedback struct {
	gorm.Model
	SubmissionID uint   `json:"submission_id"`
	LecturerID   uint   `json:"lecturer_id"`
	FeedbackDate string `json:"feedback_date"`
	Comments     string `json:"comments"`
}
