package order

import (
	"database/sql"

	"github.com/Nutan-Kum12/Ecom/types"
)

type Store struct {
	db *sql.DB
}

func NewStore(db *sql.DB) *Store {
	return &Store{db: db}
}
func (s *Store) CreateOrder(order *types.Order) (int, error) {
	result, err := s.db.Exec("INSERT INTO orders(user_id, total, status, address) VALUES(?, ?, ?, ?)",
		order.UserID, order.Total, order.Status, order.Address)
	if err != nil {
		return 0, err
	}
	id, err := result.LastInsertId()
	if err != nil {
		return 0, err
	}
	return int(id), nil
}
func (s *Store) CreateOrderItem(orderItem *types.OrderItem) error {
	_, err := s.db.Exec("INSERT INTO order_items(order_id, product_id, quantity, price) VALUES(?, ?, ?, ?)",
		orderItem.OrderID, orderItem.ProductID, orderItem.Quantity, orderItem.Price)
	return err
}
