package handler

import (
	"net/http"

	"crud-with-golang/helper"
	"crud-with-golang/user"

	"github.com/gin-gonic/gin"
)

type userHander struct {
	userService user.Service
}

func NewHandler(userService user.Service) *userHander {
	return &userHander{userService}
}

func (h *userHander) CreateUser(c *gin.Context) {
	var inputUser user.CreateUserInput

	err := c.ShouldBindJSON(&inputUser)
	if err != nil {
		errors := helper.FormatValidationError(err)
		errorMessage := gin.H{"errors": errors}
		response := helper.MessageResponse("Register account failed", http.StatusUnprocessableEntity, "error", errorMessage)
		c.JSON(http.StatusBadRequest, response)
		return
	}

	email := user.CheckEmailInput{
		Email: inputUser.Email,
	}

	emailIsAvailable, _ := h.userService.IsEmailAvailable(email)

	if !emailIsAvailable {
		response := helper.MessageResponse("Registration Failed", http.StatusBadRequest, "failed", "Email has already exist")
		c.JSON(http.StatusBadRequest, response)
		return
	}

	newUser, err := h.userService.CreateUser(inputUser)
	if err != nil {
		errorMessage := gin.H{"errors": err}
		response := helper.MessageResponse("Registration Failed", http.StatusBadRequest, "failed", errorMessage)
		c.JSON(http.StatusBadRequest, response)
		return
	}

	formater := user.CreateUserFormat(newUser)
	response := helper.MessageResponse("Account has been registered", http.StatusOK, "success", formater)

	c.JSON(http.StatusOK, response)
}

func (h *userHander) GetUsers(c *gin.Context) {
	result, err := h.userService.GetUsers()
	if err != nil {
		errorMessage := gin.H{"errors": err}
		response := helper.MessageResponse("Get User Failed", http.StatusBadRequest, "failed", errorMessage)
		c.JSON(http.StatusBadRequest, response)
		return
	}

	response := helper.MessageResponse("Get Users Success", http.StatusOK, "success", result)
	c.JSON(http.StatusOK, response)
}

func (h *userHander) GetUserByID(c *gin.Context) {
	var input user.GetUserIDInput

	err := c.ShouldBindUri(&input)
	if err != nil {
		response := helper.MessageResponse("Failed to get user", http.StatusBadRequest, "error", "Request Not Allowed")
		c.JSON(http.StatusBadRequest, response)
		return
	}

	result, err := h.userService.GetUserByID(input)
	if err != nil {
		response := helper.MessageResponse("Failed to get user", http.StatusBadRequest, "error", nil)
		c.JSON(http.StatusBadRequest, response)
		return
	}

	if result.ID == 0 {
		response := helper.MessageResponse("Failed to get user", http.StatusOK, "failed", "User not found")
		c.JSON(http.StatusOK, response)
		return
	}

	response := helper.MessageResponse("Get Users Success", http.StatusOK, "success", result)
	c.JSON(http.StatusOK, response)
}
