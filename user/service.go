package user

import (
	"crud-with-golang/helper"
)

type Service interface {
	CreateUser(input CreateUserInput) (User, error)
	IsEmailAvailable(input CheckEmailInput) (bool, error)
}

type service struct {
	repository Repository
}

func NewService(repository Repository) *service {
	return &service{repository}
}

func (s *service) CreateUser(input CreateUserInput) (User, error) {
	passwordHash := helper.HashPassword(input.Password)

	inputData := User{
		Name:        input.Name,
		Gender:      input.Gender,
		Dateofbirth: input.Dateofbirth,
		Email:       input.Email,
		Password:    passwordHash,
	}

	newUser, err := s.repository.CreateUser(inputData)
	if err != nil {
		return newUser, err
	}

	return newUser, nil
}

func (s *service) IsEmailAvailable(input CheckEmailInput) (bool, error) {
	dataEmail := input.Email

	result, err := s.repository.GetUserByEmail(dataEmail)
	if err != nil {
		return false, err
	}

	if result.ID == 0 {
		return true, nil
	}

	return false, nil
}
