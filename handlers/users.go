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
		response := helpers.ApiResponse("Register account failed", http.StatusBadRequest, "error", nil)
		c.JSON(http.StatusBadRequest, response)
		return
	}

	newUser, err := uh.userService.Register(input)
	if err != nil {
		response := helpers.ApiResponse("Register account failed", http.StatusBadRequest, "error", nil)
		c.JSON(http.StatusBadRequest, response)
		return
	}

	formatUser := users.FormatUser(newUser, "tokentokentoken")

	response := helpers.ApiResponse("Account has been registered", http.StatusCreated, "Success", formatUser)

	c.JSON(http.StatusOK, response)
}
