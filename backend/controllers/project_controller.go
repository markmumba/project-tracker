package controllers

import (
	"github.com/labstack/echo/v4"
	"github.com/markmumba/project-tracker/models"
	"github.com/markmumba/project-tracker/services"
	"net/http"
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

	return c.JSON(http.StatusCreated, models.ProjectToDTO(&project))
}

func GetProject(c echo.Context) error {
	userId := c.Get("userId").(uint)
	project, err := services.GetProject(userId)
	if err != nil {
		return c.JSON(http.StatusNotFound, err.Error())
	}

	return c.JSON(http.StatusOK, models.ProjectToDTO(project))
}

func GetAllProjectByLecturerId(c echo.Context) error {
	userId := c.Get("userId").(uint)
	projects, err := services.GetProjectsByLecturerId(userId)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, err.Error())
	}
	return c.JSON(http.StatusOK, models.ProjectToDTOs(projects))
}

func UpdateProject(c echo.Context) error {
	var project models.Project
	if err := c.Bind(&project); err != nil {
		return c.JSON(http.StatusBadRequest, err.Error())
	}

	if err := services.UpdateProject(&project); err != nil {
		return c.JSON(http.StatusInternalServerError, err.Error())
	}

	return c.JSON(http.StatusOK, models.ProjectToDTO(&project))
}

func DeleteProject(c echo.Context) error {
	var projectParams models.Project
	err := c.Bind(&projectParams)
	if err != nil {
		return c.JSON(http.StatusBadRequest, err.Error())
	}
	id := projectParams.ID
	err = services.DeleteProject(uint(id))
	if err != nil {
		return c.JSON(http.StatusNotFound, err.Error())
	}

	return c.JSON(http.StatusOK, "Project deleted successfully")
}
