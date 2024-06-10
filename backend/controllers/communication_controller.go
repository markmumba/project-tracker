package controllers

import (
	"net/http"
	"strconv"

	"github.com/labstack/echo/v4"
	"github.com/markmumba/project-tracker/models"
	"github.com/markmumba/project-tracker/services"
)

func CreateMessage(c echo.Context) error {
    var message models.CommunicationHistory
    if err := c.Bind(&message); err != nil {
        return c.JSON(http.StatusBadRequest, err.Error())
    }

    if err := services.CreateMessage(&message); err != nil {
        return c.JSON(http.StatusInternalServerError, err.Error())
    }

    return c.JSON(http.StatusCreated, message)
}

func GetCommunicationHistory(c echo.Context) error {
    projectIDStr := c.Param("project_id")
    projectID, err := strconv.Atoi(projectIDStr)
    if err != nil {
        return c.JSON(http.StatusBadRequest, "Invalid project ID")
    }

    messages, err := services.GetCommunicationHistory(uint(projectID))
    if err != nil {
        return c.JSON(http.StatusNotFound, err.Error())
    }

    return c.JSON(http.StatusOK, messages)
}
