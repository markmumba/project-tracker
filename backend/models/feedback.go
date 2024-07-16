package models

type Feedback struct {
	ID           uint   `gorm:"primaryKey"`
	Comment      string `gorm:"not null"`
	FeedbackDate string
	SubmissionID uint       `gorm:"not null"`
	LecturerID   uint       `gorm:"not null"`
	Submission   Submission `gorm:"foreignKey:SubmissionID"`
	Lecturer     User       `gorm:"foreignKey:LecturerID"`
}

type FeedbackResponse struct {
	FeedbackID     uint   `json:"feedback_id"`
	FeedbackDate   string `json:"feedback_date"`
	Comments       string `json:"comments"`
	SubmissionID   uint   `json:"submission_id"`
	SubmissionDate string `json:"submission_date"`
	DocumentPath   string `json:"document_path"`
	Description    string `json:"description"`
	StudentName    string `json:"student_name"`
	StudentEmail   string `json:"student_email"`
}
type FeedbackDTO struct {
	ID           uint          `json:"id"`
	Comment      string        `json:"comment"`
	FeedbackDate string        `json:"feedback_date"`
	Submission   SubmissionDTO `json:"submission"`
	Lecturer     UserDTO       `json:"lecturer"`
}

func FeedbackToDTO(f *Feedback) FeedbackDTO {
	return FeedbackDTO{
		ID:           f.ID,
		Comment:      f.Comment,
		FeedbackDate: f.FeedbackDate,
		Submission:   SubmissionToDTO(&f.Submission),
		Lecturer:     UserToDTO(&f.Lecturer),
	}
}

func FeedbackToDTOs(feedbacks []Feedback) []FeedbackDTO {
	feedbackDTOs := make([]FeedbackDTO, len(feedbacks))
	for i, feedback := range feedbacks {
		feedbackDTOs[i] = FeedbackToDTO(&feedback)
	}
	return feedbackDTOs
}
