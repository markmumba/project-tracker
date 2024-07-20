package main

import (
	"fmt"
	"net/http"
	"os"
	"strconv"
	"time"

	"github.com/markmumba/project-tracker/database"
	"github.com/markmumba/project-tracker/models"
	"github.com/markmumba/project-tracker/repository"
	"github.com/markmumba/project-tracker/routes"
	"github.com/markmumba/project-tracker/services"
)

func main() {

	database.ConnectDB()

	database.DB.AutoMigrate(
		&models.Role{},
		&models.User{},
		&models.Project{},
		&models.Submission{},
		&models.Feedback{},
		&models.CommunicationHistory{},
	)

	userRepository := repository.NewUserRepository()
	userService := services.NewUserService(userRepository)

	projectRepository := repository.NewProjectRepository()
	projectService := services.NewProjectService(projectRepository, userRepository)

	submissionRepository := repository.NewSubmissionRepository()
	submissionService := services.NewSubmissionService(submissionRepository, userRepository)

	feedbackRepository := repository.NewFeedbackRepository()
	feedbackService := services.NewFeedbackService(feedbackRepository)

	handler := routes.SetupRouter(userService, projectService, submissionService, feedbackService)
	port, _ := strconv.Atoi(os.Getenv("BACKEND_PORT"))

	srv := &http.Server{
		Addr:        fmt.Sprintf(":%d", port),
		Handler:     handler,
		ReadTimeout: time.Second * 10,
	}
	fmt.Printf("server started on port : %v", port)
	fmt.Println()
	err := srv.ListenAndServe()
	if err != nil {
		fmt.Println("server failed")
	}

}
