package models

type Submission struct {
	ID             uint   `gorm:"primaryKey"`
	Description    string `gorm:"not null"`
	SubmissionDate string
	DocumentPath   string  `gorm:"not null"`
	Reviewed       bool    `gorm:"default:false"`
	ProjectID      uint    `gorm:"not null"`
	StudentID      uint    `gorm:"not null"`
	Project        Project `gorm:"foreignKey:ProjectID"`
	Student        User    `gorm:"foreignKey:StudentID"`
}

type SubmissionDTO struct {
	ID          uint       `json:"id"`
	Description string     `json:"description"`
	Reviewed    bool       `json:"reviewed"`
	Project     ProjectDTO `json:"project"`
	Student     UserDTO    `json:"student"`
}

func SubmissionToDTO(s *Submission) SubmissionDTO {
	return SubmissionDTO{
		ID:          s.ID,
		Description: s.Description,
		Reviewed:    s.Reviewed,
		Project:     ProjectToDTO(&s.Project),
		Student:     UserToDTO(&s.Student),
	}
}

func SubmissionToDTOs(submissions []Submission) []SubmissionDTO {
	submissionDTOs := make([]SubmissionDTO, len(submissions))
	for i, submission := range submissions {
		submissionDTOs[i] = SubmissionToDTO(&submission)
	}
	return submissionDTOs
}
