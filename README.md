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

### Executando
A aplicação está configurada para ser executada com Docker Compose. Para iniciá-la, entre no diretório raiz da aplicação e execute o seguinte comando:

```bash
docker-compose up --build
```

Após a inicialização dos serviços pelo Docker Compose, acesse a aplicação através da URL http://localhost:8080 em seu navegador. Depois navegue nas categorias e nos itens para visualizar as recomendações sendo realizadas em tempo real pela API de Recomendação.

![Application in Action](documentation/recommendation_in_action.gif)