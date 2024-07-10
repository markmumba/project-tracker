package controllers

import (
	"net/http"
	"strconv"

	"github.com/labstack/echo/v4"
	"github.com/markmumba/project-tracker/helpers"
	"github.com/markmumba/project-tracker/models"
	"github.com/markmumba/project-tracker/services"
)

// TODO : streamline type conversion from frontend to backend and vice versa
type CreateSubmissionRequest struct {
	ProjectID      string `json:"project_id"`
	StudentID      string `json:"student_id"`
	SubmissionDate string `json:"submission_date"`
	DocumentPath   string `json:"document_path"`
	Description    string `json:"description"`
	Reviewed       bool   `json:"reviewed"`
}

func CreateSubmission(c echo.Context) error {

	var request CreateSubmissionRequest
	if err := c.Bind(&request); err != nil {
		return c.JSON(http.StatusBadRequest, err.Error())
	}
	ProjectID, err := strconv.ParseUint(request.ProjectID, 10, 32)
	if err != nil {
		return c.JSON(http.StatusBadRequest, "invalid project_id or student_id")
	}

	StudnetID, err := strconv.ParseUint(request.StudentID, 10, 32)
	if err != nil {
		return c.JSON(http.StatusBadRequest, "invalid project_id or student_id")
	}
	submission := models.Submission{
		ProjectID:      uint(ProjectID),
		StudentID:      uint(StudnetID),
		SubmissionDate: request.SubmissionDate,
		DocumentPath:   request.DocumentPath,
		Description:    request.Description,
		Reviewed: 	 request.Reviewed,
	}

	if err := services.CreateSubmission(&submission); err != nil {
		return c.JSON(http.StatusInternalServerError, err.Error())
	}

	return c.JSON(http.StatusCreated, models.SubmissionToDTO(&submission))
}

func GetSubmission(c echo.Context) error {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		return c.JSON(http.StatusBadRequest, err.Error())
	}
	submission, err := services.GetSubmission(uint(id))
	if err != nil {
		return c.JSON(http.StatusNotFound, err.Error())
	}

	return c.JSON(http.StatusOK, submission)
}

func GetSubmissionsByLecturer(c echo.Context) error {

	userID, err := helpers.ConvertUserID(c, "userId")
	if err != nil {
		return c.JSON(http.StatusInternalServerError, err.Error())
	}
	submissions, err := services.GetSubmissionsByLecturer(userID)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, err.Error())
	}
	return c.JSON(http.StatusOK, submissions)
}

func GetAllSubmissionByStudentId(c echo.Context) error {
	userID, err := helpers.ConvertUserID(c, "userId")
	if err != nil {
		return c.JSON(http.StatusInternalServerError, err.Error())
	}
	submissions, err := services.GetAllSubmissionByStudentId(userID)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, err.Error())
	}
	return c.JSON(http.StatusOK, submissions)
}

func UpdateSubmission(c echo.Context) error {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		return c.JSON(http.StatusBadRequest, err.Error())
	}
	var submission models.Submission
	if err := c.Bind(&submission); err != nil {
		return c.JSON(http.StatusBadRequest, err.Error())
	}

	if err := services.UpdateSubmission(&submission, uint(id)); err != nil {
		return c.JSON(http.StatusInternalServerError, err.Error())
	}

	return c.JSON(http.StatusOK, models.SubmissionToDTO(&submission))
}

func DeleteSubmission(c echo.Context) error {
	var submissionParams models.Submission
	err := c.Bind(&submissionParams)
	if err != nil {
		return c.JSON(http.StatusBadRequest, err.Error())
	}
	id := submissionParams.ID
	err = services.DeleteSubmission(uint(id))
	if err != nil {
		return c.JSON(http.StatusNotFound, err.Error())
	}

	return c.JSON(http.StatusOK, "Submission deleted successfully")
}
