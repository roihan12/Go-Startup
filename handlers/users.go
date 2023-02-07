package handlers

import (
	"bwastartup/helpers"
	"bwastartup/users"
	"net/http"

	"github.com/gin-gonic/gin"
)

type userHandler struct {
	userService users.Service
}

func NewUserHandler(userService users.Service) *userHandler {
	return &userHandler{userService}
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

	formatUser := users.FormatUser(newUser, "tokentokentoken")

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

	formater := users.FormatUser(loginUser, "tokentoekekenff")

	response := helpers.ApiResponse("Succesfuly Login", http.StatusOK, "Success", formater)
	c.JSON(http.StatusOK, response)
}
