package cart

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
	result, err := s.db.Exec("INSERT INTO orders(user_id, total_amount) VALUES(?, ?)",
		order.UserID, order.Total)
	if err != nil {
		return 0, err
	}
	id, err := result.LastInsertId()
	if err != nil {
		return 0, err
	}
	return int(id), nil
}
