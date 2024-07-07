package controllers

import (
	"net/http"
	"strconv"

	"github.com/labstack/echo/v4"
	"github.com/markmumba/project-tracker/helpers"
	"github.com/markmumba/project-tracker/models"
	"github.com/markmumba/project-tracker/services"
)

// TODO : update the update project function to use the id and the new data
// TODO : update the delete project function to use the id
// TODO : update all remaining functions to use the helper function

type CreateProjectRequest struct {
	Title       string `json:"title"`
	LecturerID  string `json:"lecturer_id"`
	Description string `json:"description"`
	StartDate   string `json:"start_date"`
	EndDate     string `json:"end_date"`
}

func CreateProject(c echo.Context) error {
	userID, err := helpers.ConvertUserID(c, "userId")
	if err != nil {
		return c.JSON(http.StatusInternalServerError, err.Error())
	}

	var request CreateProjectRequest
	if err := c.Bind(&request); err != nil {
		return c.JSON(http.StatusBadRequest, err.Error())
	}

	// Convert lecturer_id to uint
	lecturerID, err := strconv.ParseUint(request.LecturerID, 10, 32)
	if err != nil {
		return c.JSON(http.StatusBadRequest, "invalid lecturer_id")
	}

	// Populate the final Project struct
	project := models.Project{
		Title:       request.Title,
		LecturerID:  uint(lecturerID),
		Description: request.Description,
		StartDate:   request.StartDate,
		EndDate:     request.EndDate,
		StudentID:   userID,
	}

	if err := services.CreateProject(&project); err != nil {
		return c.JSON(http.StatusInternalServerError, err.Error())
	}

	return c.JSON(http.StatusCreated, models.ProjectToDTO(&project))
}

func GetProject(c echo.Context) error {

	userID, err := helpers.ConvertUserID(c, "userId")

	if err != nil {
		return c.JSON(http.StatusInternalServerError, err.Error())
	}
	project, err := services.GetProject(userID)
	if err != nil {
		return c.JSON(http.StatusNotFound, err.Error())
	}

	return c.JSON(http.StatusOK, models.ProjectToDTO(project))
}

func GetAllProjectByLecturerId(c echo.Context) error {

	userID, err := helpers.ConvertUserID(c, "userId")
	if err != nil {
		return c.JSON(http.StatusInternalServerError, err.Error())
	}
	projects, err := services.GetProjectsByLecturerId(userID)
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
