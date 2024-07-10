package services

import (
	"github.com/markmumba/project-tracker/database"
	"github.com/markmumba/project-tracker/models"
)

func CreateSubmission(submission *models.Submission) error {
	result := database.DB.Create(submission)
	if result.Error != nil {
		return result.Error
	}
	err := database.DB.Preload("Project").First(submission, submission.ID).Error
	if err != nil {
		return err
	}

	return nil
}

func GetSubmission(id uint) (*models.Submission, error) {
	var submission models.Submission
	result := database.DB.Preload("Project").Preload("Student").First(&submission, id)
	if result.Error != nil {
		return nil, result.Error
	}

	// Populate the ProjectName and StudentName fields
	submission.ProjectName = submission.Project.Title
	submission.StudentName = submission.Student.Name

	return &submission, nil
}

func GetAllSubmissionByStudentId(studentId uint) ([]models.Submission, error) {
	user, err := GetUser(studentId)
	if err != nil {
		return nil, err
	}
	if user.RoleID != 2 {
		return nil, nil
	}
	var submissions []models.Submission
	result := database.DB.Where("student_id = ?", studentId).Find(&submissions)
	return submissions, result.Error
}

func GetSubmissionsByLecturer(lecturerID uint) ([]models.SubmissionDTO, error) {
	var submissions []models.SubmissionDTO
	err := database.DB.Table("submissions").
		Select("submissions.id as submission_id, projects.id as project_id, users.id as student_id, submissions.submission_date, submissions.document_path, submissions.description, projects.title as project_name, users.name as student_name, submissions.reviewed").
		Joins("JOIN projects ON projects.id = submissions.project_id").
		Joins("JOIN users ON users.id = submissions.student_id").
		Where("projects.lecturer_id = ?", lecturerID).
		Scan(&submissions).Error
	if err != nil {
		return nil, err
	}
	return submissions, nil
}

func UpdateSubmission(submission *models.Submission, id uint) error {
	// Find the submission by ID
	var existingSubmission models.Submission
	result := database.DB.First(&existingSubmission, id)
	if result.Error != nil {
		return result.Error
	}

	// Update fields of the existing submission
	existingSubmission.Description = submission.Description
	existingSubmission.SubmissionDate = submission.SubmissionDate
	existingSubmission.ProjectName = submission.ProjectName
	existingSubmission.StudentName = submission.StudentName
	existingSubmission.Reviewed = submission.Reviewed // Assuming Reviewed is a field in models.Submission

	// Save the updated submission
	result = database.DB.Save(&existingSubmission)
	return result.Error
}

func DeleteSubmission(id uint) error {
	var submission models.Submission
	result := database.DB.Delete(&submission, id)
	return result.Error
}
