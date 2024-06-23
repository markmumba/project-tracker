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
	}

	projectGroup := r.Group("/projects")
	{
		projectGroup.POST("", controllers.CreateProject)

		// Restrict GET project by ID to lecturers only
		projectGroup.GET("/:id", func(c echo.Context) error {
			userRole := c.Get("userRole").(string)
			if userRole != "lecturer" {
				return c.JSON(http.StatusForbidden, echo.Map{
					"message": "Forbidden",
				})
			}
			return controllers.GetProject(c)
		})
	}

	submissionGroup := r.Group("/submissions")
	{
		submissionGroup.POST("", controllers.CreateSubmission)

		// Restrict GET submission by ID to students only
		submissionGroup.GET("/:id", func(c echo.Context) error {
			userRole := c.Get("userRole").(string)
			if userRole != "student" {
				return c.JSON(http.StatusForbidden, echo.Map{
					"message": "Forbidden",
				})
			}
			return controllers.GetSubmission(c)
		})
	}

	feedbackGroup := r.Group("/feedbacks")
	{
		// Middleware to restrict access to lecturers only
		feedbackGroup.Use(func(next echo.HandlerFunc) echo.HandlerFunc {
			return func(c echo.Context) error {
				userRole := c.Get("userRole").(string)
				if userRole != "lecturer" {
					return c.JSON(http.StatusForbidden, echo.Map{
						"message": "Forbidden",
					})
				}
				return next(c)
			}
		})

		feedbackGroup.POST("", controllers.CreateFeedback)
		feedbackGroup.GET("/:id", controllers.GetFeedback)
	}

	communicationGroup := r.Group("/communications")
	{
		communicationGroup.POST("/", controllers.CreateMessage)
		communicationGroup.GET("/:project_id", controllers.GetCommunicationHistory)
	}

	return e
}
