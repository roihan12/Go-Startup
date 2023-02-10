package handlers

import (
	"bwastartup/auth"
	"bwastartup/helpers"
	"bwastartup/users"
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
)

type userHandler struct {
	userService users.Service
	authService auth.Service
}

func NewUserHandler(userService users.Service, authService auth.Service) *userHandler {
	return &userHandler{userService, authService}
}

func (uh *userHandler) RegisterUser(c *gin.Context) {
	var input users.RegisterUserInput

	err := c.ShouldBind(&input)
	if err != nil {
		errors := helpers.FormatValidationErrors(err)
		errorMessage := gin.H{"errors": errors}

		response := helpers.ApiResponse("Register account failed", http.StatusUnprocessableEntity, "error", errorMessage)
		c.JSON(http.StatusUnprocessableEntity, response)
		return
	}

	newUser, err := uh.userService.Register(input)
	if err != nil {
		response := helpers.ApiResponse("Register account failed", http.StatusBadRequest, "error", nil)
		c.JSON(http.StatusBadRequest, response)
		return
	}

	token, err := uh.authService.GenerateToken(newUser.ID)
	if err != nil {
		response := helpers.ApiResponse("Register account failed", http.StatusBadRequest, "error", nil)
		c.JSON(http.StatusBadRequest, response)
		return
	}

	formatUser := users.FormatUser(newUser, token)

	response := helpers.ApiResponse("Account has been registered", http.StatusOK, "Success", formatUser)

	c.JSON(http.StatusOK, response)
}

func (uh *userHandler) Login(c *gin.Context) {

	var input users.LoginInput

	err := c.ShouldBind(&input)
	if err != nil {
		errors := helpers.FormatValidationErrors(err)
		errorMessage := gin.H{"errors": errors}
		response := helpers.ApiResponse("Login failed", http.StatusUnprocessableEntity, "error", errorMessage)
		c.JSON(http.StatusUnprocessableEntity, response)
		return
	}

	loginUser, err := uh.userService.Login(input)
	if err != nil {
		errorMessage := gin.H{"errors": err.Error()}
		response := helpers.ApiResponse("Login failed", http.StatusUnprocessableEntity, "error", errorMessage)
		c.JSON(http.StatusUnprocessableEntity, response)
		return
	}

	token, err := uh.authService.GenerateToken(loginUser.ID)
	if err != nil {
		response := helpers.ApiResponse("Login failed", http.StatusBadRequest, "error", nil)
		c.JSON(http.StatusBadRequest, response)
		return
	}

	formater := users.FormatUser(loginUser, token)

	response := helpers.ApiResponse("Succesfuly Login", http.StatusOK, "Success", formater)
	c.JSON(http.StatusOK, response)
}

func (uh *userHandler) CheckEmailAvailability(c *gin.Context) {

	var input users.CheckEmail

	err := c.ShouldBind(&input)
	if err != nil {
		errors := helpers.FormatValidationErrors(err)
		errorMessage := gin.H{"errors": errors}
		response := helpers.ApiResponse("Email cheking failed", http.StatusUnprocessableEntity, "error", errorMessage)
		c.JSON(http.StatusUnprocessableEntity, response)
		return
	}

	isEmailAvailable, err := uh.userService.IsEmailAvailable(input)
	if err != nil {
		errorMessage := gin.H{"errors": "Server error"}
		response := helpers.ApiResponse("Email cheking failed", http.StatusUnprocessableEntity, "error", errorMessage)
		c.JSON(http.StatusUnprocessableEntity, response)
		return
	}

	data := gin.H{
		"is_available": isEmailAvailable,
	}
	metaMessage := "Email has been registered"
	if isEmailAvailable {

		metaMessage = "Email is available"
	}

	response := helpers.ApiResponse(metaMessage, http.StatusOK, "Success", data)
	c.JSON(http.StatusOK, response)
}

func (uh *userHandler) UploadAvatar(c *gin.Context) {
	file, err := c.FormFile("avatar")
	if err != nil {
		data := gin.H{"is_uploaded": false}
		response := helpers.ApiResponse("failed to upload avatar image", http.StatusBadRequest, "error", data)
		c.JSON(http.StatusBadRequest, response)
		return
	}

	currentUser := c.MustGet("currentUser").(users.User)

	userID := currentUser.ID

	path := fmt.Sprintf("images/%d-%s", userID, file.Filename)

	err = c.SaveUploadedFile(file, path)
	if err != nil {
		data := gin.H{"is_uploaded": false}
		response := helpers.ApiResponse("failed to upload avatar image", http.StatusBadRequest, "error", data)
		c.JSON(http.StatusBadRequest, response)
		return
	}

	_, err = uh.userService.SaveAvatar(userID, path)
	if err != nil {
		data := gin.H{"is_uploaded": false}
		response := helpers.ApiResponse("failed to upload avatar image", http.StatusBadRequest, "error", data)
		c.JSON(http.StatusBadRequest, response)
		return
	}
	data := gin.H{"is_uploaded": true}
	response := helpers.ApiResponse("Succesfuly upload avatar image", http.StatusOK, "success", data)
	c.JSON(http.StatusOK, response)

}
