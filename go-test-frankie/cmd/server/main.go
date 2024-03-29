package main

import (
	"log"
	"os"

	"go-test-frankie/router"

	"github.com/joho/godotenv"
)

func main() {
	err := godotenv.Load()

	if err != nil {
		log.Fatal("Error loading .env file")
	}

	r := router.Setup()
	r.Run(":" + os.Getenv("APP_PORT"))
}
