package configs

import (
	"log"
	"os"

	"github.com/joho/godotenv"
)

func EnvMongoURI() string {
	err := godotenv.Load()

	if err != nil {
		log.Println("No .env file found, using environment variables")
	}

	mongoURI := os.Getenv("MONGO_URI")

	if mongoURI == "" {
		log.Fatal("MONGO_URI is not set")
	}

	return mongoURI
}

func EnvPort() string {
	err := godotenv.Load()

	if err != nil {
		log.Println("No .env file found, using environment variables")
	}

	port := os.Getenv("PORT")

	if port == "" {
		log.Fatal("PORT is not set")
	}

	return port
}
