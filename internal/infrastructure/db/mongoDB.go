package db

import (
	"context"
	"log"
	"os"

	"github.com/joho/godotenv"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func Connection() (*mongo.Client, error) {
	err := godotenv.Load(".env")
	if err != nil {
		log.Fatal("Error loading .env file %w", err)
		return nil, err
	}

	clientOptions := options.Client().ApplyURI("mongodb://localhost:27017")

	username := os.Getenv("MONGO_DB_USERNAME")
	password := os.Getenv("MONGO_DB_PASSWORD")

	clientOptions.SetAuth(options.Credential{
		Username: username,
		Password: password,
	})

	client, err := mongo.Connect(context.Background(), clientOptions)

	if err != nil {
		log.Fatal(err)
		return nil, err
	}

	log.Println("Conected to mongo...")

	return client, nil
}
