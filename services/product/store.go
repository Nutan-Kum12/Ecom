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
func (s *Store) CreateProduct(product *types.Product) error {
	_, err := s.db.Exec("INSERT INTO products(name, description, image_url, price, quantity) VALUES(?, ?, ?, ?, ?)",
		product.Name, product.Description, product.ImageURL, product.Price, product.Quantity)
	if err != nil {
		return err
	}
	return nil
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

// func (s *Store) GetProductByID(id int) (*types.Product, error) {
// 	rows, err := s.db.Query("SELECT * FROM products WHERE id = ?", id)
// 	if err != nil {
// 		return nil, err
// 	}
// 	defer rows.Close()

//		if rows.Next() {
//			return ScanRowIntoProduct(rows)
//		}
//		return nil, sql.ErrNoRows
//	}
func ScanRowIntoProduct(rows *sql.Rows) (*types.Product, error) {
	product := new(types.Product) // Create a new Product instance

	err := rows.Scan(
		&product.ID,
		&product.Name,
		&product.Description,
		&product.Price,
		&product.Quantity,
		&product.CreatedAt,
		&product.ImageURL,
	)
	if err != nil {
		return nil, err
	}
	return product, nil
}
