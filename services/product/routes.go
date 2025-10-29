package product

import (
	"fmt"
	"log"
	"net/http"
	"strconv"

	"github.com/Nutan-Kum12/Ecom/services/auth"
	"github.com/Nutan-Kum12/Ecom/types"
	"github.com/Nutan-Kum12/Ecom/utils"
	"github.com/go-playground/validator"
	"github.com/gorilla/mux"
)

type Handler struct {
	store     types.ProductStore
	userstore types.UserStore
}

func NewHandler(store types.ProductStore, userstore types.UserStore) *Handler {
	return &Handler{store: store, userstore: userstore} // Store injected and userstore injected
}
func (h *Handler) RegisterRoutes(router *mux.Router) {
	router.HandleFunc("/products", h.handleGetProducts).Methods(http.MethodGet)
	router.HandleFunc("/products/{productID}", h.handleGetProduct).Methods(http.MethodGet)

	// admin routes
	router.HandleFunc("/products", auth.WithJWTAuth(h.handleCreateProduct, h.userstore)).Methods(http.MethodPost)
}
func (h *Handler) handleGetProducts(w http.ResponseWriter, r *http.Request) {
	ps, err := h.store.GetProducts()
	if err != nil {
		utils.WriteError(w, http.StatusInternalServerError, err)
		return
	}
	utils.WriteJSON(w, http.StatusOK, ps)
}
func (h *Handler) handleGetProduct(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	str, ok := vars["productID"]
	if !ok {
		utils.WriteError(w, http.StatusBadRequest, fmt.Errorf("missing product ID"))
		return
	}

	productID, err := strconv.Atoi(str)
	if err != nil {
		utils.WriteError(w, http.StatusBadRequest, fmt.Errorf("invalid product ID"))
		return
	}

	product, err := h.store.GetProductByID(productID)
	if err != nil {
		utils.WriteError(w, http.StatusInternalServerError, err)
		return
	}

	utils.WriteJSON(w, http.StatusOK, product)
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
