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
		&models.User{},
		&models.Project{},
		&models.Submission{},
		&models.Feedback{},
		&models.CommunicationHistory{},
	)
	
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
