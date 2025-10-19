package user

import (
	"fmt"
	"log"
	"net/http"

	"github.com/Nutan-Kum12/Ecom/services/auth"
	"github.com/Nutan-Kum12/Ecom/types"
	"github.com/Nutan-Kum12/Ecom/utils"
	"github.com/go-playground/validator"
	"github.com/gorilla/mux"
)

type Handler struct {
	// You can add dependencies like Store here if needed
	store types.UserStore
}

func NewHandler(store types.UserStore) *Handler {
	return &Handler{store: store}
}
func (h *Handler) RegisterRoutes(router *mux.Router) {
	// Define user-related routes here
	router.HandleFunc("/login", h.handleLogin).Methods("POST")
	router.HandleFunc("/register", h.handleRegister).Methods("POST")
}
func (h *Handler) handleLogin(w http.ResponseWriter, r *http.Request) {
	//get json payload
	log.Println("=== Login Handler Start ===")
	var payload types.LoginUserPayload

	log.Printf("Request body available: %v", r.Body != nil)
	log.Printf("Content-Type: %s", r.Header.Get("Content-Type"))

	if err := utils.ParseJSON(r, &payload); err != nil {
		log.Printf("ParseJSON error: %v", err)
		utils.WriteError(w, http.StatusBadRequest, fmt.Errorf("invalid JSON: %v", err))
		return
	}
	u, err := h.store.GetUserByEmail(payload.Email)
	if err != nil {
		log.Printf("GetUserByEmail error: %v", err)
		utils.WriteError(w, http.StatusUnauthorized, fmt.Errorf("invalid email or password"))
		return
	}

}
func (h *Handler) handleRegister(w http.ResponseWriter, r *http.Request) {
	//get json payload
	log.Println("=== Register Handler Start ===")
	var payload types.RegisterUserPayload

	log.Printf("Request body available: %v", r.Body != nil)
	log.Printf("Content-Type: %s", r.Header.Get("Content-Type"))

	if err := utils.ParseJSON(r, &payload); err != nil {
		log.Printf("ParseJSON error: %v", err)
		utils.WriteError(w, http.StatusBadRequest, fmt.Errorf("invalid JSON: %v", err))
		return
	}

	// Always log what was parsed
	log.Printf("Parsed payload: %+v", payload)
	log.Printf("FirstName: '%s', LastName: '%s', Email: '%s', Password: '%s'",
		payload.FirstName, payload.LastName, payload.Email, payload.Password)

	//validate payload
	if err := utils.Validate.Struct(payload); err != nil {
		log.Printf("Validation error: %v", err)
		error := err.(validator.ValidationErrors)
		utils.WriteError(w, http.StatusBadRequest, fmt.Errorf("invalid payload %v", error))
		return
	}
	//check if user already exists
	_, err := h.store.GetUserByEmail(payload.Email)
	if err == nil {
		utils.WriteError(w, http.StatusConflict, fmt.Errorf("user with email %s already exists", payload.Email))
		return
	}
	hashPassword, err := auth.HashPassword(payload.Password)
	if err != nil {
		utils.WriteError(w, http.StatusInternalServerError, err)
		return
	}
	//if it doesn't we create NEW user
	err = h.store.CreateUser(&types.User{
		FirstName: payload.FirstName,
		LastName:  payload.LastName,
		Email:     payload.Email,
		Password:  hashPassword,
	})
	if err != nil {
		utils.WriteError(w, http.StatusInternalServerError, err)
		return
	}
	utils.WriteJSON(w, http.StatusCreated, nil)
}
