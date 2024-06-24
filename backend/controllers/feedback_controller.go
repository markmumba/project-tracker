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
		return c.JSON(http.StatusBadRequest, err.Error())
	}
	feedback.LecturerID = userId
	if err := services.CreateFeedback(&feedback); err != nil {
		return c.JSON(http.StatusInternalServerError, err.Error())
	}

	return c.JSON(http.StatusCreated, models.FeedbackToDTO(&feedback))
}

func GetFeedback(c echo.Context) error {
	var feedbackParams models.Feedback
	err := c.Bind(feedbackParams)
	if err != nil {
		return c.JSON(http.StatusBadRequest, err.Error())
	}
	id := feedbackParams.ID
	feedback, err := services.GetFeedback(id)
	if err != nil {
		return c.JSON(http.StatusNotFound, err.Error())
	}

	return c.JSON(http.StatusOK, models.FeedbackToDTO(feedback))
}

func GetAllFeedback(c echo.Context) error {
	feedbacks, err := services.GetAllFeedback()
	if err != nil {
		return c.JSON(http.StatusInternalServerError, err.Error())
	}
	return c.JSON(http.StatusOK, models.FeedbackToDTOs(feedbacks))
}

func GetFeedbackBySubmissionId(c echo.Context) error {
	var submissionParams models.Submission
	err := c.Bind(submissionParams)
	if err != nil {
		return c.JSON(http.StatusBadRequest, err.Error())
	}
	id := submissionParams.ID
	feedbacks, err := services.GetFeedbackBySubmissionId(uint(id))
	if err != nil {
		return c.JSON(http.StatusNotFound, err.Error())
	}

	return c.JSON(http.StatusOK, models.FeedbackToDTOs(feedbacks))
}

func DeleteFeedback(c echo.Context) error {
    var feedbackParams models.Feedback
    err := c.Bind(feedbackParams)
    if err != nil {
        return c.JSON(http.StatusBadRequest, err.Error())
    }
    id := feedbackParams.ID
	err = services.DeleteFeedback(uint(id))
	if err != nil {
		return c.JSON(http.StatusNotFound, err.Error())
	}

	return c.NoContent(http.StatusNoContent)
}
