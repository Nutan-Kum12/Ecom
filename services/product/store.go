package product

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
func (s *Store) GetProducts() ([]*types.Product, error) {
	rows, err := s.db.Query("SELECT * FROM products")
	if err != nil {
		return nil, err
	}
	var products []*types.Product // Slice to hold products
	for rows.Next() {
		p, err := ScanRowIntoProduct(rows)
		if err != nil {
			return nil, err
		}
		products = append(products, p) // Append product to slice
	}
	return products, nil
}
func ScanRowIntoProduct(rows *sql.Rows) (*types.Product, error) {
	product := new(types.Product) // Create a new Product instance

	err := rows.Scan(
		&product.ID,
		&product.Name,
		&product.Description,
		&product.ImageURL,
		&product.Quantity,
		&product.Price,
		&product.CreatedAt,
	)
	if err != nil {
		return nil, err
	}
	return product, nil
}
