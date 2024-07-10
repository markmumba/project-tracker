package models

import (
	"gorm.io/gorm"
)

type Submission struct {
	gorm.Model
	ProjectID      uint    `json:"project_id"`
	StudentID      uint    `json:"student_id"`
	SubmissionDate string  `json:"submission_date"`
	DocumentPath   string  `json:"document_path"`
	Description    string  `json:"description"`
	Project        Project `gorm:"foreignKey:ProjectID" json:"-"`
	Student        User    `gorm:"foreignKey:StudentID" json:"-"`
	Reviewed       bool    `json:"reviewed"`
	ProjectName    string  `json:"project_name" gorm:"-"`
	StudentName    string  `json:"student_name" gorm:"-"`
}
type SubmissionDTO struct {
	SubmissionID   uint   `json:"submission_id"`
	ProjectID      uint   `json:"project_id"`
	ProjectName    string `json:"project_name"`
	StudentID      uint   `json:"student_id"`
	StudentName    string `json:"student_name"`
	SubmissionDate string `json:"submission_date"`
	DocumentPath   string `json:"document_path"`
	Reviewed 	 bool   `json:"reviewed"`
	Description    string `json:"description"`
}

func SubmissionToDTO(s *Submission) SubmissionDTO {
	return SubmissionDTO{
		SubmissionID:   s.ID,
		ProjectID:      s.ProjectID,
		StudentID:      s.StudentID,
		SubmissionDate: s.SubmissionDate,
		DocumentPath:   s.DocumentPath,
		Description:    s.Description,
		Reviewed:       s.Reviewed,
	}
}

func SubmissionToDTOs(submissions []Submission) []SubmissionDTO {
	submissionDTOs := make([]SubmissionDTO, len(submissions))
	for i, submission := range submissions {
		submissionDTOs[i] = SubmissionToDTO(&submission)
	}
	return submissionDTOs
}
