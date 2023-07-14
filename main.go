package main

import (
	"fmt"
	"log"
	"os"
	"service-user-campaign/auth"
	"service-user-campaign/config"
	"service-user-campaign/core"
	"service-user-campaign/database"
	"service-user-campaign/handler"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
)

func main() {
	fmt.Println("Service User Campaign")
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}
	// setup log
	// config.InitLog()
	// setup repository
	db := database.NewConnectionDB()
	userCampaignRepository := core.NewRepository(db)

	// setup service
	userCampaignService := core.NewService(userCampaignRepository)
	authService := auth.NewService()

	// setup handler
	userHandler := handler.NewUserHandler(userCampaignService, authService)

	// END SETUP

	// RUN SERVICE
	router := gin.Default()

	// setup cors
	corsConfig := config.InitCors()
	router.Use(cors.New(corsConfig))

	// group route
	api := router.Group("api/v1")

	// Rounting admin
	api.POST("/register_campaign", userHandler.RegisterUser)

	url := fmt.Sprintf("%s:%s", os.Getenv("SERVICE_HOST"), os.Getenv("SERVICE_PORT"))
	router.Run(url)
}
