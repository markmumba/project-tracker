package database

import (
	"fmt"
	"github.com/joho/godotenv"
	"log"
	"os"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var DB *gorm.DB

func ConnectDB() {
    var err error
    err = godotenv.Load()
    if err != nil {
        log.Fatal("Error loading .env file")
    }
    
    dbHost := os.Getenv("PGHOST")
    dbUser := os.Getenv("PGUSER")
    dbPassword := os.Getenv("PGPASSWORD")
    dbName := os.Getenv("PGDATABASE")
    dbPort := os.Getenv("PGPORT")

    dsn := fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%s sslmode=disable TimeZone=Asia/Shanghai",
        dbHost, dbUser, dbPassword, dbName, dbPort)

    DB, err = gorm.Open(postgres.Open(dsn), &gorm.Config{})
    if err != nil {
        log.Fatal("Error connecting to database: ", err)
    }
    log.Println("Connection established successfully")
}
