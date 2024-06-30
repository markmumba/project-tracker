package main

import (
	"fmt"
	"html/template"
	"io"
	"log"
	"net/http"
	"os"
	"strconv"
	"time"

	"github.com/joho/godotenv"
	"github.com/labstack/echo/v4"
	"github.com/markmumba/project-tracker/database"
	"github.com/markmumba/project-tracker/models"
	"github.com/markmumba/project-tracker/routes"
)

type Template struct {
	templates *template.Template
}

func (t *Template) Render(w io.Writer, name string, data interface{}, c echo.Context) error {
	return t.templates.ExecuteTemplate(w, name, data)
}

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
	e := echo.New()

	e.Static("/static", "./UI/static")
	tmpl := template.New("")
	tmpl = template.Must(tmpl.ParseGlob("./UI/Views/partials/*.html"))
	tmpl = template.Must(tmpl.ParseGlob("./UI/Views/student/*.html"))
	tmpl = template.Must(tmpl.ParseGlob("./UI/Views/lecturer/*.html"))
	tmpl = template.Must(tmpl.ParseGlob("./UI/Views/*.html"))

	t := &Template{
		templates: tmpl,
	}
	e.Renderer = t
	handler := routes.SetupRouter(e)
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
		fmt.Println("server failed", err.Error())
	}

}
