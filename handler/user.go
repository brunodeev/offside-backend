package handler

import (
	"context"
	"fmt"

	"github.com/brunodeev/offside-backend/database"
	"github.com/brunodeev/offside-backend/model"
	"github.com/gofiber/fiber/v2"
	"go.mongodb.org/mongo-driver/bson"
)

type UserHandler struct{}

func NewUserHandler() *UserHandler {
	return &UserHandler{}
}

func (u *UserHandler) GetUsers(c *fiber.Ctx) error {
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
}
