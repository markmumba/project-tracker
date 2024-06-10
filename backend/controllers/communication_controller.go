package controllers

import (
	"net/http"

	"github.com/labstack/echo/v4"
	"github.com/markmumba/project-tracker/database"
	"github.com/markmumba/project-tracker/models"
)

func CreateMessage(c echo.Context) error {
	var message models.CommunicationHistory
	if err := c.Bind(&message); err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{"error": err.Error()})
	}
	if err := database.DB.Create(&message).Error; err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"error": err.Error()})
	}
	return c.JSON(http.StatusOK, message)
}

func GetCommunicationHistory(c echo.Context) error {
	projectID := c.Param("project_id")
	var history []models.CommunicationHistory
	if err := database.DB.Where("project_id = ?", projectID).Find(&history).Error; err != nil {
		return c.JSON(http.StatusNotFound, map[string]string{"error": "Communication history not found"})
	}
	return c.JSON(http.StatusOK, history)
}
