package main

import (
	"github.com/brunodeev/offside-backend/database"
	"github.com/brunodeev/offside-backend/handler"
	"github.com/brunodeev/offside-backend/utils"
	"github.com/gofiber/fiber/v2"
)

func main() {
	utils.LoadEnv()
	app := fiber.New()

	err := database.ConnectToDB(utils.UserDB, utils.PasswordDB)
	if err != nil {
		panic("Deu pau geral!")
	}

	userHandler := handler.NewUserHandler()

	app.Get("/", userHandler.GetUsers)
	app.Post("/register", userHandler.RegisterUser)
	app.Get("/login", userHandler.LoginUser)

	app.Listen(":8080")
}
