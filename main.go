package main

import (
	"fmt"
	"log"

	"github.com/joho/godotenv"
)

func main() {
	fmt.Println("Service User Campaign")
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	// RUN SERVICE
	// router := gin.Default()

	// api := router.Group("api/v1")

}
