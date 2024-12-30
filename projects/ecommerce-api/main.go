package main

import (
	"encoding/json"
	"log"
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
)

// Category representa uma categoria
type Category struct {
	ID   int    `json:"id"`
	Name string `json:"name"`
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
	r := chi.NewRouter()

	// Middlewares
	r.Use(middleware.Logger)
	r.Use(middleware.Recoverer)
	r.Use(corsMiddleware)

	// Rota para obter as categorias
	r.Get("/api/categories", func(w http.ResponseWriter, r *http.Request) {
		categories := []Category{
			{ID: 1, Name: "Casa"},
			{ID: 2, Name: "Eletrônicos"},
			{ID: 3, Name: "Livros"},
			{ID: 4, Name: "Moda"},
			{ID: 5, Name: "Esportes"},
		}

		w.Header().Set("Content-Type", "application/json")
		if err := json.NewEncoder(w).Encode(categories); err != nil {
			http.Error(w, "Erro ao codificar as categorias", http.StatusInternalServerError)
			return
		}
	})

	// Porta do servidor
	port := ":8181"
	log.Printf("API rodando em http://localhost%s\n", port)
	log.Fatal(http.ListenAndServe(port, r))
}
