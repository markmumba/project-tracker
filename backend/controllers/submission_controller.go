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
		return c.Render(http.StatusBadRequest, "error.html", map[string]interface{}{
			"error": err.Error(),
		})
	}
	submission.StudentID = userId
	if err := services.CreateSubmission(&submission); err != nil {
		return c.Render(http.StatusInternalServerError, "error.html", map[string]interface{}{
			"error": err.Error(),
		})
	}

	return c.Redirect(http.StatusFound, "/submissions")
}

func GetSubmission(c echo.Context) error {
	var submissionParams models.Submission

	err := c.Bind(&submissionParams)
	if err != nil {
		return c.Render(http.StatusBadRequest, "error.html", map[string]interface{}{
			"error": err.Error(),
		})
	}

	id := submissionParams.ID
	submission, err := services.GetSubmission(id)
	if err != nil {
		return c.Render(http.StatusNotFound, "error.html", map[string]interface{}{
			"error": err.Error(),
		})
	}

	return c.Render(http.StatusOK, "submission.html", models.SubmissionToDTO(submission))
}

func GetAllSubmissionByStudentId(c echo.Context) error {
	userId := c.Get("userId").(uint)
	submissions, err := services.GetAllSubmissionByStudentId(userId)
	if err != nil {
		return c.Render(http.StatusInternalServerError, "error.html", map[string]interface{}{
			"error": err.Error(),
		})
	}
	return c.Render(http.StatusOK, "submissions.html", models.SubmissionToDTOs(submissions))
}

func DeleteSubmission(c echo.Context) error {
	var submissionParams models.Submission
	err := c.Bind(&submissionParams)
	if err != nil {
		return c.Render(http.StatusBadRequest, "error.html", map[string]interface{}{
			"error": err.Error(),
		})
	}
	id := submissionParams.ID
	err = services.DeleteSubmission(id)
	if err != nil {
		return c.Render(http.StatusNotFound, "error.html", map[string]interface{}{
			"error": err.Error(),
		})
	}

	return c.Render(http.StatusOK, "delete_success.html", map[string]interface{}{
		"message": "Submission deleted successfully",
	})
}
