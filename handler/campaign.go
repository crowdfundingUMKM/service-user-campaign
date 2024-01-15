package handler

import (
	"fmt"
	"io"
	"net/http"
	"net/url"
	"os"
	api_admin "service-user-campaign/api/admin"
	"service-user-campaign/auth"
	"service-user-campaign/core"
	"service-user-campaign/helper"

	"cloud.google.com/go/storage"
	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
	"google.golang.org/api/option"
	"google.golang.org/appengine"
)

type userCampaignHandler struct {
	userService core.Service
	authService auth.Service
}

func NewUserHandler(userService core.Service, authService auth.Service) *userCampaignHandler {
	return &userCampaignHandler{userService, authService}
}

var (
	storageClient *storage.Client
)

func (h *userCampaignHandler) ServiceHealth(c *gin.Context) {
	// check env open or not
	errEnv := godotenv.Load()
	if errEnv != nil {
		response := helper.APIResponse("Failed to get env for service campaign", http.StatusBadRequest, "error", nil)
		c.JSON(http.StatusBadRequest, response)
		return
	}

	adminID := c.Param("admin_id")
	adminInput := api_admin.AdminIdInput{UnixID: adminID}
	getAdminValueId, err := api_admin.GetAdminId(adminInput)

	if err != nil {
		response := helper.APIResponse(err.Error(), http.StatusBadRequest, "error", nil)
		c.JSON(http.StatusBadRequest, response)
		return
	}

	// middleware api admin
	currentAdmin := c.MustGet("currentUserAdmin").(api_admin.AdminId)

	if c.Param("admin_id") != getAdminValueId && currentAdmin.UnixAdmin != getAdminValueId {
		response := helper.APIResponse("Your not Admin, cannot Access", http.StatusUnprocessableEntity, "error", nil)
		c.JSON(http.StatusNotFound, response)
		return
	}

	envVars := []string{
		"ADMIN_ID",
		"DB_USER",
		"DB_PASS",
		"DB_NAME",
		"DB_PORT",
		"INSTANCE_HOST",
		"SERVICE_HOST",
		"SERVICE_PORT",
		"JWT_SECRET",
		"STATUS_ACCOUNT",
	}

	data := make(map[string]interface{})
	for _, key := range envVars {
		data[key] = os.Getenv(key)
	}

	errService := c.Errors
	if errService != nil {
		response := helper.APIResponse("Service campaign is not running", http.StatusInternalServerError, "error", nil)
		c.JSON(http.StatusInternalServerError, response)
		return
	}
	response := helper.APIResponse("Service campaign is running", http.StatusOK, "success", data)
	c.JSON(http.StatusOK, response)
}

func (h *userCampaignHandler) DeactiveUser(c *gin.Context) {
	var input core.DeactiveUserInput
	// check input from user
	err := c.ShouldBindJSON(&input)
	if err != nil {
		errors := helper.FormatValidationError(err)
		errorMessage := gin.H{"errors": errors}

		response := helper.APIResponse("User Not Found", http.StatusUnprocessableEntity, "error", errorMessage)
		c.JSON(http.StatusUnprocessableEntity, response)
		return
	}

	// cheack id from get param and fetch data from service admin to check id admin and status account admin
	// var adminInput campaign.AdminIdInput
	adminID := c.Param("admin_id")
	adminInput := api_admin.AdminIdInput{UnixID: adminID}
	getAdminValueId, err := api_admin.GetAdminId(adminInput)

	if err != nil {
		response := helper.APIResponse(err.Error(), http.StatusBadRequest, "error", nil)
		c.JSON(http.StatusBadRequest, response)
		return
	}

	// middleware api admin
	currentAdmin := c.MustGet("currentUserAdmin").(api_admin.AdminId)

	// check id admin
	// adminId := getAdminValueId
	if c.Param("admin_id") == getAdminValueId && currentAdmin.UnixAdmin == getAdminValueId {
		// get id user

		// deactive user
		deactive, err := h.userService.DeactivateAccountUser(input, currentAdmin.UnixAdmin)

		data := gin.H{
			"success_deactive": deactive,
		}

		if err != nil {
			response := helper.APIResponse("Failed to deactive user", http.StatusBadRequest, "error", data)
			c.JSON(http.StatusBadRequest, response)
			return
		}
		response := helper.APIResponse("User has been deactive", http.StatusOK, "success", data)
		c.JSON(http.StatusOK, response)
	} else {
		response := helper.APIResponse("Your not Admin, cannot Access", http.StatusUnprocessableEntity, "error", nil)
		c.JSON(http.StatusNotFound, response)
		return
	}
}

