package controllers

import (
	"net/http"

	"github.com/labstack/echo/v4"
	"github.com/markmumba/project-tracker/database"
	"github.com/markmumba/project-tracker/models"
)

func CreateFeedback(c echo.Context) error {
	var feedback models.Feedback
	if err := c.Bind(&feedback); err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{"error": err.Error()})
	}
	if err := database.DB.Create(&feedback).Error; err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"error": err.Error()})
	}
	return c.JSON(http.StatusOK, feedback)
}

func GetFeedback(c echo.Context) error {
	id := c.Param("id")
	var feedback models.Feedback
	if err := database.DB.First(&feedback, id).Error; err != nil {
		return c.JSON(http.StatusNotFound, map[string]string{"error": "Feedback not found"})
	}
	return c.JSON(http.StatusOK, feedback)
}
