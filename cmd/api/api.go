package api

import (
	"database/sql"
	"log"
	"net/http"

	"github.com/Nutan-Kum12/Ecom/services/cart"
	"github.com/Nutan-Kum12/Ecom/services/order"
	"github.com/Nutan-Kum12/Ecom/services/product"
	"github.com/Nutan-Kum12/Ecom/services/user"
	"github.com/gorilla/mux"
)

type APIserver struct {
	addr string
	db   *sql.DB
}

// Go does not have traditional constructors like some other languages, but you can create a
//  function that initializes and returns a new instance of a struct.
//  This is often referred to as a "constructor" in Go, even though it's just a regular function.

func NewAPIserver(addr string, db *sql.DB) *APIserver {
	return &APIserver{
		addr: addr,
		db:   db,
	}
}
func (s *APIserver) Run() error {
	router := mux.NewRouter()
	subrouter := router.PathPrefix("/api/v1").Subrouter()

	userStore := user.NewStore(s.db)          // userStore now has database access and is injected
	userHandler := user.NewHandler(userStore) // userHandler can use userStore for DB operations and is injected
	userHandler.RegisterRoutes(subrouter)     // Register user routes

	productStore := product.NewStore(s.db)
	productHandler := product.NewHandler(productStore, userStore)
	productHandler.RegisterRoutes(subrouter)

	orderStore := order.NewStore(s.db)
	cartHandler := cart.NewHandler(orderStore, productStore, userStore)
	cartHandler.RegisterRoutes(subrouter)

	log.Println("Starting server on", s.addr)
	return http.ListenAndServe(s.addr, router)
}
