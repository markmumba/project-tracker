package controllers

import (
	"net/http"

	"github.com/labstack/echo/v4"
	"github.com/markmumba/project-tracker/helpers"
	"github.com/markmumba/project-tracker/models"
	"github.com/markmumba/project-tracker/services"
)

// TODO : get the latest feedback for the student

func CreateFeedback(c echo.Context) error {

	userID, err := helpers.ConvertUserID(c, "userId")
	if err != nil {
		return c.JSON(http.StatusInternalServerError, err.Error())
	}
	var feedback models.Feedback
	if err := c.Bind(&feedback); err != nil {
		return c.JSON(http.StatusBadRequest, err.Error())
	}
	feedback.LecturerID = userID
	if err := services.CreateFeedback(&feedback); err != nil {
		return c.JSON(http.StatusInternalServerError, err.Error())
	}

	return c.JSON(http.StatusCreated, models.FeedbackToDTO(&feedback))
}

func GetFeedbackByStudent(c echo.Context) error {
	userID, err := helpers.ConvertUserID(c, "userId")
	if err != nil {
		return c.JSON(http.StatusInternalServerError, err.Error())
	}
	feedback, err := services.GetFeedbackByStudent(uint(userID))
	if err != nil {
		return c.JSON(http.StatusNotFound, err.Error())
	}

	return c.JSON(http.StatusOK, feedback)
}

func GetAllFeedback(c echo.Context) error {
	feedbacks, err := services.GetAllFeedback()
	if err != nil {
		return c.JSON(http.StatusInternalServerError, err.Error())
	}
	return c.JSON(http.StatusOK, models.FeedbackToDTOs(feedbacks))
}
func UpdateFeedback(c echo.Context) error {
	var feedback models.Feedback
	if err := c.Bind(&feedback); err != nil {
		return c.JSON(http.StatusBadRequest, err.Error())
	}

	if err := services.UpdateFeedback(&feedback); err != nil {
		return c.JSON(http.StatusInternalServerError, err.Error())
	}

	return c.JSON(http.StatusOK, models.FeedbackToDTO(&feedback))
}

func DeleteFeedback(c echo.Context) error {
	var feedbackParams models.Feedback
	err := c.Bind(&feedbackParams)
	if err != nil {
		return c.JSON(http.StatusBadRequest, err.Error())
	}
	id := feedbackParams.ID
	err = services.DeleteFeedback(uint(id))
	if err != nil {
		return c.JSON(http.StatusNotFound, err.Error())
	}
	return c.JSON(http.StatusOK, "Feedback deleted successfully")
}
