package user

type CreateUserInput struct {
	Name        string `json:"name"`
	Gender      string `json:"gender"`
	Dateofbirth string `json:"dateofbirth"`
	Email       string `json:"email"`
	Password    string `json:"password"`
}

type CheckEmailInput struct {
	Email string `json:"email" binding:"required,email"`
}
