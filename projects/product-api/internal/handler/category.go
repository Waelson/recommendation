package handler

import (
	"encoding/json"
	"github.com/Waelson/recommendation/ecommerce-api/internal/repository"
	"net/http"
)

type CategoryHandler struct {
	repository repository.CategoryRepository
}

func (c *CategoryHandler) ListAll(w http.ResponseWriter, r *http.Request) {

	categories, err := c.repository.ListCategories()
	if err != nil {
		http.Error(w, "Categorias não encontrado", http.StatusNotFound)
		return
	}

	w.Header().Set("Content-Type", "application/json")

	if err := json.NewEncoder(w).Encode(categories); err != nil {
		http.Error(w, "Erro ao codificar as categorias", http.StatusInternalServerError)
		return
	}

}

func NewCategoryHandler(repository repository.CategoryRepository) CategoryHandler {
	return CategoryHandler{
		repository: repository,
	}
}
