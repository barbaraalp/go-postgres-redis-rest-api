# Go Docker PostgreSQL Redis

Projeto **Go Docker PostgreSQL Redis**! 🌟 Este projeto mostra como criar uma aplicação em Go que se conecta a um banco de dados PostgreSQL usando Docker e utiliza o Redis para caching.

### Sumário

1. [Tecnologias](#tecnologias)
2. [Pré-requisitos](#pré-requisitos)
3. [Instruções de Compilação](#instruções-de-compilação)
4. [Estrutura do Projeto](#estrutura-do-projeto)

### Tecnologias:

- Go: Linguagem de programação rápida e eficiente.
- PostgreSQL: Banco de dados relacional confiável e escalável.
- Redis: Sistema de cache de memória in-memory de alta performance.
- Docker: Ferramenta para containerização e orquestração de aplicações.

## Pré-requisitos

Antes de começar, certifique-se de ter o seguinte item instalado em sua máquina:

- [Docker](https://www.docker.com/get-started)

## Instruções de Compilação

Para compilar e executar o projeto, siga estes passos:

1. Clone este repositório em sua máquina local.

2. Navegue até o diretório raiz do projeto.

3. Use o Docker Compose para iniciar o banco de dados PostgreSQL, o cache Redis e a aplicação Go:

```sh
docker-compose up -d
```
4. Para verificar se o caching está funcionando corretamente, verifique os logs do contêiner da aplicação Go:

```sh
docker logs golang
```

5. Acessar a aplicação:

- Para recuperar todas as listagens de alimentos: `GET  http://localhost:8080/food-listings`
- Para recuperar uma listagem de alimentos específica por ID: `GET  http://localhost:8080/food-listings/{id}`

## Estrutura do Projeto

- **Dockerfile**: Define a configuração da imagem Docker para a aplicação Go.
- **docker-compose.yml**: Define os serviços e suas configurações usando o Docker Compose.
- **go.mod**, **go.sum**: Arquivos de configuração do Go Modules especificando dependências do projeto.
- **main.go**: Arquivo principal da aplicação Go.
- **router/router.go**: Define os roteamentos para os diferentes endpoints da API.
- **handler/foodlisting/food_listing_handler.go**: Contém os manipuladores para os endpoints relacionados a listagem de alimentos.
- **model/foodlisting/food_listing.go**: Define a estrutura de dados para um item alimentício.
- **repository/foodlisting/food_listing_repository.go**: Contém funções para acessar o banco de dados PostgreSQL.
- **config/postgres_init.sql**: Script SQL para inicialização do banco de dados PostgreSQL com dados de exemplo.
🚀
