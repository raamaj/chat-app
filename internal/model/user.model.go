package model

import "time"

type UserResponse struct {
	ID        int64     `json:"id,omitempty" example:"1"`
	Username  string    `json:"username,omitempty" example:"johndoe"`
	Email     string    `json:"email,omitempty" example:"johndoe@example.com"`
	CreatedAt time.Time `json:"created_at,omitempty" example:"2020-01-01T00:00:00+09:00"`
}

type RegisterUserRequest struct {
	Username string `json:"username" validate:"required,max=20" example:"johndoe"`
	Email    string `json:"email" validate:"required,max=100" example:"johndoe@example.com"`
	Password string `json:"password" validate:"required,max=50" example:"123456"`
}
