package models

import (
	"gorm.io/gorm"
)

type Project struct {
	gorm.Model
	StudentID   uint   `json:"student_id"`
	LecturerID  uint   `json:"lecturer_id"`
	Title       string `json:"title"`
	Description string `json:"description"`
	StartDate   string `json:"start_date"`
	EndDate     string `json:"end_date"`
}
