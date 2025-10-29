package cart

import (
	"fmt"

	"github.com/Nutan-Kum12/Ecom/types"
)

func getCartItemIDs(items []types.CartItem) ([]int, error) {
	ids := make([]int, len(items)) // Initialize slice to hold product IDs
	for i, item := range items {
		if item.Quantity <= 0 {
			return nil, fmt.Errorf("invalid quantity for product ID %d", item.ProductID)
		}
		ids[i] = item.ProductID
	}
	return ids, nil
}
func (h *Handler) createOrder(ps []*types.Product, items []types.CartItem, userID int) (int, float64, error) {
	productMap := make(map[int]*types.Product) //map product ID to product
	for _, p := range ps {
		productMap[p.ID] = p
	}
	//check if all products are available in sufficient quantity
	if err := checkIfCartInStock(productMap, items); err != nil {
		return 0, 0, err
	}
	//calculate total amount
	var totalAmount float64
	for _, item := range items {
		product, ok := productMap[item.ProductID]
		if !ok {
			return 0, 0, fmt.Errorf("product ID %d not found", item.ProductID)
		}
		totalAmount += product.Price * float64(item.Quantity)
	}
	//reduce product quantity
	for _, item := range items {
		product, ok := productMap[item.ProductID]
		if !ok {
			return 0, 0, fmt.Errorf("product ID %d not found", item.ProductID)
		}
		product.Quantity -= item.Quantity
		h.productstore.UpdateProduct(product)
	}
	//create order and order items
	orderID, err := h.store.CreateOrder(&types.Order{
		UserID:  userID,
		Total:   totalAmount,
		Status:  "pending",
		Address: "some address",
	})
	if err != nil {
		return 0, 0, err
	}
	for _, item := range items {
		product, ok := productMap[item.ProductID]
		if !ok {
			return 0, 0, fmt.Errorf("product ID %d not found", item.ProductID)
		}
		orderItem := &types.OrderItem{
			OrderID:   orderID,
			ProductID: product.ID,
			Quantity:  item.Quantity,
			Price:     product.Price,
		}
		if err := h.store.CreateOrderItem(orderItem); err != nil {
			return 0, 0, err
		}
	}
	return orderID, totalAmount, nil
}

func checkIfCartInStock(productMap map[int]*types.Product, items []types.CartItem) error {
	if len(items) == 0 {
		return fmt.Errorf("no products found")
	}
	for _, item := range items {
		product, ok := productMap[item.ProductID]
		if !ok {
			return fmt.Errorf("product ID %d not found", item.ProductID)
		}
		if product.Quantity < item.Quantity {
			return fmt.Errorf("insufficient stock for product ID %d", item.ProductID)
		}
	}
	return nil
}