func (h *userCampaignHandler) ActiveUser(c *gin.Context) {
	var input core.DeactiveUserInput
	// check input from user
	err := c.ShouldBindJSON(&input)
	if err != nil {
		errors := helper.FormatValidationError(err)
		errorMessage := gin.H{"errors": errors}

		response := helper.APIResponse("User Not Found", http.StatusUnprocessableEntity, "error", errorMessage)
		c.JSON(http.StatusUnprocessableEntity, response)
		return
	}

	// cheack id from get param and fetch data from service admin to check id admin and status account admin
	// var adminInput campaign.AdminIdInput
	adminID := c.Param("admin_id")
	adminInput := api_admin.AdminIdInput{UnixID: adminID}
	getAdminValueId, err := api_admin.GetAdminId(adminInput)

	if err != nil {
		response := helper.APIResponse(err.Error(), http.StatusBadRequest, "error", nil)
		c.JSON(http.StatusBadRequest, response)
		return
	}

	// middleware api admin
	currentAdmin := c.MustGet("currentUserAdmin").(api_admin.AdminId)

	// check id admin
	// adminId := getAdminValueId
	if c.Param("admin_id") == getAdminValueId && currentAdmin.UnixAdmin == getAdminValueId {
		// get id user

		// deactive user
		active, err := h.userService.ActivateAccountUser(input, currentAdmin.UnixAdmin)

		data := gin.H{
			"success_active": active,
		}

		if err != nil {
			response := helper.APIResponse("Failed to active user", http.StatusBadRequest, "error", data)
			c.JSON(http.StatusBadRequest, response)
			return
		}
		response := helper.APIResponse("User has been active", http.StatusOK, "success", data)
		c.JSON(http.StatusOK, response)
	} else {
		response := helper.APIResponse("Your not Admin, cannot Access", http.StatusUnprocessableEntity, "error", nil)
		c.JSON(http.StatusNotFound, response)
		return
	}
}

// get info id admin not use middleware
func (h *userCampaignHandler) GetInfoAdminID(c *gin.Context) {
	var inputID core.GetUserIdInput

	// check id is valid or not
	err := c.ShouldBindUri(&inputID)
	if err != nil {
		response := helper.APIResponse("Failed get user admin and status", http.StatusBadRequest, "error", nil)
		c.JSON(http.StatusBadRequest, response)
		return
	}

	user, err := h.userService.GetUserByUnixID(inputID.UnixID)
	if err != nil {
		response := helper.APIResponse("Get user failed", http.StatusBadRequest, "error", nil)
		c.JSON(http.StatusBadRequest, response)
		return
	}

	formatter := core.FormatterUserCampaignID(user)

	response := helper.APIResponse("Successfuly get user id and status", http.StatusOK, "success", formatter)
	c.JSON(http.StatusOK, response)

}

func (h *userCampaignHandler) GetAllUserData(c *gin.Context) {
	// cheack id from get param and fetch data from service admin to check id admin and status account admin
	// var adminInput campaign.AdminIdInput
	adminID := c.Param("admin_id")
	adminInput := api_admin.AdminIdInput{UnixID: adminID}
	getAdminValueId, err := api_admin.GetAdminId(adminInput)

	currentAdmin := c.MustGet("currentUserAdmin").(api_admin.AdminId)

	if err != nil {
		response := helper.APIResponse(err.Error(), http.StatusBadRequest, "error", nil)
		c.JSON(http.StatusBadRequest, response)
		return
	}

	if c.Param("admin_id") == getAdminValueId && currentAdmin.UnixAdmin == getAdminValueId {
		users, err := h.userService.GetAllUsers()
		if err != nil {
			response := helper.APIResponse("Failed to get user", http.StatusBadRequest, "error", nil)
			c.JSON(http.StatusBadRequest, response)
			return
		}
		response := helper.APIResponse("List of user campaign", http.StatusOK, "success", users)
		c.JSON(http.StatusOK, response)
	} else {
		response := helper.APIResponse("Your not Admin, cannot Access", http.StatusUnprocessableEntity, "error", nil)
		c.JSON(http.StatusNotFound, response)
		return
	}
}

