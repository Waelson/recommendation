package main

import (
	"encoding/csv"
	"fmt"
	"os"
	"strings"
)

func main() {
	// Abrir o arquivo CSV
	file, err := os.Open("/Users/waelson/Documents/workspace/recommendation/machine-learning/dataset/products.csv")
	if err != nil {
		fmt.Printf("Erro ao abrir o arquivo: %v\n", err)
		return
	}
	defer file.Close()

	// Criar o arquivo category.sql
	sqlFile, err := os.Create("category.sql")
	if err != nil {
		fmt.Printf("Erro ao criar o arquivo category.sql: %v\n", err)
		return
	}
	defer sqlFile.Close()

	// Escrever instrução de criação da tabela
	tableCreation := `CREATE TABLE category (
	name VARCHAR(255) PRIMARY KEY
);

`
	sqlFile.WriteString(tableCreation)

	// Ler o conteúdo do CSV
	reader := csv.NewReader(file)
	reader.Comma = ','
	rows, err := reader.ReadAll()
	if err != nil {
		fmt.Printf("Erro ao ler o arquivo CSV: %v\n", err)
		return
	}

	// Verificar se há linhas no arquivo
	if len(rows) < 2 {
		fmt.Println("Arquivo CSV está vazio ou contém apenas cabeçalhos.")
		return
	}

	// Obter o índice da coluna "Category"
	headers := rows[0]
	categoryIndex := -1
	for i, header := range headers {
		if strings.ToLower(header) == "category" {
			categoryIndex = i
			break
		}
	}

	if categoryIndex == -1 {
		fmt.Println("Coluna 'Category' não encontrada no arquivo CSV.")
		return
	}

	// Extrair categorias únicas
	categories := make(map[string]struct{})
	for _, row := range rows[1:] {
		if len(row) > categoryIndex {
			category := strings.TrimSpace(row[categoryIndex])
			categories[category] = struct{}{}
		}
	}

	// Gerar instruções de INSERT
	for category := range categories {
		insert := fmt.Sprintf("INSERT INTO category (name) VALUES ('%s');\n", escapeSingleQuotes(category))
		_, err := sqlFile.WriteString(insert)
		if err != nil {
			fmt.Printf("Erro ao escrever no arquivo category.sql: %v\n", err)
			return
		}
	}

	fmt.Println("Arquivo category.sql gerado com sucesso!")
}

// Função para escapar aspas simples em strings
func escapeSingleQuotes(input string) string {
	return strings.ReplaceAll(input, "'", "''")
}
