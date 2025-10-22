package product

import (
	"net/http"

	"github.com/Nutan-Kum12/Ecom/types"
	"github.com/gorilla/mux"
)

type Handler struct {
	store types.ProductStore
}

func NewHandler(store types.ProductStore) *Handler {
	return &Handler{store: store}
}
func (h *Handler) RegisterRoutes(router *mux.Router) {

}
func (h *Handler) handleCreateProduct(w http.ResponseWriter, r *http.Request) {

}
