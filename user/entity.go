package user

import "time"

type User struct {
	ID          int
	Name        string
	Gender      string
	Dateofbirth string
	Email       string
	Password    string
	Role        string
	Avatar      string
	IsActived   int
	CreatedAt   time.Time
	UpdatedAt   time.Time
}
