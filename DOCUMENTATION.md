# Go E-commerce Project - Complete Beginner's Guide

## Table of Contents
1. [Project Overview](#project-overview)
2. [Project Structure](#project-structure)
3. [Understanding Go Structs](#understanding-go-structs)
4. [Configuration System](#configuration-system)
5. [Database Layer](#database-layer)
6. [Types and Data Models](#types-and-data-models)
7. [API Layer](#api-layer)
8. [User Service](#user-service)
9. [Authentication](#authentication)
10. [Testing](#testing)
11. [Complete Flow Examples](#complete-flow-examples)
12. [Common Patterns](#common-patterns)

---

## Project Overview

This is a **REST API e-commerce backend** built in Go. It handles:
- User registration and authentication
- Product management
- Shopping cart functionality
- Order processing

### What is a Struct?
A **struct** in Go is like a blueprint for creating objects. Think of it as a container that groups related data together.

```go
// A struct is like a form with fields
type Person struct {
    Name string  // Field 1
    Age  int     // Field 2
}

// Creating an instance (filling out the form)
person := Person{
    Name: "John",
    Age:  25,
}
```

---

## Project Structure

```
Ecom/
├── cmd/                    # Application entry points
│   ├── main.go            # 🚀 START HERE - Application startup
│   └── api/
│       └── api.go         # HTTP server setup
├── config/                # Configuration management
│   └── env.go            # Environment variables
├── db/                   # Database connection
│   └── db.go            # MySQL connection setup
├── types/               # Data models (structs)
│   └── types.go        # User, Product structs
├── services/           # Business logic
│   ├── user/          # User-related operations
│   │   ├── routes.go  # User API handlers
│   │   └── routes_test.go # Tests
│   └── auth/         # Authentication logic
├── utils/            # Helper functions
├── .env             # Environment variables file
├── Makefile        # Build commands
└── go.mod         # Dependencies
```

---

## Understanding Go Structs

### 1. Basic Struct Definition

```go
// Define a struct (like creating a template)
type User struct {
    ID        int       `json:"id"`         // Tags for JSON conversion
    FirstName string    `json:"first_name"` // Field name in JSON
    LastName  string    `json:"last_name"`
    Email     string    `json:"email"`
    Password  string    `json:"-"`          // "-" means hide from JSON
    CreatedAt time.Time `json:"created_at"`
}
```

**What each part means:**
- `type User struct` - Creates a new type called "User"
- `ID int` - A field named "ID" that stores integers
- `json:"id"` - When converting to JSON, use "id" as the field name
- `json:"-"` - Don't include this field in JSON output (security)

### 2. Creating and Using Structs

```go
// Method 1: Create with field names
user := User{
    FirstName: "John",
    LastName:  "Doe",
    Email:     "john@example.com",
    Password:  "hashedpassword",
}

// Method 2: Create empty and set fields
var user User
user.FirstName = "John"
user.Email = "john@example.com"

// Method 3: Create pointer to struct
user := &User{
    FirstName: "John",
    Email:     "john@example.com",
}
```

### 3. Methods on Structs

```go
// You can add functions (methods) to structs
func (u *User) GetFullName() string {
    return u.FirstName + " " + u.LastName
}

// Usage
user := User{FirstName: "John", LastName: "Doe"}
fullName := user.GetFullName() // Returns "John Doe"
```

---

## Configuration System

### File: `config/env.go`

```go
package config

import (
    "fmt"
    "os"
    "github.com/joho/godotenv"
)

// 1. Define configuration struct
type Config struct {
    PublicHost string  // Server host
    Port       string  // Server port
    DBUser     string  // Database username
    DBPassword string  // Database password
    DBName     string  // Database name
    DBAddress  string  // Database address
}

// 2. Global variable to hold config
var Envs = initConfig()

// 3. Function to load configuration
func initConfig() Config {
    godotenv.Load()  // Load .env file
    
    return Config{
        PublicHost: getEnv("PUBLIC_HOST", "http://localhost"),
        Port:       getEnv("PORT", "8080"),
        DBUser:     getEnv("DB_USER", "root"),
        DBPassword: getEnv("DB_PASSWORD", ""),
        DBName:     getEnv("DB_NAME", "ecom"),
        DBAddress:  fmt.Sprintf("%s:%s", 
                     getEnv("DB_HOST", "127.0.0.1"), 
                     getEnv("DB_PORT", "3306")),
    }
}

// 4. Helper function to get environment variables
func getEnv(key, fallback string) string {
    if value, ok := os.LookupEnv(key); ok {
        return value
    }
    return fallback  // Use default if not found
}
```

**How it works:**
1. **Struct Definition**: `Config` struct holds all configuration values
2. **Global Variable**: `Envs` is automatically initialized when package loads
3. **Environment Loading**: `initConfig()` reads from `.env` file
4. **Fallback Values**: If environment variable doesn't exist, use defaults

**Why use this pattern?**
- Centralized configuration
- Environment-specific settings (dev, staging, prod)
- Security (sensitive data in environment variables)

---

## Database Layer

### File: `db/db.go`

```go
package db

import (
    "database/sql"
    "log"
    "github.com/go-sql-driver/mysql"
)

// Function to create database connection
func NewMySQLStorage(cfg mysql.Config) (*sql.DB, error) {
    // cfg.FormatDSN() creates connection string like:
    // "user:password@tcp(localhost:3306)/database"
    db, err := sql.Open("mysql", cfg.FormatDSN())
    if err != nil {
        log.Fatal(err)
    }
    return db, nil
}
```

**How it's used in main.go:**

```go
// Create database configuration
dbConfig := mysql.Config{
    User:                 config.Envs.DBUser,     // "root"
    Passwd:               config.Envs.DBPassword, // ""
    Net:                  "tcp",
    Addr:                 config.Envs.DBAddress,  // "127.0.0.1:3306"
    DBName:               config.Envs.DBName,     // "ecom"
    AllowNativePasswords: true,
    ParseTime:            true,
}

// Create connection
db, err := db.NewMySQLStorage(dbConfig)
```

---

## Types and Data Models

### File: `types/types.go`

```go
package types

import "time"

// 1. DATABASE MODEL - Represents table structure
type User struct {
    ID        int       `json:"id"`
    FirstName string    `json:"first_name"`
    LastName  string    `json:"last_name"`
    Email     string    `json:"email"`
    Password  string    `json:"-"`          // Hidden from JSON responses
    CreatedAt time.Time `json:"created_at"`
}

// 2. REQUEST MODEL - What client sends for registration
type RegisterUserPayload struct {
    FirstName string `json:"first_name"`
    LastName  string `json:"last_name"`
    Email     string `json:"email"`
    Password  string `json:"password"`
    // Note: No ID or CreatedAt (auto-generated)
}

// 3. INTERFACE - Contract for database operations
type UserStore interface {
    GetUserByEmail(email string) (*User, error)
    GetUserById(id int) (*User, error)
    CreateUser(user *User) error
}
```

**Why different structs for same data?**

1. **Security**: `User` hides password in JSON, `RegisterUserPayload` doesn't
2. **Validation**: Different rules for input vs stored data
3. **API Design**: Client doesn't need to send ID (auto-generated)

**Example Usage:**

```go
// Client sends this JSON:
{
    "first_name": "John",
    "last_name": "Doe", 
    "email": "john@example.com",
    "password": "mypassword"
}

// Server parses into RegisterUserPayload
var payload RegisterUserPayload
json.Unmarshal(jsonData, &payload)

// Server converts to User for database
user := &User{
    FirstName: payload.FirstName,
    LastName:  payload.LastName,
    Email:     payload.Email,
    Password:  hashedPassword, // Hashed version
    CreatedAt: time.Now(),     // Current time
}

// Server responds with User (password hidden)
{
    "id": 123,
    "first_name": "John",
    "last_name": "Doe",
    "email": "john@example.com",
    "created_at": "2025-10-15T10:30:00Z"
    // Note: No password field
}
```

---

## API Layer

### File: `cmd/api/api.go`

```go
package api

import (
    "database/sql"
    "log"
    "net/http"
    "github.com/gorilla/mux"
)

// 1. API Server struct
type APIServer struct {
    addr string   // Server address ":8080"
    db   *sql.DB  // Database connection
}

// 2. Constructor function
func NewAPIserver(addr string, db *sql.DB) *APIServer {
    return &APIServer{
        addr: addr,
        db:   db,
    }
}

// 3. Start server method
func (s *APIServer) Run() error {
    router := mux.NewRouter()
    
    // Set up routes
    s.setupRoutes(router)
    
    log.Printf("Starting server on %s", s.addr)
    return http.ListenAndServe(s.addr, router)
}

// 4. Route setup
func (s *APIServer) setupRoutes(router *mux.Router) {
    // Health check
    router.HandleFunc("/health", s.handleHealth).Methods("GET")
    
    // API routes
    api := router.PathPrefix("/api/v1").Subrouter()
    
    // User routes
    userHandler := user.NewHandler(s.db)
    userHandler.RegisterRoutes(api)
}

// 5. Handler example
func (s *APIServer) handleHealth(w http.ResponseWriter, r *http.Request) {
    w.Header().Set("Content-Type", "application/json")
    w.WriteHeader(http.StatusOK)
    w.Write([]byte(`{"status": "ok"}`))
}
```

**Key Concepts:**

1. **Struct Methods**: `(s *APIServer)` means this function belongs to APIServer
2. **Dependency Injection**: Database is passed to the server
3. **Route Organization**: Different handlers for different features
4. **HTTP Status Codes**: 200 OK, 201 Created, 400 Bad Request, etc.

---

## User Service

### File: `services/user/routes.go`

```go
package user

import (
    "fmt"
    "log"
    "net/http"
    
    "github.com/Nutan-Kum12/Ecom/types"
    "github.com/Nutan-Kum12/Ecom/utils"
    "github.com/gorilla/mux"
)

// 1. Handler struct
type Handler struct {
    store types.UserStore  // Interface, not concrete type
}

// 2. Constructor
func NewHandler(store types.UserStore) *Handler {
    return &Handler{store: store}
}

// 3. Register routes
func (h *Handler) RegisterRoutes(router *mux.Router) {
    router.HandleFunc("/auth/register", h.handleRegister).Methods("POST")
    router.HandleFunc("/auth/login", h.handleLogin).Methods("POST")
}

// 4. Register handler
func (h *Handler) handleRegister(w http.ResponseWriter, r *http.Request) {
    // Step 1: Parse JSON into struct
    var payload types.RegisterUserPayload
    if err := utils.ParseJSON(r, &payload); err != nil {
        utils.WriteError(w, http.StatusBadRequest, err)
        return
    }
    
    // Step 2: Validate required fields
    if payload.FirstName == "" || payload.LastName == "" || 
       payload.Email == "" || payload.Password == "" {
        utils.WriteError(w, http.StatusBadRequest, 
                        fmt.Errorf("all fields are required"))
        return
    }
    
    // Step 3: Check if user already exists
    _, err := h.store.GetUserByEmail(payload.Email)
    if err == nil {
        utils.WriteError(w, http.StatusConflict, 
                        fmt.Errorf("user already exists"))
        return
    }
    
    // Step 4: Hash password
    hashPassword, err := auth.HashPassword(payload.Password)
    if err != nil {
        utils.WriteError(w, http.StatusInternalServerError, err)
        return
    }
    
    // Step 5: Create user
    user := &types.User{
        FirstName: payload.FirstName,
        LastName:  payload.LastName,
        Email:     payload.Email,
        Password:  hashPassword,
    }
    
    err = h.store.CreateUser(user)
    if err != nil {
        utils.WriteError(w, http.StatusInternalServerError, err)
        return
    }
    
    // Step 6: Success response
    utils.WriteJSON(w, http.StatusCreated, map[string]string{
        "message": "User created successfully",
    })
}
```

**Flow Explanation:**

1. **Input**: Client sends JSON with user data
2. **Parsing**: Convert JSON to Go struct
3. **Validation**: Check required fields
4. **Business Logic**: Check duplicates, hash password
5. **Storage**: Save to database
6. **Response**: Send success/error back to client

---

## Authentication

### Password Hashing

```go
package auth

import (
    "golang.org/x/crypto/bcrypt"
)

// Hash password before storing
func HashPassword(password string) (string, error) {
    hash, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
    return string(hash), err
}

// Verify password during login
func CheckPassword(password, hash string) bool {
    err := bcrypt.CompareHashAndPassword([]byte(hash), []byte(password))
    return err == nil
}
```

**Why hash passwords?**
- Security: Even if database is compromised, passwords are protected
- One-way: You can't reverse a hash to get original password
- Verification: You can verify if a password matches the hash

---

## Testing

### File: `services/user/routes_test.go`

```go
package user

import (
    "bytes"
    "encoding/json"
    "net/http"
    "net/http/httptest"
    "testing"
    
    "github.com/Nutan-Kum12/Ecom/types"
    "github.com/gorilla/mux"
)

// 1. Test function
func TestUserServiceHandlers(t *testing.T) {
    // Setup
    userStore := &mockUserStore{}
    handler := NewHandler(userStore)
    
    t.Run("Should fail if the user payload is invalid", func(t *testing.T) {
        // Create test payload
        payload := types.RegisterUserPayload{
            FirstName: "",    // Invalid - empty
            LastName:  "",    // Invalid - empty
            Email:     "",    // Invalid - empty
            Password:  "",    // Invalid - empty
        }
        
        // Convert to JSON
        marshalled, _ := json.Marshal(payload)
        
        // Create HTTP request
        req, err := http.NewRequest(http.MethodPost, "/register", 
                                   bytes.NewBuffer(marshalled))
        if err != nil {
            t.Fatal(err)
        }
        
        // Create response recorder
        rr := httptest.NewRecorder()
        
        // Set up router
        router := mux.NewRouter()
        router.HandleFunc("/register", handler.handleRegister).Methods("POST")
        
        // Execute request
        router.ServeHTTP(rr, req)
        
        // Check result
        if rr.Code != http.StatusBadRequest {
            t.Errorf("Expected status %d but got %d", 
                    http.StatusBadRequest, rr.Code)
        }
    })
}

// 2. Mock store for testing
type mockUserStore struct{}

func (m *mockUserStore) GetUserByEmail(email string) (*types.User, error) {
    return nil, fmt.Errorf("user not found")
}

func (m *mockUserStore) GetUserById(id int) (*types.User, error) {
    return nil, fmt.Errorf("user not found")
}

func (m *mockUserStore) CreateUser(user *types.User) error {
    return nil  // Success
}
```

**Testing Concepts:**

1. **Mock Objects**: Fake implementations for testing
2. **HTTP Testing**: Simulate HTTP requests/responses
3. **Test Cases**: Different scenarios (valid/invalid data)
4. **Assertions**: Check if results match expectations

---

## Complete Flow Examples

### 1. User Registration Flow

```
1. Client Request:
   POST /api/v1/auth/register
   {
       "first_name": "John",
       "last_name": "Doe",
       "email": "john@example.com", 
       "password": "mypassword"
   }

2. Server Processing:
   ├── Parse JSON → RegisterUserPayload struct
   ├── Validate fields (not empty)
   ├── Check if email exists → Query database
   ├── Hash password → bcrypt.GenerateFromPassword()
   ├── Create User struct
   └── Save to database → INSERT INTO users...

3. Server Response:
   HTTP 201 Created
   {
       "message": "User created successfully"
   }
```

### 2. Application Startup Flow

```
1. main.go starts
   ├── Load config (config.Envs)
   ├── Connect to database (db.NewMySQLStorage)
   ├── Create API server (api.NewAPIserver)
   └── Start HTTP server (server.Run)

2. Server running on :8080
   ├── Listening for HTTP requests
   ├── Router matches URLs to handlers
   └── Handlers process requests
```

---

## Common Patterns

### 1. Constructor Pattern

```go
// Instead of exposing struct fields directly
func NewHandler(store types.UserStore) *Handler {
    return &Handler{store: store}
}

// Benefits:
// - Validation during creation
// - Initialization logic
// - Consistent object creation
```

### 2. Interface Pattern

```go
// Define what you need, not how it's implemented
type UserStore interface {
    GetUserByEmail(email string) (*User, error)
    CreateUser(user *User) error
}

// Benefits:
// - Easy testing (mock implementations)
// - Flexible (switch database types)
// - Dependency inversion
```

### 3. Error Handling Pattern

```go
func (h *Handler) handleRegister(w http.ResponseWriter, r *http.Request) {
    var payload types.RegisterUserPayload
    
    // Early return on error
    if err := utils.ParseJSON(r, &payload); err != nil {
        utils.WriteError(w, http.StatusBadRequest, err)
        return  // Stop processing
    }
    
    // Continue with happy path...
}

// Benefits:
// - Clear error handling
// - Avoid deep nesting
// - Consistent error responses
```

### 4. Dependency Injection Pattern

```go
// Don't create dependencies inside structs
type Handler struct {
    store types.UserStore  // Injected dependency
}

// Inject from outside
func NewHandler(store types.UserStore) *Handler {
    return &Handler{store: store}
}

// Benefits:
// - Testable (inject mocks)
// - Flexible (different implementations)
// - Loose coupling
```

---

## Key Takeaways

### 1. **Structs are Data Containers**
- Group related data together
- Add methods for behavior
- Use tags for serialization

### 2. **Interfaces Define Contracts**
- What methods a type must have
- Enable polymorphism and testing
- Decouple dependencies

### 3. **Separation of Concerns**
- `config/` - Configuration management
- `db/` - Database connections
- `types/` - Data models
- `services/` - Business logic
- `cmd/` - Application entry points

### 4. **HTTP Request Lifecycle**
```
Request → Router → Handler → Business Logic → Database → Response
```

### 5. **Error Handling is Critical**
- Validate input early
- Return appropriate HTTP status codes
- Use consistent error response format

This documentation should give you a solid understanding of how Go structs, interfaces, and patterns work together to build a maintainable e-commerce API! 🚀