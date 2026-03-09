package config

import (
	"fmt"
	"log"
	"os"

	"github.com/joho/godotenv"
)

type Config struct {
	Port          string
	MongoURI      string
	MongoDatabase string
	AdminPassword string
	AdminUsername string
	AdminEmail    string
	JWTSecret     []byte
}

var Env *Config

func init() {
	Env = LoadConfig()
}

func LoadConfig() *Config {
	if err := godotenv.Load(".env"); err != nil {
		fmt.Println("Warning: .env file not found")
	}

	port := os.Getenv("PORT")
	if port == "" {
		port = "3000"
	}

	mongoURI := os.Getenv("MONGO_URI")
	if mongoURI == "" {
		mongoURI = "mongodb://root:root@mongodb:27017"
	}

	mongoDatabase := os.Getenv("MONGO_DATABASE")
	if mongoDatabase == "" {
		mongoDatabase = "s3-drive"
	}

	adminPassword := os.Getenv("ADMIN_PASSWORD")
	if adminPassword == "" {
		log.Fatal("Error: ADMIN_PASSWORD not found")
	}

	adminUsername := os.Getenv("ADMIN_USERNAME")
	if adminUsername == "" {
		log.Fatal("Error: ADMIN_USERNAME not found")
	}

	adminEmail := os.Getenv("ADMIN_EMAIL")
	if adminEmail == "" {
		log.Fatal("Error: ADMIN_EMAIL not found")
	}

	jwtSecret := os.Getenv("JWT_SECRET")
	if jwtSecret == "" {
		log.Fatal("Error: JWT_SECRET not found")
	}

	return &Config{
		Port:          port,
		MongoURI:      mongoURI,
		MongoDatabase: mongoDatabase,
		AdminPassword: adminPassword,
		AdminUsername: adminUsername,
		AdminEmail:    adminEmail,
		JWTSecret:     []byte(jwtSecret),
	}
}
