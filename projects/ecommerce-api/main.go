package main

import (
	"encoding/json"
	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"github.com/google/uuid"
	"log"
	"net/http"
)

// Category representa uma categoria
type Category struct {
	ID   string `json:"id"`
	Name string `json:"name"`
}

// Product representa um produto
type Product struct {
	ID          string  `json:"id"`
	Name        string  `json:"name"`
	Description string  `json:"description"`
	Price       float64 `json:"price"`
	Image       string  `json:"image"`
	CategoryID  string  `json:"categoryId"`
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
	// Criar categorias com UUID
	categories := []Category{
		{ID: uuid.NewString(), Name: "Casa"},
		{ID: uuid.NewString(), Name: "Eletrônicos"},
		{ID: uuid.NewString(), Name: "Livros"},
		{ID: uuid.NewString(), Name: "Moda"},
		{ID: uuid.NewString(), Name: "Esportes"},
	}

	// Criar produtos com UUID e associar as categorias
	products := []Product{
		{ID: uuid.NewString(), Name: "Smartphone XYZ", Description: "Um smartphone moderno com todas as funcionalidades avançadas.", Price: 1999.99, Image: "https://via.placeholder.com/200", CategoryID: categories[1].ID},
		{ID: uuid.NewString(), Name: "Notebook ABC", Description: "Notebook potente para profissionais e gamers.", Price: 3499.99, Image: "https://via.placeholder.com/200", CategoryID: categories[3].ID},
		{ID: uuid.NewString(), Name: "TV 4K Ultra HD", Description: "Experiência de cinema em casa com esta TV 4K.", Price: 2599.99, Image: "https://via.placeholder.com/200", CategoryID: categories[0].ID},
		{ID: uuid.NewString(), Name: "Fone de Ouvido Bluetooth", Description: "Áudio de alta qualidade com conectividade sem fio.", Price: 299.99, Image: "https://via.placeholder.com/200", CategoryID: categories[4].ID},
		{ID: uuid.NewString(), Name: "Câmera DSLR", Description: "Capture momentos inesquecíveis com alta resolução.", Price: 4999.99, Image: "https://via.placeholder.com/200", CategoryID: categories[2].ID},
		{ID: uuid.NewString(), Name: "Geladeira Smart", Description: "Geladeira inteligente com controle remoto e economia de energia.", Price: 5999.99, Image: "https://via.placeholder.com/200", CategoryID: categories[0].ID},
		{ID: uuid.NewString(), Name: "Máquina de Lavar", Description: "Alta eficiência e baixo consumo de água e energia.", Price: 2499.99, Image: "https://via.placeholder.com/200", CategoryID: categories[0].ID},
		{ID: uuid.NewString(), Name: "Monitor 4K", Description: "Qualidade de imagem superior para trabalho e entretenimento.", Price: 1299.99, Image: "https://via.placeholder.com/200", CategoryID: categories[1].ID},
		{ID: uuid.NewString(), Name: "Teclado Mecânico", Description: "Teclado gamer com iluminação RGB e alta durabilidade.", Price: 499.99, Image: "https://via.placeholder.com/200", CategoryID: categories[3].ID},
		{ID: uuid.NewString(), Name: "Mouse Sem Fio", Description: "Design ergonômico com alta precisão.", Price: 199.99, Image: "https://via.placeholder.com/200", CategoryID: categories[3].ID},
		{ID: uuid.NewString(), Name: "Liquidificador Turbo", Description: "Potência ideal para qualquer receita.", Price: 299.99, Image: "https://via.placeholder.com/200", CategoryID: categories[0].ID},
		{ID: uuid.NewString(), Name: "Ar Condicionado Split", Description: "Climatização eficiente e silenciosa.", Price: 3199.99, Image: "https://via.placeholder.com/200", CategoryID: categories[4].ID},
		{ID: uuid.NewString(), Name: "Livro de Romance", Description: "Uma história emocionante para todas as idades.", Price: 49.99, Image: "https://via.placeholder.com/200", CategoryID: categories[2].ID},
		{ID: uuid.NewString(), Name: "Enciclopédia Ilustrada", Description: "Conteúdo educacional para toda a família.", Price: 199.99, Image: "https://via.placeholder.com/200", CategoryID: categories[2].ID},
		{ID: uuid.NewString(), Name: "Kit de Ferramentas", Description: "Completo para todas as suas necessidades de reparo.", Price: 599.99, Image: "https://via.placeholder.com/200", CategoryID: categories[4].ID},
		{ID: uuid.NewString(), Name: "Relógio Inteligente", Description: "Monitore sua saúde e atividades diárias.", Price: 899.99, Image: "https://via.placeholder.com/200", CategoryID: categories[1].ID},
		{ID: uuid.NewString(), Name: "Churrasqueira Elétrica", Description: "Ideal para churrascos rápidos e fáceis.", Price: 899.99, Image: "https://via.placeholder.com/200", CategoryID: categories[0].ID},
		{ID: uuid.NewString(), Name: "Mochila de Viagem", Description: "Design ergonômico e material resistente.", Price: 249.99, Image: "https://via.placeholder.com/200", CategoryID: categories[4].ID},
		{ID: uuid.NewString(), Name: "Drone Profissional", Description: "Capture imagens incríveis com alta precisão.", Price: 7999.99, Image: "https://via.placeholder.com/200", CategoryID: categories[1].ID},
		{ID: uuid.NewString(), Name: "Fritadeira Sem Óleo", Description: "Prepare refeições saudáveis de forma prática.", Price: 599.99, Image: "https://via.placeholder.com/200", CategoryID: categories[0].ID},
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

		for _, product := range products {
			if product.ID == idParam {
				w.Header().Set("Content-Type", "application/json")
				if err := json.NewEncoder(w).Encode(product); err != nil {
					http.Error(w, "Erro ao codificar o produto", http.StatusInternalServerError)
				}
				return
			}
		}

		http.Error(w, "Produto não encontrado", http.StatusNotFound)
	})

	// Rota para listar produtos por categoria
	r.Get("/api/categories/{categoryId}/products", func(w http.ResponseWriter, r *http.Request) {
		categoryIDParam := chi.URLParam(r, "categoryId")

		var filteredProducts []Product
		for _, product := range products {
			if product.CategoryID == categoryIDParam {
				filteredProducts = append(filteredProducts, product)
			}
		}

		w.Header().Set("Content-Type", "application/json")
		if err := json.NewEncoder(w).Encode(filteredProducts); err != nil {
			http.Error(w, "Erro ao codificar os produtos", http.StatusInternalServerError)
			return
		}
	})

	// Porta do servidor
	port := ":8181"
	log.Printf("API rodando em http://localhost%s\n", port)
	log.Fatal(http.ListenAndServe(port, r))
}
