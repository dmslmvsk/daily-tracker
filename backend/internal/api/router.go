package api

import (
	"net/http"

	"github.com/dmslmvsk/daily-tracker/backend/internal/repository"
)

func NewRouter(store *repository.Queries) *http.ServeMux {
	userHandler := NewUserHandler(store)
	mux:=http.NewServeMux()
	mux.HandleFunc("POST /users",userHandler.CreateUser)
	mux.HandleFunc("GET /hello",HelloWorld)
	return mux
}