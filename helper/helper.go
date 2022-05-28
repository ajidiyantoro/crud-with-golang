package helper

import (
	"golang.org/x/crypto/bcrypt"
)

type MgsResponse struct {
	Meta Meta        `json:"meta"`
	Data interface{} `json:"data"`
}

type Meta struct {
	Message string `json:"message"`
	Code    int    `json:"code"`
	Status  string `json:"status"`
}

func HashPassword(password string) string {
	passwordHash, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.MinCost)
	if err != nil {
		return err.Error()
	}

	return string(passwordHash)
}

func MessageResponse(message string, code int, status string, data interface{}) MgsResponse {
	meta := Meta{
		Message: message,
		Code:    code,
		Status:  status,
	}

	jsonResponse := MgsResponse{
		Meta: meta,
		Data: data,
	}

	return jsonResponse
}
