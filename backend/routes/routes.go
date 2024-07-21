package routes

import (
	"net/http"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"github.com/markmumba/project-tracker/controllers"
	"github.com/markmumba/project-tracker/custommiddleware"
	"github.com/markmumba/project-tracker/services"
)

func SetupRouter(
	userService *services.UserService,
	projectService *services.ProjectService,
	submissionService *services.SubmissionService,
	feedbackService *services.FeedbackService,
) *echo.Echo {
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

	userController := controllers.NewUserController(userService)
	projectController := controllers.NewProjectController(projectService)
	submissionController := controllers.NewSubmissionController(submissionService)
	feedbackController := controllers.NewFeedbackController(feedbackService)

	e.POST("/register", userController.CreateUser)
	e.POST("/login", userController.Login)
	e.GET("/logout", userController.Logout)

	r := e.Group("")
	r.Use(custommiddleware.Authentication)

	userGroup := r.Group("/users")
	{
		userGroup.GET("", userController.GetUser)
		userGroup.GET("/all", userController.GetAllUsers)
		userGroup.GET("/students", userController.GetStudentsByLecturerId)
		userGroup.GET("/lecturers", userController.GetLecturers)
		userGroup.PUT("", userController.UpdateUser)
		userGroup.POST("/profile", userController.UpdateUserProfileImage)
		userGroup.DELETE("", userController.DeleteUser)
	}

	projectGroup := r.Group("/projects")
	{
		projectGroup.POST("", projectController.CreateProject)
		projectGroup.GET("", projectController.GetProject)
		projectGroup.PUT("", projectController.UpdateProject)
		projectGroup.DELETE("", projectController.DeleteProject)
	}

	submissionGroup := r.Group("/submissions")
	{
		submissionGroup.POST("", submissionController.CreateSubmission)
		submissionGroup.GET("/:id", submissionController.GetSubmission)
		submissionGroup.GET("/student", submissionController.GetAllSubmissionByStudentId)
		submissionGroup.GET("/lecturer", submissionController.GetSubmissionsByLecturer)
		submissionGroup.PUT("/:id", submissionController.UpdateSubmission)
		submissionGroup.GET("", submissionController.GetAllSubmissions)
		submissionGroup.DELETE("/:id", submissionController.DeleteSubmission)
	}

	feedbackGroup := r.Group("/feedbacks")
	{
		feedbackGroup.POST("", feedbackController.CreateFeedback)
		feedbackGroup.GET("/student", feedbackController.GetFeedbackByStudent)
		feedbackGroup.GET("", feedbackController.GetAllFeedback)
		feedbackGroup.PUT("/:id", feedbackController.UpdateFeedback)
		feedbackGroup.DELETE("/:id", feedbackController.DeleteFeedback)
	}

	communicationGroup := r.Group("/communications")
	{
		communicationGroup.POST("/", controllers.CreateMessage)
		communicationGroup.GET("/:project_id", controllers.GetCommunicationHistory)
	}

	return e
}
