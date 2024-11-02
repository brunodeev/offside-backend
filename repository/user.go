package repository

import (
	"context"
	"fmt"

	"github.com/brunodeev/offside-backend/model"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

type UserRepository struct {
	collection *mongo.Collection
}

func NewUserRepository(client *mongo.Client, dbName, collName string) *UserRepository {
	return &UserRepository{
		collection: client.Database(dbName).Collection(collName),
	}
}

func (u *UserRepository) Insert(user *model.User) error {
	user.ID = primitive.NewObjectID()

	_, err := u.collection.InsertOne(context.TODO(), user)
	if err != nil {
		return fmt.Errorf("falha na inserção do usuário no banco: %w", err)
	}

	return nil
}
