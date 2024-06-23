package controllers

import (
	"net/http"
	"strconv"

	"github.com/labstack/echo/v4"
	"github.com/markmumba/project-tracker/models"
	"github.com/markmumba/project-tracker/services"
)

func CreateSubmission(c echo.Context) error {
    userId := c.Get("userId").(uint)
    var submission models.Submission
    if err := c.Bind(&submission); err != nil {
        return c.JSON(http.StatusBadRequest, err.Error())
    }
    submission.StudentID = userId
    if err := services.CreateSubmission(&submission); err != nil {
        return c.JSON(http.StatusInternalServerError, err.Error())
    }

    return c.JSON(http.StatusCreated, submission)
}

func GetSubmission(c echo.Context) error {
    idStr := c.Param("id")
    id, err := strconv.Atoi(idStr)
    if err != nil {
        return c.JSON(http.StatusBadRequest, "Invalid ID")
    }

    submission, err := services.GetSubmission(uint(id))
    if err != nil {
        return c.JSON(http.StatusNotFound, err.Error())
    }

    return c.JSON(http.StatusOK, submission)
}
