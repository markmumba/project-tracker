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
		AllowOrigins:     []string{"http://localhost:3000"},
		AllowMethods:     []string{http.MethodGet, http.MethodPut, http.MethodPost, http.MethodDelete},
		AllowCredentials: true,
	}))

	e.POST("/register", controllers.CreateUser)
	e.POST("/login", controllers.Login)
	e.GET("/logout", controllers.Logout)

	r := e.Group("")
	r.Use(custommiddleware.Authentication)

	userGroup := r.Group("/users")
	{
		userGroup.GET("", controllers.GetUser)
		userGroup.GET("/all", controllers.GetAllUsers)
		userGroup.GET("/students", controllers.GetStudentsByLecturerId)
		userGroup.GET("/lecturers", controllers.GetLecturers)
		userGroup.PUT("", controllers.UpdateUser)
		userGroup.POST("/profile", controllers.UpdateUserProfileImage)	
		userGroup.DELETE("", controllers.DeleteUser)
	}

	projectGroup := r.Group("/projects")
	{
		projectGroup.POST("", controllers.CreateProject)
		projectGroup.GET("", controllers.GetProject)
		projectGroup.PUT("", controllers.UpdateProject)
		projectGroup.DELETE("", controllers.DeleteProject)
	}

	submissionGroup := r.Group("/submissions")
	{
		submissionGroup.POST("", controllers.CreateSubmission)
		submissionGroup.GET("/:id", controllers.GetSubmission)
		submissionGroup.GET("/student", controllers.GetAllSubmissionByStudentId)
		submissionGroup.GET("/lecturer", controllers.GetSubmissionsByLecturer)
		submissionGroup.PUT("/:id", controllers.UpdateSubmission)
		submissionGroup.GET("", controllers.GetAllSubmissions)
		submissionGroup.DELETE("", controllers.DeleteSubmission)
	}

	feedbackGroup := r.Group("/feedbacks")
	{
		feedbackGroup.POST("", controllers.CreateFeedback)
		feedbackGroup.GET("/student", controllers.GetFeedbackByStudent)
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
