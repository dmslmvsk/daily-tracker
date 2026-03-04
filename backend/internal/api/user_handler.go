package api

import (
	"time"
)
type CreateUserRequest struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

type CreateUserResponse struct {
	ID int32 `json:"id"`
	Email string `json:"email"`
	CreatedAt time.Time `json:"created_at"`
}