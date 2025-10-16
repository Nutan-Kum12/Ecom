package user

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/Nutan-Kum12/Ecom/types"
	"github.com/gorilla/mux"
)

func TestUserServiceHandlers(t *testing.T) {
	// Test user creation, retrieval, etc.
	userStore := &mockUserStore{}
	handler := NewHandler(userStore)
	t.Run("Should fail if the user payload is invalid", func(t *testing.T) {
		// Simulate an invalid user payload
		payload := types.RegisterUserPayload{
			FirstName: "asd",
			LastName:  "kkk",
			Email:     "grtek",
			Password:  "asd16",
		}
		marshalled, _ := json.Marshal(payload)
		req, err := http.NewRequest(http.MethodPost, "/register", bytes.NewBuffer(marshalled))
		// req, err := http.NewRequest(http.MethodPost, "/register", nil)
		if err != nil {
			t.Fatal(err)
		}
		rr := httptest.NewRecorder()
		router := mux.NewRouter()
		router.HandleFunc("/register", handler.handleRegister).Methods(http.MethodPost)
		router.ServeHTTP(rr, req)
		if rr.Code != http.StatusBadRequest {
			t.Errorf("Expected status %d but got %d", http.StatusBadRequest, rr.Code)
		}
	})
	t.Run("Should correctly create a new user", func(t *testing.T) {
		// Simulate a valid user payload
		payload := types.RegisterUserPayload{
			FirstName: "Test",
			LastName:  "User",
			Email:     "test@example.com",
			Password:  "password123",
		}
		marshalled, _ := json.Marshal(payload)
		req, err := http.NewRequest(http.MethodPost, "/register", bytes.NewBuffer(marshalled))
		if err != nil {
			t.Fatal(err)
		}
		rr := httptest.NewRecorder()
		router := mux.NewRouter()
		router.HandleFunc("/register", handler.handleRegister).Methods(http.MethodPost)
		router.ServeHTTP(rr, req)
		if rr.Code != http.StatusCreated {
			t.Errorf("Expected status %d but got %d", http.StatusCreated, rr.Code)
		}
	})
}

type mockUserStore struct{}

func (m *mockUserStore) GetUserByEmail(email string) (*types.User, error) {
	// return &types.User{ID: 1, FirstName: "Test", LastName: "User", Email: email}, nil
	return nil, fmt.Errorf("user not found") // Simulate user not found
}
func (m *mockUserStore) GetUserById(id int) (*types.User, error) {
	// return &types.User{ID: id, FirstName: "Test", LastName: "User", Email: "test@example.com"}, nil
	return nil, fmt.Errorf("user not found") // Simulate user not found
}
func (m *mockUserStore) CreateUser(user *types.User) error {
	return nil
}
