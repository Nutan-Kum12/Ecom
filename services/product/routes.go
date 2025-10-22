package product

import (
	"fmt"
	"log"
	"net/http"

	"github.com/Nutan-Kum12/Ecom/types"
	"github.com/Nutan-Kum12/Ecom/utils"
	"github.com/go-playground/validator"
	"github.com/gorilla/mux"
)

type Handler struct {
	store types.ProductStore
}

func NewHandler(store types.ProductStore) *Handler {
	return &Handler{store: store}
}
func (h *Handler) RegisterRoutes(router *mux.Router) {
	router.HandleFunc("/products", h.handleGetProducts).Methods(http.MethodGet)
	router.HandleFunc("/products", h.handleCreateProduct).Methods("POST")
}
func (h *Handler) handleGetProducts(w http.ResponseWriter, r *http.Request) {
	ps, err := h.store.GetProducts()
	if err != nil {
		utils.WriteError(w, http.StatusInternalServerError, err)
		return
	}
	utils.WriteJSON(w, http.StatusOK, ps)
}

func (h *Handler) handleCreateProduct(w http.ResponseWriter, r *http.Request) {
	var product types.ProductPayload

	log.Printf("Request body available: %v", r.Body != nil)
	log.Printf("Content-Type: %s", r.Header.Get("Content-Type"))
	if err := utils.ParseJSON(r, &product); err != nil {
		utils.WriteError(w, http.StatusBadRequest, err)
		return
	}
	// Always log what was parsed
	log.Printf("Parsed payload: %+v", product)
	log.Printf("Name: '%s', Description: '%s', Price: '%f', Quantity: '%d', ImageURL: '%s'",
		product.Name, product.Description, product.Price, product.Quantity, product.ImageURL)

	//validate payload
	if err := utils.Validate.Struct(product); err != nil {
		log.Printf("Validation error: %v", err)
		error := err.(validator.ValidationErrors)
		utils.WriteError(w, http.StatusBadRequest, fmt.Errorf("invalid payload %v", error))
		return
	}

	err := h.store.CreateProduct(&types.Product{
		Name:        product.Name,
		Description: product.Description,
		Price:       product.Price,
		Quantity:    product.Quantity,
		ImageURL:    product.ImageURL,
	})
	if err != nil {
		utils.WriteError(w, http.StatusInternalServerError, err)
		return
	}
	utils.WriteJSON(w, http.StatusCreated, map[string]string{"message": "Product created successfully"})
}
