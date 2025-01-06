package handler

import (
	"encoding/json"
	"github.com/Waelson/recommendation/ecommerce-api/internal/repository"
	"github.com/Waelson/recommendation/ecommerce-api/internal/service"
	"github.com/go-chi/chi/v5"
	"net/http"
	"strconv"
)

type RecommendationHandler struct {
	recommendationService service.RecommendationService
	productRepository     repository.ProductRepository
}

func (h *RecommendationHandler) Recommend(w http.ResponseWriter, r *http.Request) {
	productIDParam := chi.URLParam(r, "productID")
	if productIDParam == "" {
		http.Error(w, "Parametro productID nao informado", http.StatusBadRequest)
		return
	}

	numItemsParam := chi.URLParam(r, "numItems")
	if numItemsParam == "" {
		http.Error(w, "Parametro numItems nao informado", http.StatusBadRequest)
		return
	}

	numItem, err := strconv.Atoi(numItemsParam)
	if err != nil {
		http.Error(w, "Parametro numItems e invalido", http.StatusBadRequest)
		return
	}

	recommendationResponse, err := h.recommendationService.GetRecommendations(productIDParam, numItem)
	if err != nil {
		http.Error(w, "Erro ao obter a lista de recomencacoes", http.StatusInternalServerError)
		return
	}

	if recommendationResponse == nil || len(recommendationResponse.Recommendations) == 0 {
		http.Error(w, "Erro ao obter a lista de recomencacoes", http.StatusNotFound)
		return
	}

	products, err := h.productRepository.FindByProductIDs(recommendationResponse.Recommendations)
	if err != nil {
		http.Error(w, "Erro ao obter a lista de recomendacoes", http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	if err := json.NewEncoder(w).Encode(products); err != nil {
		http.Error(w, "Erro ao codificar os produtos", http.StatusInternalServerError)
		return
	}

}

func NewRecommendationHandler(recommendationService service.RecommendationService, productRepository repository.ProductRepository) RecommendationHandler {
	return RecommendationHandler{
		recommendationService: recommendationService,
		productRepository:     productRepository,
	}
}
