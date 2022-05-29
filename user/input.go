package user

type CreateUserInput struct {
	Name        string `json:"name" binding:"required"`
	Gender      string `json:"gender" binding:"required"`
	Dateofbirth string `json:"dateofbirth" binding:"required"`
	Email       string `json:"email" binding:"required,email"`
	Password    string `json:"password" binding:"required"`
}

type CheckEmailInput struct {
	Email string `json:"email" binding:"required,email"`
}

type GetUserIDInput struct {
	ID int `uri:"id" binding:"required"`
}
