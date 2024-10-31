package main

import (
	"context"
	"fmt"
	"os"

	"github.com/brunodeev/offside-backend/database"
	"github.com/brunodeev/offside-backend/model"
	"github.com/gofiber/fiber/v2"
	"github.com/joho/godotenv"
	"go.mongodb.org/mongo-driver/bson"
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

	app.Get("/", func(c *fiber.Ctx) error {
		collection := database.Client.Database("offside-db").Collection("users")

		cur, err := collection.Find(context.TODO(), bson.M{})
		if err != nil {
			return fmt.Errorf("falha na busca dos documentos da collection users: %w", err)
		}

		var users []model.User

		for cur.Next(context.TODO()) {
			var user model.User

			cur.Decode(&user)

			users = append(users, user)
		}

		if users == nil {
			return fmt.Errorf("não há documentos da collection users")
		}

		return c.Status(200).JSON(users)
	})

	app.Listen(":8080")
}
