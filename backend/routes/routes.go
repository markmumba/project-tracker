package routes

import (
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"github.com/markmumba/project-tracker/controllers"
)

func SetupRouter() *echo.Echo {
	e := echo.New()

	e.Use(middleware.Logger())
	e.Use(middleware.Recover())

	userGroup := e.Group("/users")
	{
		userGroup.POST("/", controllers.CreateUser)
		userGroup.GET("/:id", controllers.GetUser)
	}

	projectGroup := e.Group("/projects")
	{
		projectGroup.POST("/", controllers.CreateProject)
		projectGroup.GET("/:id", controllers.GetProject)
	}

	submissionGroup := e.Group("/submissions")
	{
		submissionGroup.POST("/", controllers.CreateSubmission)
		submissionGroup.GET("/:id", controllers.GetSubmission)
	}

	feedbackGroup := e.Group("/feedbacks")
	{
		feedbackGroup.POST("/", controllers.CreateFeedback)
		feedbackGroup.GET("/:id", controllers.GetFeedback)
	}

	communicationGroup := e.Group("/communications")
	{
		communicationGroup.POST("/", controllers.CreateMessage)
		communicationGroup.GET("/:project_id", controllers.GetCommunicationHistory)
	}

	return e
}
