package controllers

import (
	"net/http"

	"github.com/labstack/echo/v4"
	"github.com/markmumba/project-tracker/models"
	"github.com/markmumba/project-tracker/services"
)

func CreateProject(c echo.Context) error {
	userId := c.Get("userId").(uint)
	var project models.Project
	if err := c.Bind(&project); err != nil {
		return c.Render(http.StatusBadRequest, "error.html", map[string]interface{}{
			"error": err.Error(),
		})
	}
	project.StudentID = userId
	if err := services.CreateProject(&project); err != nil {
		return c.Render(http.StatusInternalServerError, "error.html", map[string]interface{}{
			"error": err.Error(),
		})
	}

	return c.Redirect(http.StatusFound, "/projects")
}

func GetProject(c echo.Context) error {
	var projectParams models.Project
	err := c.Bind(&projectParams)
	if err != nil {
		return c.Render(http.StatusBadRequest, "error.html", map[string]interface{}{
			"error": err.Error(),
		})
	}
	id := projectParams.ID
	project, err := services.GetProject(id)
	if err != nil {
		return c.Render(http.StatusNotFound, "error.html", map[string]interface{}{
			"error": err.Error(),
		})
	}

	return c.Render(http.StatusOK, "project.html", models.ProjectToDTO(project))
}

func GetAllProjectByLecturerId(c echo.Context) error {
	userId := c.Get("userId").(uint)
	projects, err := services.GetProjectsByLecturerId(userId)
	if err != nil {
		return c.Render(http.StatusInternalServerError, "error.html", map[string]interface{}{
			"error": err.Error(),
		})
	}
	return c.Render(http.StatusOK, "projects.html", models.ProjectToDTOs(projects))
}

func DeleteProject(c echo.Context) error {
	var projectParams models.Project
	err := c.Bind(&projectParams)
	if err != nil {
		return c.Render(http.StatusBadRequest, "error.html", map[string]interface{}{
			"error": err.Error(),
		})
	}
	id := projectParams.ID
	err = services.DeleteProject(id)
	if err != nil {
		return c.Render(http.StatusNotFound, "error.html", map[string]interface{}{
			"error": err.Error(),
		})
	}

	return c.Render(http.StatusOK, "delete_success.html", map[string]interface{}{
		"message": "Project deleted successfully",
	})
}
