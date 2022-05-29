package user

import (
	"crud-with-golang/helper"
)

type Service interface {
	CreateUser(input CreateUserInput) (User, error)
	IsEmailAvailable(input CheckEmailInput) (bool, error)
	GetUsers() ([]GetUsersResponse, error)
	GetUserByID(input GetUserIDInput) (GetUsersResponse, error)
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

func (s *service) GetUsers() ([]GetUsersResponse, error) {
	var dataUser GetUsersResponse
	var dataUserList []GetUsersResponse

	result, err := s.repository.GetUsers()
	if err != nil {
		return nil, err
	}

	for _, field := range result {
		dataUser = GetUsersResponse{
			ID:     field.ID,
			Name:   field.Name,
			Gender: field.Gender,
			Email:  field.Email,
			Avatar: field.Avatar,
		}
		dataUserList = append(dataUserList, dataUser)
	}

	return dataUserList, nil
}

func (s *service) GetUserByID(input GetUserIDInput) (GetUsersResponse, error) {
	result, _ := s.repository.GetUserByID(input.ID)

	output := GetUsersResponse{
		ID:     result.ID,
		Name:   result.Name,
		Gender: result.Gender,
		Email:  result.Email,
		Avatar: result.Avatar,
	}

	return output, nil
}
