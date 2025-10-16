package api

import (
	"database/sql"
	"log"
	"net/http"

	"github.com/Nutan-Kum12/Ecom/services/user"
	"github.com/gorilla/mux"
)

type APIserver struct {
	addr string
	db   *sql.DB
}

func NewAPIserver(addr string, db *sql.DB) *APIserver {
	return &APIserver{
		addr: addr,
		db:   db,
	}
}
func (s *APIserver) Run() error {
	router := mux.NewRouter()
	subrouter := router.PathPrefix("/api/v1").Subrouter()
	userStore := user.NewStore(s.db)
	userHandler := user.NewHandler(userStore)
	userHandler.RegisterRoutes(subrouter)
	log.Println("Starting server on", s.addr)
	return http.ListenAndServe(s.addr, router)
}
