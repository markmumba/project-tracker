package routes

import (
	"net/http"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"github.com/markmumba/project-tracker/controllers"
	"github.com/markmumba/project-tracker/custommiddleware"
)

func SetupRouter(e *echo.Echo) *echo.Echo {
	e.Use(middleware.LoggerWithConfig(middleware.LoggerConfig{
		Format: "method=${method}, uri=${uri}, status=${status}\n",
	}))
	e.Use(middleware.Recover())
	e.Use(middleware.CORSWithConfig(middleware.CORSConfig{
		AllowOrigins: []string{"https://labstack.com", "https://labstack.net"},
		AllowMethods: []string{http.MethodGet, http.MethodPut, http.MethodPost, http.MethodDelete},
	}))

	e.POST("/register", controllers.CreateUser)
	e.POST("/login", controllers.Login)
	e.GET("/logout", controllers.Logout)
	e.GET("/", controllers.HomePage)

	// Group routes that require authentication
	r := e.Group("")
	r.Use(custommiddleware.Authentication)

	// User routes
	userGroup := r.Group("/users")
	{
		userGroup.GET("/:id", controllers.GetUser)
		userGroup.DELETE("/:id", controllers.DeleteUser)
	}

	// Project routes
	projectGroup := r.Group("/projects")
	{
		projectGroup.POST("", controllers.CreateProject)
		projectGroup.GET("/:id", controllers.GetProject)
		projectGroup.DELETE("/:id", controllers.DeleteProject)
		projectGroup.GET("/lecturer/:id", controllers.GetAllProjectByLecturerId)
	}

	// Submission routes
	submissionGroup := r.Group("/submissions")
	{
		submissionGroup.POST("", controllers.CreateSubmission)
		submissionGroup.GET("/:id", controllers.GetSubmission)
		submissionGroup.DELETE("/:id", controllers.DeleteSubmission)
		submissionGroup.GET("/student/:id", controllers.GetAllSubmissionByStudentId)
	}

	// Feedback routes
	feedbackGroup := r.Group("/feedbacks")
	{
		feedbackGroup.POST("", controllers.CreateFeedback)
		feedbackGroup.GET("/:id", controllers.GetFeedback)
		feedbackGroup.DELETE("/:id", controllers.DeleteFeedback)
		feedbackGroup.GET("/submission/:id", controllers.GetFeedbackBySubmissionId)
	}

	// Communication routes
	communicationGroup := r.Group("/communications")
	{
		communicationGroup.POST("/", controllers.CreateMessage)
		communicationGroup.GET("/:project_id", controllers.GetCommunicationHistory)
	}

	

	return e

}
