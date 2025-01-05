package repository

import (
	"database/sql"
	"fmt"
	"github.com/Waelson/recommendation/ecommerce-api/internal/model"
	"strings"
)

type ProductRepository struct {
	db *sql.DB
}

// NewProductRepository creates a new instance of ProductRepository
func NewProductRepository(db *sql.DB) ProductRepository {
	return ProductRepository{db: db}
}

// FindByID retrieves a single Product by its product_id
func (r *ProductRepository) FindByID(productID string) (*model.Product, error) {
	query := fmt.Sprintf(`SELECT product_id, name, description, price, 'A' as image, category FROM items WHERE product_id = '%s'`, productID)
	row := r.db.QueryRow(query)

	var product model.Product
	err := row.Scan(&product.ID, &product.Name, &product.Description, &product.Price, &product.Image, &product.CategoryID)
	if err != nil {
		if err == sql.ErrNoRows {
			return nil, nil // No product found
		}
		return nil, fmt.Errorf("failed to retrieve product: %v", err)
	}

	return &product, nil
}

// FindByProductIDs retrieves a list of Products by an array of product_ids
func (r *ProductRepository) FindByProductIDs(productIDs []string) ([]model.Product, error) {
	if len(productIDs) == 0 {
		return nil, nil
	}

	query := `SELECT product_id, name, description, price, 'A' as image, category FROM items WHERE product_id IN (`
	placeholders := make([]string, len(productIDs))
	args := make([]interface{}, len(productIDs))

	for i, id := range productIDs {
		placeholders[i] = "?"
		args[i] = id
	}

	query += strings.Join(placeholders, ",") + ")"

	rows, err := r.db.Query(query, args...)
	if err != nil {
		return nil, fmt.Errorf("failed to retrieve products: %v", err)
	}
	defer rows.Close()

	var products []model.Product
	for rows.Next() {
		var product model.Product
		err := rows.Scan(&product.ID, &product.Name, &product.Description, &product.Price, &product.Image, &product.CategoryID)
		if err != nil {
			return nil, fmt.Errorf("failed to scan product: %v", err)
		}
		products = append(products, product)
	}

	return products, nil
}

// FindByCategoryID retrieves a list of Products by category_id
func (r *ProductRepository) FindByCategoryID(categoryID string) ([]model.Product, error) {
	query := fmt.Sprintf(`SELECT product_id, name, description, price, 'A' as image, category FROM items WHERE category = '%s'  LIMIT 20`, categoryID)

	rows, err := r.db.Query(query)
	if err != nil {
		return nil, fmt.Errorf("failed to retrieve products by category: %v", err)
	}
	defer rows.Close()

	var products []model.Product
	for rows.Next() {
		var product model.Product
		err := rows.Scan(&product.ID, &product.Name, &product.Description, &product.Price, &product.Image, &product.CategoryID)
		if err != nil {
			return nil, fmt.Errorf("failed to scan product: %v", err)
		}
		products = append(products, product)
	}

	return products, nil
}
