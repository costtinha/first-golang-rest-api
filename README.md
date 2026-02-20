# 🛒 API de Gerenciamento de Usuários em Golang

API REST simples para gerenciamento de usuários (users) com operações CRUD básicas.
Implementa camadas de domínio, serviço, repositório e handlers com Gin, persistência via GORM e PostgreSQL, middlewares para recuperação de pânico e ID de requisição, além de configuração via ambiente e containerização com Docker.
O foco é em arquitetura limpa e escalável para um backend básico, ideal para iniciantes ou como base para projetos maiores.

---

## 📋 Índice

- [Sobre o projeto](#-sobre-o-projeto)
- [Tecnologias](#-tecnologias)
- [Funcionalidades principais](#-funcionalidades-principais)
- [Pré-requisitos](#-pré-requisitos)
- [Instalação e Execução](#-instalação-e-execução)
- [Endpoints da API](#-endpoints-da-api)
- [Exemplos de Requisições](#-exemplos-de-requisições)
- [Estrutura do Projeto](#-estrutura-do-projeto)
- [Decisões & Aprendizados](#-decisões--aprendizados)
- [Documentação Swagger](#-documentação-swagger)
- [Licença](#-licença)

---

## 🚀 Sobre o projeto
Sistema backend para gerenciar usuários com campos como nome, email e ID UUID.
Inclui validação de inputs, paginação offset-based para listagem, integração com PostgreSQL via GORM e roteamento com Gin.
Não inclui autenticação ou cache avançado, mas é extensível para adicionar features como JWT ou Redis.
Perfeito para demonstrar fundamentos de API REST em Go em portfólio ou estudos.

---
## 🛠 Tecnologias

| Tecnologia        | Versão    | Finalidade principal                          |
|-------------------|-----------|-----------------------------------------------|
| Go                | 1.22+     | Linguagem                                     |
| Gin               | 1.x       | Framework web para roteamento e handlers      |
| GORM              | 1.x       | ORM para PostgreSQL com migrações automáticas |
| PostgreSQL        | 16+       | Banco de dados relacional                     |
| Validator         | 10.x      | Validação de structs (ex: email, required)    |
| Zap               | —         | Logging estruturado                           |
| Docker + Compose  | 24+ / 3.x | Containerização (app + pg + redis)            |

---

## ✨ Funcionalidades principais

- CRUD completo para usuários com validação de campos (nome, email único).
- Paginação offset-based para listagem de usuários (com limite padrão de 10 itens).
- Middlewares para recuperação de pânico (recover) e geração de ID de requisição (X-Request-ID).
- Migração automática de schema via GORM.
- Configuração via variáveis de ambiente (.env).
- Endpoint de health check para monitoramento.
- Logging com níveis (debug/info em dev/prod).
---

## 📦 Pré-requisitos

- Go 1.22+ instalado (para build manual).
- Docker e Docker Compose instalados (recomendado para execução).
- PostgreSQL (se não usar Docker).

---

## 🚀 Instalação e Execução

### Com Docker Compose (recomendado)

```bash
git clone https://github.com/costtinha/first-golang-rest-api.git
cd first-golang-rest-api
docker compose up -d --build

## 🚀 Instalação e Execução

# A API estará disponível em http://localhost:8080
```

## Crie um arquivo .env com variáveis (exemplo baseado em config.go):
```text
APP_ENV=dev
APP_PORT=8080
DB_HOST=postgres
DB_PORT=5432
DB_USER=postgres
DB_PASS=postgres
DB_NAME=appdb
DB_SSLMODE=disable
LOG_LEVEL=info
```
## Inicie os containers
```bash
docker compose up -d --build
# A API estará disponível em http://localhost:8080.
# O banco PostgreSQL roda em localhost:5433 (mapeado de 5432).
```
## Pare os containers
```bash
docker-compose down
```

### Sem Docker

## Instale as dependências
```bash
go mod tidy
```

## Configure o banco PostgreSQL localmente e atualize o .env ou variáveis de ambiente.

## Execute
```bash
go run main.go
```
---

## 🔗 Endpoints da API

### Públicos (sem autenticação)
## 🔗 Endpoints da API

| Método    | Endpoint              | Descrição                              |
|-----------|-----------------------|----------------------------------------|
| `GET`     | `/health`             | Health check da API                    |
| `POST`    | `/v1/users`           | Criar usuário                          |
| `GET`     | `/v1/users/:id`       | Buscar usuário por ID                  |
| `GET`     | `/v1/users`           | Listar usuários (paginado)             |
| `PUT`     | `/v1/users/:id`       | Atualizar usuário (parcial)            |
| `DELETE`  | `/v1/users/:id`       | Deletar usuário                        |

**Paginação (GET /v1/users)**:  
`?page=1&size=10` (padrão: page=1, size=10)
---

## 📝 Exemplos de Requisições

### Criar usuário

```bash
curl -X POST http://localhost:8080/v1/users \
  -H "Content-Type: application/json" \
  -d '{
    "name": "João Silva",
    "email": "joao@example.com"
  }'
```

**Resposta (201 Created):**
```json
{
  "ID": "uuid-gerado",
  "Name": "João Silva",
  "Email": "joao@example.com",
  "CreatedAt": "timestamp",
  "UpdatedAt": "timestamp"
}
```
### Listar usuários (Paginados)

```bash
 curl "http://localhost:8080/v1/users?page=1&size=10"
```

**Resposta (200 OK):**
```json
{
  "items": [ ... array de usuários ... ],
  "total": 50,
  "page": 1,
  "size": 10
}
```
### Buscar por id

```bash
curl "http://localhost:8080/v1/users/uuid-aqui"
```

**Resposta (200 OK):**
```json
{
  "ID": "uuid-aqui",
  "Name": "João Silva",
  "Email": "joao@example.com",
  "CreatedAt": "timestamp",
  "UpdatedAt": "timestamp"
}
```
---

## 📁 Estrutura do Projeto

```
.
├── internal/
│   ├── config/           # Carregamento de env e DSN
│   ├── http/             # Router com Gin e middlewares
│   ├── logger/           # Logger com Zap
│   ├── platform/database/# Conexão com PostgreSQL
│   └── user/             # Domínio, service, repository, handler
├── main.go               # Entrypoint com injeção de dependências
├── docker-compose.yml    # Containers para app e postgres
├── go.mod / go.sum       # Dependências Go
└── .env                  # Variáveis de ambiente (exemplo)
```


---

## 📚 Decisões e Aprendizados

###  Este projeto foi desenvolvido como estudo prático dos seguintes conceitos em Go:

- Arquitetura em Camadas: Separação clara entre domínio (structs e inputs), serviço (lógica de negócio com validação), repositório (GORM queries) e handlers (Gin).
- Validação: Uso de go-playground/validator para structs tagged.
- Paginação Offset: Simples e funcional, mas com potencial para upgrade para cursor-based em escalas maiores.
- Middlewares: Recover para evitar crashes e RequestID para tracing.
- Configuração: Env-based com defaults, facilitando dev/prod.
- Migração: AutoMigrate do GORM para schema simples.
- Erros Padronizados: Retorno de HTTP status e mensagens claras.
- Containerização: Docker Compose para ambiente reproduzível.

### Aprendizados: Go favorece simplicidade; GORM acelera ORM, mas queries customizadas são essenciais para performance. Próximos passos: adicionar autenticação JWT e cache Redis.

---

## 📄 Licença

Este projeto é de uso educacional e está disponível sob a licença [MIT](LICENSE).

---

<p align="center">
  Desenvolvido por <a href="https://github.com/costtinha">Daniel Costa</a>
</p>
