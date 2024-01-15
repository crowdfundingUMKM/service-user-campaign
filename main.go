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
	"service-user-campaign/middleware"

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
	notifHandler := handler.NewNotifHandler(userCampaignService, authService)

	// END SETUP

	// RUN SERVICE
	router := gin.Default()

	// setup cors
	corsConfig := config.InitCors()
	router.Use(cors.New(corsConfig))

	// group route
	api := router.Group("api/v1")

	// admin request -> service user admin
	// api.GET("/admin/log_service_toAdmin/:admin_id", middleware.AuthApiAdminMiddleware(authService, userCampaignService), userHandler.GetLogtoAdmin)
	api.GET("/admin/service_status/:admin_id", middleware.AuthApiAdminMiddleware(authService, userCampaignService), userHandler.ServiceHealth)
	api.PUT("/admin/deactive_user/:admin_id", middleware.AuthApiAdminMiddleware(authService, userCampaignService), userHandler.DeactiveUser)
	api.PUT("/admin/active_user/:admin_id", middleware.AuthApiAdminMiddleware(authService, userCampaignService), userHandler.ActiveUser)

	// make endpoint get all user by admin
	api.GET("/admin/get_all_user/:admin_id", middleware.AuthApiAdminMiddleware(authService, userCampaignService), userHandler.GetAllUserData)

	// route give information to user about admin
	api.GET("/campaign/getUserCampaignID/:unix_id", userHandler.GetInfoAdminID)

	// verify token
	api.GET("/verifyTokenCampaign", middleware.AuthMiddleware(authService, userCampaignService), userHandler.VerifyToken)

	// Rounting user
	api.POST("/register_campaign", userHandler.RegisterUser)
	api.POST("/login_campaign", userHandler.Login)
	api.GET("/email_check", userHandler.CheckEmailAvailability)
	api.GET("/phone_check", userHandler.CheckPhoneAvailability)

	api.GET("/get_user", middleware.AuthMiddleware(authService, userCampaignService), userHandler.GetUser)

	api.PUT("/update_profile", middleware.AuthMiddleware(authService, userCampaignService), userHandler.UpdateUser)
	//make update password user by unix_id
	api.PUT("/update_password", middleware.AuthMiddleware(authService, userCampaignService), userHandler.UpdatePassword)
	//make create image profile user by unix_id this for update -> update same
	api.POST("/upload_avatar", middleware.AuthMiddleware(authService, userCampaignService), userHandler.UploadAvatar)

	// make logout user by unix_id
	api.DELETE("/logout_campaign", middleware.AuthMiddleware(authService, userCampaignService), userHandler.LogoutUser)

	// Notif route
	api.POST("/report_to_admin", middleware.AuthMiddleware(authService, userCampaignService), notifHandler.ReportToAdmin)
	api.GET("/admin/get_notif_admin", middleware.AuthApiAdminMiddleware(authService, userCampaignService), notifHandler.GetNotifToAdmin)

	url := fmt.Sprintf("%s:%s", os.Getenv("SERVICE_HOST"), os.Getenv("SERVICE_PORT"))
	router.Run(url)
}
