package main

import (
	"log"
	"net/http"
)

func main() {
	// Define o diretório onde estão as páginas estáticas
	staticDir := "./static"

	// Servir arquivos estáticos
	http.Handle("/", http.FileServer(http.Dir(staticDir)))

	// Porta do servidor
	port := ":8080"

	log.Printf("Servidor rodando em http://localhost%s\n", port)
	log.Fatal(http.ListenAndServe(port, nil))
}
