package cart

import (
	"fmt"
	"net/http"

	"github.com/Nutan-Kum12/Ecom/services/auth"
	"github.com/Nutan-Kum12/Ecom/types"
	"github.com/Nutan-Kum12/Ecom/utils"
	"github.com/go-playground/validator"
	"github.com/gorilla/mux"
)

type Handler struct {
	store        types.OrderStore
	productstore types.ProductStore
	userstore    types.UserStore
}

func NewHandler(store types.OrderStore, productstore types.ProductStore, userstore types.UserStore) *Handler {
	return &Handler{store: store, productstore: productstore, userstore: userstore}
}

func (h *Handler) RegisterRoutes(router *mux.Router) {
	router.HandleFunc("/cart/checkout", auth.WithJWTAuth(h.handleCheckout, h.userstore)).Methods(http.MethodPost)
}

func (h *Handler) handleCheckout(w http.ResponseWriter, r *http.Request) {
	// Get user ID from JWT token context
	userID := auth.GetUserIDFromContext(r.Context())
	if userID == -1 {
		utils.WriteError(w, http.StatusUnauthorized, fmt.Errorf("user not authenticated"))
		return
	}
	
	// Get JSON payload
	var cart types.CartCheckoutPayload
	if err := utils.ParseJSON(r, &cart); err != nil {
		utils.WriteError(w, http.StatusBadRequest, err)
		return
	}
	// Validate payload
	if err := utils.Validate.Struct(cart); err != nil {
		errors := err.(validator.ValidationErrors)
		utils.WriteError(w, http.StatusBadRequest, fmt.Errorf("invalid payload: %v", errors))
		return
	}
	//get product details and calculate total amount
	productsIDs, err := getCartItemIDs(cart.Items)
	if err != nil {
		utils.WriteError(w, http.StatusBadRequest, err)
		return
	}
	ps, err := h.productstore.GetProductsByID(productsIDs)
	if err != nil {
		utils.WriteError(w, http.StatusBadRequest, err)
		return
	}
	orderId, totalAmount, err := h.createOrder(ps, cart.Items, userID)
	if err != nil {
		utils.WriteError(w, http.StatusInternalServerError, err)
		return
	}
	utils.WriteJSON(w, http.StatusOK, map[string]interface{}{
		"order_id":     orderId,
		"total_amount": totalAmount,
	})
}