// verify token
// cen acces to veriefy token VerifyToken
func (h *userCampaignHandler) VerifyToken(c *gin.Context) {
	currentUser := c.MustGet("currentUser").(core.User)

	// check f account deactive
	if currentUser.StatusAccount == "deactive" {
		errorMessage := gin.H{"errors": "Your account is deactive by admin"}
		response := helper.APIResponse("Your token failed", http.StatusUnprocessableEntity, "error", errorMessage)
		c.JSON(http.StatusUnprocessableEntity, response)
		return
	}
	// if you logout you can't get user
	if currentUser.Token == "" {
		errorMessage := gin.H{"errors": "Your account is logout"}
		response := helper.APIResponse("Your token failed", http.StatusUnprocessableEntity, "error", errorMessage)
		c.JSON(http.StatusUnprocessableEntity, response)
		return
	}

	data := gin.H{
		"success":  "Your token is valid",
		"admin_id": currentUser.UnixID,
	}

	response := helper.APIResponse("Successfuly get user by middleware", http.StatusOK, "success", data)
	c.JSON(http.StatusOK, response)
}

func (h *userCampaignHandler) RegisterUser(c *gin.Context) {
	// tangkap input dari user
	// map input dari user ke struct RegisterUserInput
	// struct di atas kita passing sebagai parameter service

	var input core.RegisterUserInput

	err := c.ShouldBindJSON(&input)
	if err != nil {
		errors := helper.FormatValidationError(err)
		errorMessage := gin.H{"errors": errors}

		response := helper.APIResponse("Register account failed", http.StatusUnprocessableEntity, "error", errorMessage)
		c.JSON(http.StatusUnprocessableEntity, response)
		return
	}

	newUser, err := h.userService.RegisterUser(input)
	if err != nil {
		response := helper.APIResponse("Register account failed", http.StatusBadRequest, "error", nil)
		c.JSON(http.StatusBadRequest, response)
		return
	}
	// generate token
	token, err := h.authService.GenerateToken(newUser.UnixID)
	if err != nil {
		if err != nil {
			response := helper.APIResponse("Register account failed", http.StatusBadRequest, "error", nil)
			c.JSON(http.StatusBadRequest, response)
			return
		}
	}

	formatter := core.FormatterUser(newUser, token)

	if formatter.StatusAccount == "active" {
		_, err = h.userService.SaveToken(newUser.UnixID, token)

		if err != nil {
			response := helper.APIResponse("Register account failed", http.StatusBadRequest, "error", nil)
			c.JSON(http.StatusBadRequest, response)
			return
		}

		response := helper.APIResponse("Account has been registered and active", http.StatusOK, "success", formatter)
		c.JSON(http.StatusOK, response)
		return
	}

	data := gin.H{
		"status": "Account has been registered, but you must wait admin to active your account",
	}

	response := helper.APIResponse("Account has been registered but you must wait admin or review to active your account", http.StatusOK, "success", data)

	c.JSON(http.StatusOK, response)
}

// Login User Admin
func (h *userCampaignHandler) Login(c *gin.Context) {

	var input core.LoginInput

	err := c.ShouldBindJSON(&input)
	if err != nil {
		errors := helper.FormatValidationError(err)
		errorMessage := gin.H{"errors": errors}

		response := helper.APIResponse("Login failed", http.StatusUnprocessableEntity, "error", errorMessage)
		c.JSON(http.StatusUnprocessableEntity, response)
		return
	}

	loggedinUser, err := h.userService.Login(input)
	if err != nil {
		errorMessage := gin.H{"errors": err.Error()}
		response := helper.APIResponse("Login failed", http.StatusUnprocessableEntity, "error", errorMessage)
		c.JSON(http.StatusUnprocessableEntity, response)
		return
	}
	// generate token
	token, err := h.authService.GenerateToken(loggedinUser.UnixID)

	// save toke to database
	_, err = h.userService.SaveToken(loggedinUser.UnixID, token)

	if err != nil {
		response := helper.APIResponse("Login failed", http.StatusBadRequest, "error", nil)
		c.JSON(http.StatusBadRequest, response)
		return
	}
	// end save token to database

	if err != nil {
		if err != nil {
			response := helper.APIResponse("Login failed", http.StatusBadRequest, "error", nil)
			c.JSON(http.StatusBadRequest, response)
			return
		}
	}

	// check role acvtive and not send massage your account deactive
	if loggedinUser.StatusAccount == "deactive" {
		errorMessage := gin.H{"errors": "Your account is deactive by admin"}
		response := helper.APIResponse("Login failed", http.StatusUnprocessableEntity, "error", errorMessage)
		c.JSON(http.StatusUnprocessableEntity, response)
		return
	}
	formatter := core.FormatterUser(loggedinUser, token)

	response := helper.APIResponse("Succesfuly loggedin", http.StatusOK, "success", formatter)

	c.JSON(http.StatusOK, response)
}

