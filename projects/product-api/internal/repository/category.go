package repository

import (
	"database/sql"
	"fmt"
	"github.com/Waelson/recommendation/ecommerce-api/internal/model"
)

type CategoryRepository struct {
	db *sql.DB
}

// NewCategoryRepository creates a new instance of CategoryRepository
func NewCategoryRepository(db *sql.DB) CategoryRepository {
	return CategoryRepository{db: db}
}

// ListCategories retrieves all categories in alphabetical order
func (r *CategoryRepository) ListCategories() ([]model.Category, error) {
	query := `SELECT name FROM category ORDER BY name ASC`

	rows, err := r.db.Query(query)
	if err != nil {
		return nil, fmt.Errorf("failed to retrieve categories: %v", err)
	}
	defer rows.Close()

	var categories []model.Category
	for rows.Next() {
		var category model.Category
		err := rows.Scan(&category.Name)
		if err != nil {
			return nil, fmt.Errorf("failed to scan category: %v", err)
		}
		categories = append(categories, category)
	}

	return categories, nil
}
