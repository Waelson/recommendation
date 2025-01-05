package main

import (
	"github.com/Waelson/recommendation/ecommerce-api/internal/db"
	"github.com/Waelson/recommendation/ecommerce-api/internal/handler"
	"github.com/Waelson/recommendation/ecommerce-api/internal/repository"
	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	_ "github.com/lib/pq"
	"log"
	"net/http"
	"os"
	"strconv"
)

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
	// Configurações do banco de dados obtidas das variáveis de ambiente
	dbConfig := db.Config{
		Host:     getEnv("POSTGRES_HOST", "localhost"),
		Port:     getEnvAsInt("POSTGRES_PORT", 5432),
		User:     getEnv("POSTGRES_USER", "postgres"),
		Password: getEnv("POSTGRES_PASSWORD", "password"),
		DBName:   getEnv("POSTGRES_DB", "product_db"),
	}

	// Conexão com o banco de dados
	conn, err := db.Connect(dbConfig)
	if err != nil {
		log.Fatalf("Failed to connect to database: %v", err)
	}
	defer conn.Close()

	categoryRepository := repository.NewCategoryRepository(conn)
	productRepository := repository.NewProductRepository(conn)

	categoryHandler := handler.NewCategoryHandler(categoryRepository)
	productHandler := handler.NewProductHandler(productRepository)

	r := chi.NewRouter()

	// Middlewares
	r.Use(middleware.Logger)
	r.Use(middleware.Recoverer)
	r.Use(corsMiddleware)

	// Rota para obter as categorias
	r.Get("/api/categories", categoryHandler.ListAll)

	// Rota para obter os detalhes de um produto pelo ID
	r.Get("/api/products/{id}", productHandler.FindByID)

	// Rota para listar produtos por categoria
	r.Get("/api/categories/{categoryId}/products", productHandler.FindByCategoryID)

	// Porta do servidor
	port := ":8181"
	log.Printf("API rodando em http://localhost%s\n", port)
	log.Fatal(http.ListenAndServe(port, r))
}

// getEnv retorna o valor da variável de ambiente ou um valor padrão
func getEnv(key, defaultValue string) string {
	if value, exists := os.LookupEnv(key); exists {
		return value
	}
	return defaultValue
}

// getEnvAsInt retorna o valor da variável de ambiente como int ou um valor padrão
func getEnvAsInt(key string, defaultValue int) int {
	if value, exists := os.LookupEnv(key); exists {
		if intValue, err := strconv.Atoi(value); err == nil {
			return intValue
		}
	}
	return defaultValue
}
