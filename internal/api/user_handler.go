package api

import (
	"fmt"
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/google/uuid"
)

type UserHandler struct{}

func NewUserHandler() *UserHandler {
	return &UserHandler{}
}

func (uh *UserHandler) HandlerGetUserByID(w http.ResponseWriter, r *http.Request) {
	paramsHandlerID := chi.URLParam(r, "id")

	if paramsHandlerID == "" {
		http.NotFound(w, r)
		return
	}

	userID, err := uuid.Parse(paramsHandlerID)

	if err != nil {
		http.Error(w, "Invalid UUID", http.StatusBadRequest)
		return
	}

	fmt.Fprintf(w, "User ID: %s\n", userID)

}

func (uh *UserHandler) HandlerCreateUser(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Create User")
}
