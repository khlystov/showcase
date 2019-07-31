package main

import (
	"log"
	"os"

	"go-test-frankie/handlers"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
)

func main() {
	err := godotenv.Load()

	if err != nil {
		log.Fatal("Error loading .env file")
	}

	r := gin.Default()
	r.POST("/isgood", handlers.IsGood)
	r.Run(":" + os.Getenv("APP_PORT"))
}
