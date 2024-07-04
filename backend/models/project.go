package models

import (
    "gorm.io/gorm"
)

type Project struct {
    gorm.Model
    StudentID    uint   `json:"student_id"`
    LecturerID   uint   `json:"lecturer_id"`
    Title        string `json:"title"`
    Description  string `json:"description"`
    StartDate    string `json:"start_date"`
    EndDate      string `json:"end_date"`
    Student      User   `gorm:"foreignKey:StudentID"`
    Lecturer     User   `gorm:"foreignKey:LecturerID"`
}

type ProjectDTO struct {
    StudentID    uint   `json:"student_id"`
    LecturerID   uint   `json:"lecturer_id"`
    LecturerName string `json:"lecturer_name"`
    Title        string `json:"title"`
    Description  string `json:"description"`
    StartDate    string `json:"start_date"`
    EndDate      string `json:"end_date"`
}

func   ProjectToDTO(p *Project) ProjectDTO {
    return ProjectDTO{
        StudentID:    p.StudentID,
        LecturerID:   p.LecturerID,
        LecturerName: p.Lecturer.Name,
        Title:        p.Title,
        Description:  p.Description,
        StartDate:    p.StartDate,
        EndDate:      p.EndDate,
    }
}

func   ProjectToDTOs(projects []Project) []ProjectDTO {
    projectDTOs := make([]ProjectDTO, len(projects))
    for i, project := range projects {
        projectDTOs[i] = ProjectToDTO(&project)
    }
    return projectDTOs
}