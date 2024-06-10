package controllers

import (
	"net/http"

	"github.com/labstack/echo/v4"
	"github.com/markmumba/project-tracker/database"
	"github.com/markmumba/project-tracker/models"
)

func CreateSubmission(c echo.Context) error {
	var submission models.Submission
	if err := c.Bind(&submission); err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{"error": err.Error()})
	}
	if err := database.DB.Create(&submission).Error; err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"error": err.Error()})
	}
	return c.JSON(http.StatusOK, submission)
}

func GetSubmission(c echo.Context) error {
	id := c.Param("id")
	var submission models.Submission
	if err := database.DB.First(&submission, id).Error; err != nil {
		return c.JSON(http.StatusNotFound, map[string]string{"error": "Submission not found"})
	}
	return c.JSON(http.StatusOK, submission)
}
