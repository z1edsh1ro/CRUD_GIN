package main

import (
	"context"
	"log"
	"main/internal/infrastructure/db"
	router "main/internal/routes"
	"time"
)

func main() {
	mongoClient, err := db.Connection()

	if err != nil {
		log.Panic(err)
	}

	context, cancel := context.WithTimeout(context.Background(), 15*time.Second)
	defer cancel()

	defer func() {
		err = mongoClient.Disconnect(context)

		if err != nil {
			panic(err)
		}
	}()

	r := router.SetupRoutes(mongoClient)
	r.Run()
}
