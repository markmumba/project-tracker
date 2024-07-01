package routes

import (
	"net/http"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"github.com/markmumba/project-tracker/controllers"
	"github.com/markmumba/project-tracker/custommiddleware"
)


func SetupRouter() *echo.Echo {
	e := echo.New()

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

	r := e.Group("")
	r.Use(custommiddleware.Authentication)

	userGroup := r.Group("/users")
	{
		userGroup.GET("/:id", controllers.GetUser)
		userGroup.GET("", controllers.GetAllUsers)
		userGroup.PUT("/:id", controllers.UpdateUser)
		userGroup.DELETE("/:id", controllers.DeleteUser)
	}

	projectGroup := r.Group("/projects")
	{
		projectGroup.POST("", controllers.CreateProject)
		projectGroup.GET("/:id", controllers.GetProject)
		projectGroup.PUT("/:id", controllers.UpdateProject)
		projectGroup.DELETE("/:id", controllers.DeleteProject)
	}

	submissionGroup := r.Group("/submissions")
	{
		submissionGroup.POST("", controllers.CreateSubmission)
		submissionGroup.GET("/:id", controllers.GetSubmission)
		submissionGroup.PUT("/:id", controllers.UpdateSubmission)
		submissionGroup.DELETE("/:id", controllers.DeleteSubmission)
	}

	feedbackGroup := r.Group("/feedbacks")
	{
		feedbackGroup.POST("", controllers.CreateFeedback)
		feedbackGroup.GET("/:id", controllers.GetFeedback)
		feedbackGroup.GET("", controllers.GetAllFeedback)
		feedbackGroup.PUT("/:id", controllers.UpdateFeedback)
		feedbackGroup.DELETE("/:id", controllers.DeleteFeedback)
	}

	communicationGroup := r.Group("/communications")
	{
		communicationGroup.POST("/", controllers.CreateMessage)
		communicationGroup.GET("/:project_id", controllers.GetCommunicationHistory)
	}

	return e	
}
