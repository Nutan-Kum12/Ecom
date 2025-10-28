package product

import (
	"database/sql"
	"fmt"
	"strings"

	"github.com/Nutan-Kum12/Ecom/types"
)

// Generic function to convert any slice to []interface{}
// use for SQL queries
func sliceToInterface[T any](slice []T) []interface{} {
	result := make([]interface{}, len(slice))
	for i, v := range slice {
		result[i] = v
	}
	return result
}

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

func (s *Store) GetProductsByID(productIDs []int) ([]*types.Product, error) {
	placeholder := strings.Repeat(",?", len(productIDs)-1)
	// placeholder = strings.TrimRight(placeholder, ",")    //
	query := fmt.Sprintf("SELECT * FROM products WHERE id IN (?%s)", placeholder) //
	//convert produuctIDs to interface slice
	// args := make([]interface{}, len(productIDs)) //
	// for i, v := range productIDs {
	// 	args[i] = v
	// }

	//convert productIDs to interface slice using generics
	args := sliceToInterface(productIDs)
	rows, err := s.db.Query(query, args...)
	if err != nil {
		return nil, err
	}
	products := []*types.Product{} // Slice to hold products
	// Iterate through the rows and scan each product
	for rows.Next() {
		p, err := ScanRowIntoProduct(rows)
		if err != nil {
			return nil, err
		}
		products = append(products, p)
	}
	return products, nil
}
func (s *Store) GetProductByID(productID int) (*types.Product, error) {
	rows, err := s.db.Query("SELECT * FROM products WHERE id = ?", productID)
	if err != nil {
		return nil, err
	}

	p := new(types.Product)
	for rows.Next() {
		p, err = ScanRowIntoProduct(rows)
		if err != nil {
			return nil, err
		}
	}

	return p, nil
}

func (s *Store) UpdateProduct(product *types.Product) error {
	_, err := s.db.Exec("UPDATE products SET name = ?, description = ?, image_url = ?, price = ?, quantity = ? WHERE id = ?",
		product.Name, product.Description, product.ImageURL, product.Price, product.Quantity, product.ID)
	if err != nil {
		return err
	}
	return nil
}
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
