package main

import (
	"os"

	"github.com/brunodeev/offside-backend/database"
	"github.com/brunodeev/offside-backend/handler"
	"github.com/gofiber/fiber/v2"
	"github.com/joho/godotenv"
)

func main() {
	app := fiber.New()

	err := godotenv.Load(".env")
	if err != nil {
		panic("Erro ao buscar variáveis de ambiente")
	}

	dbUser := os.Getenv("DB_USER")
	dbPassword := os.Getenv("DB_PASSWORD")

	err = database.ConnectToDB(dbUser, dbPassword)
	if err != nil {
		panic("Deu pau geral!")
	}

	userHandler := handler.NewUserHandler()

	app.Get("/", userHandler.GetUsers)
	app.Post("/register", userHandler.RegisterUser)

	app.Listen(":8080")
}
