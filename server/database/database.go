package database

import (
	"fmt"
	"log"
	"os"

	"fst/project/models"

	"github.com/joho/godotenv"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

type DbInstance struct {
	Db *gorm.DB
}

var Database DbInstance

func ConnectDb() {

	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}
	// Generate encoded token and send it as response.
	dsn := os.Getenv("DATABASE_URL")
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})

	if err != nil {
		log.Fatal("Failed to connect to the database! \n", err)
		os.Exit(2)
	}

	log.Println("Connected Successfully to Database")

	db.Logger = logger.Default.LogMode(logger.Info)

	log.Println("Running Migrations")

	res := db.AutoMigrate(&models.User{}, &models.Token{}, &models.Heartbeat{}, &models.Step{}, &models.Sleep{})
	fmt.Println(res)
	Database = DbInstance{
		Db: db,
	}
}
