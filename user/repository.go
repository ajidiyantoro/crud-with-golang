package user

import "gorm.io/gorm"

type Repository interface {
	CreateUser(user User) (User, error)
	GetUserByID(id int) (User, error)
	GetUserByEmail(input string) (User, error)
	GetUsers() ([]User, error)
	UpdateUser(user User) (User, error)
}

type repository struct {
	db *gorm.DB
}

func NewRepository(db *gorm.DB) *repository {
	return &repository{db}
}

func (r *repository) CreateUser(user User) (User, error) {
	err := r.db.Create(&user).Error
	if err != nil {
		return user, err
	}

	return user, nil
}

func (r *repository) GetUserByID(id int) (User, error) {
	var dataUser User

	err := r.db.Where("id = ?", id).Find(&dataUser).Error
	if err != nil {
		return dataUser, err
	}

	return dataUser, nil
}

func (r *repository) GetUserByEmail(input string) (User, error) {
	var dataUser User
	err := r.db.Where("email = ?", input).Find(&dataUser).Error
	if err != nil {
		return dataUser, err
	}

	return dataUser, nil
}

func (r *repository) GetUsers() ([]User, error) {
	var dataUserList []User

	err := r.db.Find(&dataUserList).Error
	if err != nil {
		return dataUserList, err
	}

	return dataUserList, nil
}

func (r *repository) UpdateUser(user User) (User, error) {
	var dataUser User

	err := r.db.Save(&dataUser).Error
	if err != nil {
		return dataUser, nil
	}

	return dataUser, nil
}