func (h *userCampaignHandler) GetUser(c *gin.Context) {
	currentUser := c.MustGet("currentUser").(core.User)

	// check f account deactive
	if currentUser.StatusAccount == "deactive" {
		errorMessage := gin.H{"errors": "Your account is deactive by admin"}
		response := helper.APIResponse("Get user failed", http.StatusUnprocessableEntity, "error", errorMessage)
		c.JSON(http.StatusUnprocessableEntity, response)
		return
	}
	// if you logout you can't get user
	if currentUser.Token == "" {
		errorMessage := gin.H{"errors": "Your account is logout"}
		response := helper.APIResponse("Get user failed", http.StatusUnprocessableEntity, "error", errorMessage)
		c.JSON(http.StatusUnprocessableEntity, response)
		return
	}

	formatter := core.FormatterUserDetail(currentUser, currentUser)

	response := helper.APIResponse("Successfuly get user by middleware", http.StatusOK, "success", formatter)
	c.JSON(http.StatusOK, response)
}

func (h *userCampaignHandler) UpdateUser(c *gin.Context) {
	currentUser := c.MustGet("currentUser").(core.User)

	// if account deactive
	if currentUser.StatusAccount == "deactive" {
		errorMessage := gin.H{"errors": "Your account is deactive by admin"}
		response := helper.APIResponse("Update user failed", http.StatusUnprocessableEntity, "error", errorMessage)
		c.JSON(http.StatusUnprocessableEntity, response)
		return
	}

	// if you logout you can't get user
	if currentUser.Token == "" {
		errorMessage := gin.H{"errors": "Your account is logout"}
		response := helper.APIResponse("Get user failed", http.StatusUnprocessableEntity, "error", errorMessage)
		c.JSON(http.StatusUnprocessableEntity, response)
		return
	}

	var inputData core.UpdateUserInput

	err := c.ShouldBindJSON(&inputData)
	if err != nil {
		errors := helper.FormatValidationError(err)
		errorMessage := gin.H{"errors": errors}

		response := helper.APIResponse("Update user failed, input data failure", http.StatusUnprocessableEntity, "error", errorMessage)
		c.JSON(http.StatusUnprocessableEntity, response)
		return
	}

	updatedUser, err := h.userService.UpdateUserByUnixID(currentUser.UnixID, inputData)
	if err != nil {
		response := helper.APIResponse("Update user failed", http.StatusBadRequest, "error", nil)
		c.JSON(http.StatusBadRequest, response)
		return
	}

	formatter := core.FormatterUserDetail(currentUser, updatedUser)

	response := helper.APIResponse("User has been updated", http.StatusOK, "success", formatter)
	c.JSON(http.StatusOK, response)
	return
}

func (h *userCampaignHandler) UpdatePassword(c *gin.Context) {
	currentUser := c.MustGet("currentUser").(core.User)

	// if you logout you can't get user
	if currentUser.Token == "" {
		errorMessage := gin.H{"errors": "Your account is logout"}
		response := helper.APIResponse("Get user failed", http.StatusUnprocessableEntity, "error", errorMessage)
		c.JSON(http.StatusUnprocessableEntity, response)
		return
	}

	var inputData core.UpdatePasswordInput

	err := c.ShouldBindJSON(&inputData)
	if err != nil {
		errors := helper.FormatValidationError(err)
		errorMessage := gin.H{"errors": errors}

		response := helper.APIResponse("Update password failed, input data failure", http.StatusUnprocessableEntity, "error", errorMessage)
		c.JSON(http.StatusUnprocessableEntity, response)
		return
	}

	updatedUser, err := h.userService.UpdatePasswordByUnixID(currentUser.UnixID, inputData)
	if err != nil {
		response := helper.APIResponse("Update password failed", http.StatusBadRequest, "error", nil)
		c.JSON(http.StatusBadRequest, response)
		return
	}

	formatter := core.FormatterUserDetail(currentUser, updatedUser)

	response := helper.APIResponse("Password has been updated", http.StatusOK, "success", formatter)
	c.JSON(http.StatusOK, response)
	return
}

