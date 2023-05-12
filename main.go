package main

import (
	"fmt"
	"log"
	"os"
	"service-user-campaign/auth"
	"service-user-campaign/campaign"
	"service-user-campaign/database"
	"service-user-campaign/handler"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
)

func main() {
	fmt.Println("Service User Campaign")

	// Initiate service
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}
	// setup log
	// L.InitLog()
	// setup repository
	db := database.NewConnectionDB()
	userCampaignRepository := campaign.NewRepository(db)

	// setup service
	userCampaignService := campaign.NewService(userCampaignRepository)
	authService := auth.NewService()

	// setup handler
	userHandler := handler.NewUserHandler(userCampaignService, authService)

	// END SETUP

	// RUN SERVICE
	router := gin.Default()
	api := router.Group("api/v1")

	// Rounting admin
	api.POST("register_user_campaign", userHandler.RegisterUser)

	url := fmt.Sprintf("%s:%s", os.Getenv("SERVICE_HOST"), os.Getenv("SERVICE_PORT"))
	router.Run(url)
}
