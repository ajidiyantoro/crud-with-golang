package user

type CreateUserResponse struct {
	ID          int    `json:"id"`
	Name        string `json:"name"`
	Gender      string `json:"gender"`
	Dateofbirth string `json:"dateofbirth"`
	Email       string `json:"email"`
	Avatar      string `json:"avatar"`
}

type GetUsersResponse struct {
	ID     int    `json:"id"`
	Name   string `json:"name"`
	Gender string `json:"gender"`
	Email  string `json:"email"`
	Avatar string `json:"avatar"`
}

func CreateUserFormat(user User) CreateUserResponse {
	formater := CreateUserResponse{
		ID:          user.ID,
		Name:        user.Name,
		Gender:      user.Gender,
		Dateofbirth: user.Dateofbirth,
		Email:       user.Email,
		Avatar:      user.Avatar,
	}
	return formater
}

// func GetUserFormat(user User) []GetUsersResponse {
// 	formater := []GetUsersResponse{
// 		ID:     user.ID,
// 		Name:   user.Name,
// 		Gender: user.Gender,
// 		Email:  user.Email,
// 		Avatar: user.Avatar,
// 	}
// 	return formater
// }
