package user

import (
	"fmt"
	"github.com/doneskkk/order-api/service/auth"
	"github.com/doneskkk/order-api/types"
	"github.com/doneskkk/order-api/utils"
	"github.com/go-chi/chi/v5"
	"github.com/go-playground/validator/v10"
	"net/http"
)

type Handler struct {
	repo types.UserRepo
}

func NewHandler(repo types.UserRepo) *Handler {
	return &Handler{repo: repo}
}

func (h *Handler) RegisterRoutes(router chi.Router) {
	router.Post("/login", h.handleLogin)
	router.Post("/register", h.handleRegister)
}

func (h *Handler) handleLogin(w http.ResponseWriter, r *http.Request) {
}

func (h *Handler) handleRegister(w http.ResponseWriter, r *http.Request) {

	//get JSON payload
	var newUser types.RegisterUser
	if err := utils.ParseJSON(r, &newUser); err != nil {
		utils.WriteError(w, http.StatusBadRequest, err)
		return
	}

	// validate the payload
	if err := utils.Validate.Struct(newUser); err != nil {
		errors := err.(validator.ValidationErrors)
		utils.WriteError(w, http.StatusBadRequest, fmt.Errorf("invalid payload %v", errors))
		return
	}
	//check if the user exists
	_, err := h.repo.GetUserByEmail(newUser.Email)
	if err == nil {
		utils.WriteError(w, http.StatusBadRequest, fmt.Errorf("user with email %s already exists", newUser.Email))
		return
	}

	hashedPassword, err := auth.HashPassword(newUser.Password)
	if err != nil {
		utils.WriteError(w, http.StatusInternalServerError, err)
		return
	}

	// if it doesn't we create the new user
	err = h.repo.CreateUser(types.User{
		FirstName: newUser.FirstName,
		LastName:  newUser.LastName,
		Email:     newUser.Email,
		Password:  hashedPassword,
	})

	if err != nil {
		utils.WriteError(w, http.StatusInternalServerError, err)
		return
	}

	utils.WriteJSON(w, http.StatusCreated, nil)
}
