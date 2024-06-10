package main

import (
	"fmt"
	"net/http"
	"os"
	"strconv"
	"time"

	"github.com/markmumba/project-tracker/database"
	"github.com/markmumba/project-tracker/models"
	"github.com/markmumba/project-tracker/routes"
)

func main() {
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
	fmt.Println("server is up and running")

	srv := &http.Server{
		Addr:        fmt.Sprintf(":%d", port),
		Handler:     handler,
		ReadTimeout: time.Second * 10,
	}
	fmt.Printf("server started on port : %v", port)
	err := srv.ListenAndServe()
	if err != nil {
		fmt.Println("server failed")
	}

}