func (h *userCampaignHandler) UploadAvatar(c *gin.Context) {
	f, _, err := c.Request.FormFile("avatar_file_name")
	if err != nil {
		data := gin.H{"is_uploaded": false}
		response := helper.APIResponse("Failed to upload avatar image", http.StatusBadRequest, "error", data)

		c.JSON(http.StatusBadRequest, response)
		return
	}

	currentUser := c.MustGet("currentUser").(core.User)
	userID := currentUser.UnixID
	userName := currentUser.Name

	// if you logout you can't get user
	if currentUser.Token == "" {
		errorMessage := gin.H{"errors": "Your account is logout"}
		response := helper.APIResponse("Get user failed", http.StatusUnprocessableEntity, "error", errorMessage)
		c.JSON(http.StatusUnprocessableEntity, response)
		return
	}

	// initiate cloud storage os.Getenv("GCS_BUCKET")
	bucket := fmt.Sprintf("%s", os.Getenv("GCS_BUCKET"))
	subfolder := fmt.Sprintf("%s", os.Getenv("GCS_SUBFOLDER"))
	// var err error
	ctx := appengine.NewContext(c.Request)

	storageClient, err = storage.NewClient(ctx, option.WithCredentialsFile("secret-keys.json"))

	if err != nil {
		// data := gin.H{"is_uploaded": false}
		response := helper.APIResponse("Failed to upload avatar image to GCP", http.StatusBadRequest, "error", err)

		c.JSON(http.StatusBadRequest, response)
		return
	}
	defer f.Close()

	objectName := fmt.Sprintf("%s/avatar-%s-%s", subfolder, userID, userName)
	sw := storageClient.Bucket(bucket).Object(objectName).NewWriter(ctx)

	if _, err := io.Copy(sw, f); err != nil {
		// data := gin.H{"is_uploaded": false}
		response := helper.APIResponse("Failed to upload avatar image to GCP", http.StatusBadRequest, "error", err)

		c.JSON(http.StatusBadRequest, response)
		return
	}

	if err := sw.Close(); err != nil {
		// data := gin.H{"is_uploaded": false}
		response := helper.APIResponse("Failed to upload avatar image to GCP", http.StatusBadRequest, "error", err)

		c.JSON(http.StatusBadRequest, response)
		return
	}

	u, err := url.Parse("/" + bucket + "/" + sw.Attrs().Name)
	if err != nil {
		// data := gin.H{"is_uploaded": false}
		response := helper.APIResponse("Failed to upload avatar image to GCP", http.StatusBadRequest, "error", err)

		c.JSON(http.StatusBadRequest, response)
		return
	}
	path := u.String()

	// save avatar to database
	_, err = h.userService.SaveAvatar(userID, path)
	if err != nil {
		data := gin.H{"is_uploaded": false}
		response := helper.APIResponse("Failed to upload avatar image", http.StatusBadRequest, "error", data)

		c.JSON(http.StatusBadRequest, response)
		return
	}

	data := gin.H{"is_uploaded": true}
	response := helper.APIResponse("Avatar successfuly uploaded", http.StatusOK, "success", data)

	c.JSON(http.StatusOK, response)
}

func (h *userCampaignHandler) LogoutUser(c *gin.Context) {
	// get data from middleware
	currentUser := c.MustGet("currentUser").(core.User)

	// check if token is empty
	if currentUser.Token == "" {
		response := helper.APIResponse("Logout failed, your logout right now", http.StatusForbidden, "error", nil)
		c.JSON(http.StatusBadRequest, response)
		return
	}

	// delete token in database
	_, err := h.userService.DeleteToken(currentUser.UnixID)
	if err != nil {
		response := helper.APIResponse("Logout failed", http.StatusBadRequest, "error", nil)
		c.JSON(http.StatusBadRequest, response)
		return
	}

	response := helper.APIResponse("Logout success", http.StatusOK, "success", nil)
	c.JSON(http.StatusOK, response)
	return
}
