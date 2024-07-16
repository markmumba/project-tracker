package controllers

import (
	"net/http"

	"github.com/labstack/echo/v4"
	"github.com/markmumba/project-tracker/helpers"
	"github.com/markmumba/project-tracker/models"
	"github.com/markmumba/project-tracker/services"
)

// TODO : get the latest feedback for the student

type FeedbackController struct {
	feedbackService *services.FeedbackService
}

func NewFeedbackController(feedbackService *services.FeedbackService) *FeedbackController {
	return &FeedbackController{
		feedbackService: feedbackService,
	}
}

func (fc *FeedbackController) CreateFeedback(c echo.Context) error {
	userID, err := helpers.ConvertUserID(c, "userId")
	if err != nil {
		return c.JSON(http.StatusInternalServerError, err.Error())
	}
	var feedback models.Feedback
	if err := c.Bind(&feedback); err != nil {
		return c.JSON(http.StatusBadRequest, err.Error())
	}
	feedback.LecturerID = userID
	if err := fc.feedbackService.CreateFeedback(&feedback); err != nil {
		return c.JSON(http.StatusInternalServerError, err.Error())
	}

	return c.JSON(http.StatusCreated, models.FeedbackToDTO(&feedback))
}

func (fc *FeedbackController) GetFeedbackByStudent(c echo.Context) error {
	userID, err := helpers.ConvertUserID(c, "userId")
	if err != nil {
		return c.JSON(http.StatusInternalServerError, err.Error())
	}
	feedback, err := fc.feedbackService.GetFeedbackByStudent(uint(userID))
	if err != nil {
		return c.JSON(http.StatusNotFound, err.Error())
	}

	return c.JSON(http.StatusOK, feedback)
}

func (fc *FeedbackController) GetAllFeedback(c echo.Context) error {
	feedbacks, err := fc.feedbackService.GetAllFeedback()
	if err != nil {
		return c.JSON(http.StatusInternalServerError, err.Error())
	}
	return c.JSON(http.StatusOK, models.FeedbackToDTOs(feedbacks))
}

func (fc *FeedbackController) UpdateFeedback(c echo.Context) error {
	var feedback models.Feedback
	if err := c.Bind(&feedback); err != nil {
		return c.JSON(http.StatusBadRequest, err.Error())
	}

	if err := fc.feedbackService.UpdateFeedback(&feedback); err != nil {
		return c.JSON(http.StatusInternalServerError, err.Error())
	}

	return c.JSON(http.StatusOK, models.FeedbackToDTO(&feedback))
}

func (fc *FeedbackController) DeleteFeedback(c echo.Context) error {
	var feedbackParams models.Feedback
	err := c.Bind(&feedbackParams)
	if err != nil {
		return c.JSON(http.StatusBadRequest, err.Error())
	}
	id := feedbackParams.ID
	err = fc.feedbackService.DeleteFeedback(uint(id))
	if err != nil {
		return c.JSON(http.StatusNotFound, err.Error())
	}
	return c.JSON(http.StatusOK, "Feedback deleted successfully")
}