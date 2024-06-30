package controllers

import (
	"net/http"

	"github.com/labstack/echo/v4"
	"github.com/markmumba/project-tracker/models"
	"github.com/markmumba/project-tracker/services"
)

func CreateFeedback(c echo.Context) error {
	userId := c.Get("userId").(uint)
	var feedback models.Feedback
	if err := c.Bind(&feedback); err != nil {
		return c.Render(http.StatusBadRequest, "error.html", map[string]interface{}{
			"error": err.Error(),
		})
	}
	feedback.LecturerID = userId
	if err := services.CreateFeedback(&feedback); err != nil {
		return c.Render(http.StatusInternalServerError, "error.html", map[string]interface{}{
			"error": err.Error(),
		})
	}

	return c.Redirect(http.StatusFound, "/feedbacks")
}

func GetFeedback(c echo.Context) error {
	var feedbackParams models.Feedback
	err := c.Bind(&feedbackParams)
	if err != nil {
		return c.Render(http.StatusBadRequest, "error.html", map[string]interface{}{
			"error": err.Error(),
		})
	}
	id := feedbackParams.ID
	feedback, err := services.GetFeedback(id)
	if err != nil {
		return c.Render(http.StatusNotFound, "error.html", map[string]interface{}{
			"error": err.Error(),
		})
	}

	return c.Render(http.StatusOK, "feedback.html", models.FeedbackToDTO(feedback))
}

func GetAllFeedback(c echo.Context) error {
	feedbacks, err := services.GetAllFeedback()
	if err != nil {
		return c.Render(http.StatusInternalServerError, "error.html", map[string]interface{}{
			"error": err.Error(),
		})
	}
	return c.Render(http.StatusOK, "feedbacks.html", models.FeedbackToDTOs(feedbacks))
}

func GetFeedbackBySubmissionId(c echo.Context) error {
	var submissionParams models.Submission
	err := c.Bind(&submissionParams)
	if err != nil {
		return c.Render(http.StatusBadRequest, "error.html", map[string]interface{}{
			"error": err.Error(),
		})
	}
	id := submissionParams.ID
	feedbacks, err := services.GetFeedbackBySubmissionId(uint(id))
	if err != nil {
		return c.Render(http.StatusNotFound, "error.html", map[string]interface{}{
			"error": err.Error(),
		})
	}

	return c.Render(http.StatusOK, "feedbacks.html", models.FeedbackToDTOs(feedbacks))
}

func DeleteFeedback(c echo.Context) error {
	var feedbackParams models.Feedback
	err := c.Bind(&feedbackParams)
	if err != nil {
		return c.Render(http.StatusBadRequest, "error.html", map[string]interface{}{
			"error": err.Error(),
		})
	}
	id := feedbackParams.ID
	err = services.DeleteFeedback(uint(id))
	if err != nil {
		return c.Render(http.StatusNotFound, "error.html", map[string]interface{}{
			"error": err.Error(),
		})
	}

	return c.Render(http.StatusOK, "delete_success.html", map[string]interface{}{
		"message": "Feedback deleted successfully",
	})
}
