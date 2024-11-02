package main

import (
	"log"

	"github.com/brunodeev/offside-backend/database"
	"github.com/brunodeev/offside-backend/handler"
	"github.com/brunodeev/offside-backend/utils"
	"github.com/gofiber/fiber/v2"
)

func main() {
	utils.LoadEnv()
	app := fiber.New()

	if err := database.ConnectToDB(utils.UserDB, utils.PasswordDB); err != nil {
		log.Fatal("não foi possível estabelecer conexão com o banco!")
	}

	userHandler := handler.NewUserHandler(database.Client)

	app.Get("/", userHandler.GetUsers)
	app.Post("/register", userHandler.RegisterUser)
	app.Get("/login", userHandler.LoginUser)

	app.Listen(":8080")
}
