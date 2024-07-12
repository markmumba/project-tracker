package main

import (
	"fmt"
	"log"
	"net/http"
	"os"
	"strconv"
	"time"

	"github.com/joho/godotenv"
	"github.com/markmumba/project-tracker/database"
	"github.com/markmumba/project-tracker/models"
	"github.com/markmumba/project-tracker/routes"
)

func main() {

	if err := godotenv.Load(); err != nil {
		log.Fatalf("Error loading .env file: %v", err)
	}
	database.ConnectDB()
	database.DB.AutoMigrate(
		&models.Role{},
		&models.User{},
		&models.Project{},
		&models.Submission{},
		&models.Feedback{},
		&models.CommunicationHistory{},
	)

	// Create roles
	roles := []models.Role{

		{Name: "lecturer"},
		{Name: "student"},
	}
	for _, role := range roles {
		database.DB.Create(&role)
	}

	// Create lecturer users
	lecturers := []models.User{

		{Name: "Paul Mwaniki", Email: "paulmwaniki@gmail.com", Password: "qwerty1234", RoleID: 1},
		{Name: "Florence Kimani", Email: "florencekimani@gmail.com", Password: "qwerty1234", RoleID: 1},

	}
	for _, lecturer := range lecturers {
		database.DB.Create(&lecturer)
	}

	log.Println("Database seeded successfully")

	handler := routes.SetupRouter()
	port, _ := strconv.Atoi(os.Getenv("PORT"))

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
