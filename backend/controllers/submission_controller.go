package controllers

import (
	"net/http"

	"github.com/labstack/echo/v4"
	"github.com/markmumba/project-tracker/models"
	"github.com/markmumba/project-tracker/services"
)

func CreateSubmission(c echo.Context) error {
	userId := c.Get("userId").(uint)
	var submission models.Submission
	if err := c.Bind(&submission); err != nil {
		return c.JSON(http.StatusBadRequest, err.Error())
	}
	submission.StudentID = userId
	if err := services.CreateSubmission(&submission); err != nil {
		return c.JSON(http.StatusInternalServerError, err.Error())
	}

	return c.JSON(http.StatusCreated, models.SubmissionToDTO(&submission))
}

func GetSubmission(c echo.Context) error {
	var submissionParams models.Submission

	err := c.Bind(&submissionParams)
	if err != nil {
		return c.JSON(http.StatusBadRequest, err.Error())
	}	
	id := submissionParams.ID
	submission, err := services.GetSubmission(uint(id))
	if err != nil {
		return c.JSON(http.StatusNotFound, err.Error())
	}

	return c.JSON(http.StatusOK, models.SubmissionToDTO(submission))
}

func GetAllSubmissionByStudentId(c echo.Context) error {
	userId := c.Get("userId").(uint)
	submissions, err := services.GetAllSubmissionByStudentId(userId)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, err.Error())
	}
	return c.JSON(http.StatusOK, models.SubmissionToDTOs(submissions))
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
