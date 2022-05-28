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
		errorMessage := gin.H{"errors": err}
		c.JSON(http.StatusBadRequest, errorMessage)
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
		c.JSON(http.StatusBadRequest, errorMessage)
		return
	}

	formater := user.CreateUserFormat(newUser)
	response := helper.MessageResponse("Account has been registered", http.StatusOK, "success", formater)

	c.JSON(http.StatusOK, response)
}
