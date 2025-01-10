# Sistema de Recomendação de Produtos

Este repositório contém a implementação de um sistema de recomendação de produtos utilizando Machine Learning. O projeto foi desenvolvido para demonstrar como integrar sistemas transacionais com modelos de machine learning que realizam inferências em tempo real, oferecendo soluções práticas e eficientes para recomendação de produtos em e-commerce.

O sistema combina uma estrutura modular que separa as responsabilidades de frontend, backend e machine learning, que permite explorar a interação entre diferentes componentes. Ele também exemplifica como conectar APIs ao modelo treinado.

Com este repositório, espera-se fornecer um exemplo claro e funcional de como aplicar técnicas de machine learning para melhorar a experiência do usuário e aumentar a eficiência de negócios baseados em dados.

## Arquitetura
![Architecture](documentation/architecture.png)

### Componentes da Arquitetura
| Componente                | Descrição                                                                                     |
|---------------------------|---------------------------------------------------------------------------------------------|
| **Front-End**             | Responsável por exibir os produtos, categorias e recomendações para os usuários.            |
| **Item API Cluster**      | API que gerencia os dados de produtos (leitura e escrita) no banco de dados.                |
| **Recommendation API**    | API que utiliza o modelo de machine learning para fornecer produtos recomendados.            |

## Requisitos para Execução

### Pré-requisitos
- [Go](https://golang.org/) 1.19 ou superior
- [PostgreSQL](https://www.postgresql.org/) instalado e configurado
- Navegador para testar o frontend

### Configuração do Banco de Dados
1. Inicialize o banco de dados usando o script `init.sql`.
   ```bash
   psql -U <seu_usuario> -d <seu_banco> -f init.sql
   ```
2. Insira as categorias usando o script `category.sql`.
   ```bash
   psql -U <seu_usuario> -d <seu_banco> -f category.sql
   ```
3. Insira os produtos manualmente ou carregue-os via script CSV.

### Configuração do Backend
1. Navegue até o diretório do projeto desejado (ex.: `product-api` ou `recommendation-api`).
2. Execute o servidor backend:
   ```bash
   go run main.go
   ```
3. O servidor estará disponível em `http://localhost:8181` (ou porta configurada).

### Configuração do Frontend
1. Navegue até o diretório `frontend`.
2. Abra o arquivo `index.html` no navegador.
3. A navegação e as funcionalidades devem estar disponíveis.

## Endpoints Disponíveis

### Requisições Principais

#### Produtos
- **`GET /api/products/:id`**: Retorna detalhes de um produto pelo ID.
- **`GET /api/products/:id/:n/recommend`**: Retorna `n` recomendações para o produto especificado.

#### Categorias
- **`GET /api/categories`**: Retorna a lista de categorias em ordem alfabética.

## Modelo de Machine Learning

### O que é Machine Learning (ML)?
Machine Learning é um campo da inteligência artificial que utiliza algoritmos para aprender padrões a partir de dados. Com base nesses padrões, os modelos podem fazer previsões ou tomar decisões sem serem explicitamente programados para cada tarefa específica.

### O que é um Sistema de Recomendação?
Um sistema de recomendação é uma aplicação que sugere itens relevantes para os usuários com base em seus interesses, comportamentos anteriores ou características de outros itens similares. No contexto de e-commerce, isso pode incluir produtos relacionados, populares ou personalizados.

### Vantagens de um Sistema de Recomendação para E-commerce
- **Aumento nas vendas**: Oferece recomendações personalizadas que incentivam compras adicionais.
- **Melhor experiência do usuário**: Os clientes encontram produtos relevantes de forma rápida e eficiente.
- **Retenção de clientes**: Torna a navegação mais agradável, promovendo o retorno dos usuários.
- **Insights valiosos**: Oferece dados sobre preferências e tendências de comportamento.

### Técnica Utilizada no Treinamento do Modelo
O modelo de recomendação implementado utiliza a abordagem de **Filtragem Baseada em Conteúdo**. Essa técnica analisa as características dos produtos, como categoria, preço, descrição e outras informações textuais, para encontrar itens semelhantes.

#### Etapas do Processo:
1. **Representação dos Dados**: Os dados do produto são transformados em vetores numéricos, utilizando técnicas como TF-IDF para textos.
2. **Cálculo de Similaridade**: É utilizada a similaridade de cosseno para medir a proximidade entre produtos.
3. **Geração de Recomendações**: Para um produto selecionado, os produtos mais similares são ordenados e retornados como recomendação.

Essa abordagem é eficiente para sistemas onde os dados de interação do usuário são limitados, mas as informações sobre os produtos são ricas e bem estruturadas.

2. Teste a navegação para categorias, detalhes do produto e recomendações.




