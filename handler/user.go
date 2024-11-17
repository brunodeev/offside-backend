package handler

import (
	"context"
	"fmt"

	"github.com/brunodeev/offside-backend/database"
	"github.com/brunodeev/offside-backend/model"
	"github.com/brunodeev/offside-backend/repository"
	"github.com/brunodeev/offside-backend/utils"
	"github.com/gofiber/fiber/v2"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

type UserHandler struct {
	userRepo *repository.UserRepository
}

func NewUserHandler(client *mongo.Client) *UserHandler {
	return &UserHandler{
		userRepo: repository.NewUserRepository(client, "offside-db", "users"),
	}
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

func (u *UserHandler) RegisterUser(c *fiber.Ctx) error {
	var user model.User

	if err := c.BodyParser(&user); err != nil {
		return fmt.Errorf("falha na conversão do usuário")
	}

	hashedPassword, err := utils.HashPassword(user.Password)
	if err != nil {
		return fmt.Errorf("falha na criação do hash da senha")
	}

	user.Password = hashedPassword

	if err := u.userRepo.Insert(&user); err != nil {
		return fmt.Errorf("falha na inserção do usuário")
	}

	if err := c.Status(201).JSON(fiber.Map{
		"message": "Usuário criado com sucesso",
	}); err != nil {
		return fmt.Errorf("falha no envio da resposta JSON: %w", err)
	}

	return nil
}

func (u *UserHandler) LoginUser(c *fiber.Ctx) error {
	var user model.User
	var userMongo model.User

	collection := database.Client.Database("offside-db").Collection("users")

	err := c.BodyParser(&user)
	if err != nil {
		return fmt.Errorf("falha na conversão do usuário")
	}

	result := collection.FindOne(context.TODO(), bson.M{"email": user.Email})
	if result == nil {
		return fmt.Errorf("falha ao encontrar o usuário")
	}

	err = result.Decode(&userMongo)
	if err != nil {
		return fmt.Errorf("falha na conversão do usuário do mongo")
	}

	if user.Email != userMongo.Email || !utils.CheckPassword(userMongo.Password, user.Password) {
		return c.Status(400).JSON(fiber.Map{
			"message": fmt.Sprint("Ih, deu ruim!"),
		})
	}

	return c.Status(200).JSON(fiber.Map{
		"message": fmt.Sprintf("Olá, %s!", userMongo.Name),
	})
}
