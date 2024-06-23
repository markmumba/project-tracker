package controllers

import (
	"net/http"
	"strconv"

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

    return c.JSON(http.StatusCreated, feedback)
}

func GetFeedback(c echo.Context) error {
    idStr := c.Param("id")
    id, err := strconv.Atoi(idStr)
    if err != nil {
        return c.JSON(http.StatusBadRequest, "Invalid ID")
    }

    feedback, err := services.GetFeedback(uint(id))
    if err != nil {
        return c.JSON(http.StatusNotFound, err.Error())
    }

    return c.JSON(http.StatusOK, feedback)
}
