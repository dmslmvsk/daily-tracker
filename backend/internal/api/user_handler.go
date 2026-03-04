package api

import (
	"encoding/json"
	"net/http"
	"time"

	"github.com/dmslmvsk/daily-tracker/backend/internal/repository"
	"github.com/dmslmvsk/daily-tracker/backend/internal/util"
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

type UserHandler struct {
	store *repository.Queries
}

func NewUserHandler(store *repository.Queries) *UserHandler{
	return &UserHandler{store:store}
}

func (h* UserHandler) CreateUser(w http.ResponseWriter, r *http.Request){
	var req CreateUserRequest
	err := json.NewDecoder(r.Body).Decode(&req)
	if err != nil{
		http.Error(w,err.Error(),http.StatusBadRequest)
		return
	}
	hashedPassword,err := util.HashPassword(req.Password)
	if err != nil{
		http.Error(w,err.Error(),http.StatusInternalServerError)
		return
	}

	arg:= repository.CreateUserParams{
		Email: req.Email,
		PasswordHash: hashedPassword,
	}
	user,err := h.store.CreateUser(r.Context(),arg)
	if err != nil{
		http.Error(w,err.Error(),http.StatusInternalServerError)
		return
	}
	resp:=CreateUserResponse{
		ID: user.ID,
		Email: user.Email,
		CreatedAt: user.CreatedAt,
	}
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(resp)
}