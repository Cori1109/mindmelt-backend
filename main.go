package main

import (
	"log"
	"os"

	"github.com/Cori1109/mindmelt-backend/models"
	"github.com/Cori1109/mindmelt-backend/service"
	"github.com/Cori1109/mindmelt-backend/storage"
	"github.com/gofiber/fiber/v2"

	"github.com/joho/godotenv"
)

func main() {
	err := godotenv.Load(".env")

	if err != nil {
		log.Fatal(err)
	}

	config := &storage.Config{
		Host: os.Getenv("DB_HOST"),
		Port: os.Getenv("DB_PORT"),
		User: os.Getenv("DB_USER"),
		Password: os.Getenv("DB_PASS"),
		DBName: os.Getenv("DB_NAME"),
		SSLMode: os.Getenv("DB_SSLMODE"),
	}

	db, err := storage.NewConnection(config)

	if err != nil {
		log.Fatal("could not load database")
	}

	err = models.MigrateUsers(db)

	if err != nil {
		log.Fatal("could not migrate db")
	}

	r := &service.Repository{
		DB: db,
	}

	app := fiber.New()

	r.SetupRoutes(app)

	app.Listen(":8080")
}