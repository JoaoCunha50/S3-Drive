package config

import (
	"fmt"
	"os"

	"github.com/joho/godotenv"
)

type env struct {
	PORT string
}

func LoadConfig() *env {
	if err := godotenv.Load(".env"); err != nil {
		fmt.Println("Warning: .env file not found")
	}

	port := os.Getenv("PORT")
	if port == "" {
		port = "3000"
	}
	return &env{PORT: port}
}
