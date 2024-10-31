package utils

import (
	"os"

	"github.com/joho/godotenv"
)

var UserDB, PasswordDB string

func LoadEnv() {
	err := godotenv.Load(".env")
	if err != nil {
		panic("Erro ao buscar variáveis de ambiente")
	}

	UserDB = os.Getenv("DB_USER")
	PasswordDB = os.Getenv("DB_PASSWORD")
}
