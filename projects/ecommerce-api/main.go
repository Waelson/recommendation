package main

import (
	"encoding/json"
	"log"
	"net/http"
	"strconv"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
)

// Category representa uma categoria
type Category struct {
	ID   int    `json:"id"`
	Name string `json:"name"`
}

// Product representa um produto
type Product struct {
	ID          int     `json:"id"`
	Name        string  `json:"name"`
	Description string  `json:"description"`
	Price       float64 `json:"price"`
	Image       string  `json:"image"`
}

// Middleware para suportar CORS
func corsMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Access-Control-Allow-Origin", "*")
		w.Header().Set("Access-Control-Allow-Methods", "GET, POST, OPTIONS")
		w.Header().Set("Access-Control-Allow-Headers", "Content-Type, Authorization")
		if r.Method == http.MethodOptions {
			w.WriteHeader(http.StatusNoContent)
			return
		}
		next.ServeHTTP(w, r)
	})
}

func main() {
	// Dados simulados de produtos
	products := []Product{
		{ID: 1, Name: "Smartphone XYZ", Description: "Um smartphone moderno com todas as funcionalidades avançadas.", Price: 1999.99, Image: "https://via.placeholder.com/300"},
		{ID: 2, Name: "Notebook ABC", Description: "Notebook potente para profissionais e gamers.", Price: 3499.99, Image: "https://via.placeholder.com/300"},
		{ID: 3, Name: "TV 4K Ultra HD", Description: "Experiência de cinema em casa com esta TV 4K.", Price: 2599.99, Image: "https://via.placeholder.com/300"},
		{ID: 4, Name: "Fone de Ouvido Bluetooth", Description: "Áudio de alta qualidade com conectividade sem fio.", Price: 299.99, Image: "https://via.placeholder.com/300"},
		{ID: 5, Name: "Câmera DSLR", Description: "Capture momentos inesquecíveis com alta resolução.", Price: 4999.99, Image: "https://via.placeholder.com/300"},
	}

	categories := []Category{
		{ID: 1, Name: "Casa"},
		{ID: 2, Name: "Eletrônicos"},
		{ID: 3, Name: "Livros"},
		{ID: 4, Name: "Moda"},
		{ID: 5, Name: "Esportes"},
	}

	r := chi.NewRouter()

	// Middlewares
	r.Use(middleware.Logger)
	r.Use(middleware.Recoverer)
	r.Use(corsMiddleware)

	// Rota para obter as categorias
	r.Get("/api/categories", func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		if err := json.NewEncoder(w).Encode(categories); err != nil {
			http.Error(w, "Erro ao codificar as categorias", http.StatusInternalServerError)
			return
		}
	})

	// Rota para obter os detalhes de um produto pelo ID
	r.Get("/api/products/{id}", func(w http.ResponseWriter, r *http.Request) {
		idParam := chi.URLParam(r, "id")
		id, err := strconv.Atoi(idParam)
		if err != nil {
			http.Error(w, "ID inválido", http.StatusBadRequest)
			return
		}

		for _, product := range products {
			if product.ID == id {
				w.Header().Set("Content-Type", "application/json")
				if err := json.NewEncoder(w).Encode(product); err != nil {
					http.Error(w, "Erro ao codificar o produto", http.StatusInternalServerError)
				}
				return
			}
		}

		http.Error(w, "Produto não encontrado", http.StatusNotFound)
	})

	// Porta do servidor
	port := ":8181"
	log.Printf("API rodando em http://localhost%s\n", port)
	log.Fatal(http.ListenAndServe(port, r))
}
