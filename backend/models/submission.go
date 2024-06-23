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

 