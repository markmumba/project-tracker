package models

import (
    "gorm.io/gorm"
)

type Submission struct {
    gorm.Model
    ProjectID     uint   `json:"project_id"`
    StudentID     uint   `json:"student_id"`
    SubmissionDate string `json:"submission_date"`
    DocumentPath  string `json:"document_path"`
    Description   string `json:"description"`
    Project       Project `gorm:"foreignKey:ProjectID"`
    Student       User    `gorm:"foreignKey:StudentID"`
}
type SubmissionDTO struct {
    SubmissionID  uint   `json:"submission_id"`
    ProjectID     uint   `json:"project_id"`
    StudentID     uint   `json:"student_id"`
    SubmissionDate string `json:"submission_date"`
    DocumentPath  string `json:"document_path"`
    Description   string `json:"description"`
}

func  SubmissionToDTO(s *Submission) SubmissionDTO {
    return SubmissionDTO{
        SubmissionID:  s.ID,
        ProjectID:     s.ProjectID,
        StudentID:     s.StudentID,
        SubmissionDate: s.SubmissionDate,
        DocumentPath:  s.DocumentPath,
        Description:   s.Description,
    }
}

func  SubmissionToDTOs(submissions []Submission) []SubmissionDTO {
    submissionDTOs := make([]SubmissionDTO, len(submissions))
    for i, submission := range submissions {
        submissionDTOs[i] = SubmissionToDTO(&submission)
    }
    return submissionDTOs
}
 