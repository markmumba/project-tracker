package controllers

import (
	"fmt"
	"net/http"
	"strconv"

	"github.com/labstack/echo/v4"
	"github.com/markmumba/project-tracker/models"
	"github.com/markmumba/project-tracker/services"
)

func CreateProject(c echo.Context) error {
    userId := c.Get("userId").(uint)
    var project models.Project
    if err := c.Bind(&project); err != nil {
        return c.JSON(http.StatusBadRequest, err.Error())
    }
    project.StudentID = userId
    if err := services.CreateProject(&project); err != nil {
        return c.JSON(http.StatusInternalServerError, err.Error())
    }

    return c.JSON(http.StatusCreated, project)
}

func GetProject(c echo.Context) error {
    idStr := c.Param("id")
    id, err := strconv.Atoi(idStr)
    fmt.Println(id)
    if err != nil {
        return c.JSON(http.StatusBadRequest, "Invalid ID")
    }

    project, err := services.GetProject(uint(id))
    if err != nil {
        return c.JSON(http.StatusNotFound, err.Error())
    }

    return c.JSON(http.StatusOK, project)
}
