package handler

import (
	"encoding/json"
	"github.com/Waelson/recommendation/ecommerce-api/internal/repository"
	"github.com/go-chi/chi/v5"
	"net/http"
)

type ProductHandler struct {
	repository repository.ProductRepository
}

func (h *ProductHandler) FindByID(w http.ResponseWriter, r *http.Request) {
	idParam := chi.URLParam(r, "id")
	product, err := h.repository.FindByID(idParam)

	if err != nil {
		http.Error(w, "Produto não encontrado", http.StatusNotFound)
		return
	}

	w.Header().Set("Content-Type", "application/json")

	if err := json.NewEncoder(w).Encode(product); err != nil {
		http.Error(w, "Erro ao codificar o produto", http.StatusInternalServerError)
		return
	}

}

func (h *ProductHandler) FindByCategoryID(w http.ResponseWriter, r *http.Request) {
	categoryIDParam := chi.URLParam(r, "categoryId")

	filteredProducts, err := h.repository.FindByCategoryID(categoryIDParam)

	if err != nil {
		http.Error(w, "Produto não encontrado", http.StatusNotFound)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	if err := json.NewEncoder(w).Encode(filteredProducts); err != nil {
		http.Error(w, "Erro ao codificar os produtos", http.StatusInternalServerError)
		return
	}
}

func NewProductHandler(repository repository.ProductRepository) *ProductHandler {
	return &ProductHandler{
		repository: repository,
	}
}
