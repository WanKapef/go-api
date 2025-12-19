# Go API – REST Backend com Gorilla Mux e SQLite

API REST escrita em **Go**, utilizando **Gorilla Mux** como router, **SQLite** como banco de dados e uma arquitetura em camadas inspirada em aplicações backend modernas (Handler → Service → Repository).

O projeto foi estruturado para ser **simples, explícito, testável e sem frameworks mágicos**, seguindo boas práticas da comunidade Go.

---

## 📦 Tecnologias

* **Go** 1.22+
* **Gorilla Mux** – HTTP Router
* **SQLite** – Banco de dados embutido
* **golang-migrate** – Migrations
* **database/sql** – Acesso ao banco
* **Homebrew** (macOS) – para ferramentas auxiliares

---

## 🧱 Arquitetura

O projeto segue uma separação clara de responsabilidades:

```
api/
├── cmd/
│   └── api/
│       └── main.go          # Bootstrap da aplicação
├── internal/
│   ├── config/              # Configurações (env, porta, db path)
│   ├── database/            # Conexão com SQLite
│   ├── handler/             # Camada HTTP (req/res)
│   ├── service/             # Regras de negócio
│   ├── repository/          # Acesso ao banco
│   ├── model/               # Estruturas de domínio
│   ├── middleware/          # Logger, middlewares HTTP
│   └── httpx/               # Helpers HTTP (ex: extrair ID da URL)
├── migrations/              # SQL migrations
├── data/                    # Arquivo SQLite (.db)
├── go.mod
└── go.sum
```

---

## 🔄 Fluxo de uma requisição

```
HTTP Request
   ↓
Middleware (logger, etc)
   ↓
Handler (HTTP)
   ↓
Service (negócio)
   ↓
Repository (SQL)
   ↓
SQLite
```

---

## ⚙️ Configuração

### Variáveis de ambiente

Por padrão o projeto usa valores locais, mas pode ser adaptado para `.env`.

Exemplo:

```env
PORT=8080
DATABASE_PATH=./data/app.db
```

---

## ▶️ Executando o projeto

### 1️⃣ Instalar dependências

```bash
go mod tidy
```

### 2️⃣ Executar migrations

Instale o migrate (caso não tenha):

```bash
brew install golang-migrate
```

Rodar migrations:

```bash
migrate -database sqlite3://data/app.db -path migrations up
```

---

### 3️⃣ Subir a API

```bash
go run ./cmd/api
```

A API ficará disponível em:

```
http://localhost:8080
```

---

## 📌 Rotas disponíveis

### ➕ Criar usuário

```http
POST /users
Content-Type: application/json
```

```json
{
  "name": "Wan",
  "email": "wan@email.com"
}
```

---

### 📄 Listar usuários

```http
GET /users
```

---

### 🔎 Buscar usuário por ID

```http
GET /users/{id}
```

Exemplo:

```bash
curl http://localhost:8080/users/1
```

Resposta:

```json
{
  "id": 1,
  "name": "Wan",
  "email": "wan@email.com"
}
```

---

### ✏️ Atualizar usuário

```http
PUT /users
Content-Type: application/json
```

```json
{
  "id": 1,
  "name": "Wan Atualizado",
  "email": "wan@email.com"
}
```

---

### ❌ Deletar usuário

```http
DELETE /users/{id}
```

Exemplo:

```bash
curl -X DELETE http://localhost:8080/users/1
```

---

### Onde ficam as rotas?

As rotas são registradas diretamente no `main.go`, mantendo:

* Router isolado
* Handlers sem dependência de framework

```go
router.HandleFunc("/users/{id}", httpx.WithID(userHandler.GetByID)).Methods("GET")
router.HandleFunc("/users/{id}", httpx.WithID(userHandler.Delete)).Methods("DELETE")
```

---

### Como funciona o `WithID`?

O helper `WithID` extrai automaticamente o `id` da URL e injeta no handler:

```go
func (h *UserHandler) Delete(w http.ResponseWriter, r *http.Request, id int64)
```

Isso mantém os handlers **limpos e explícitos**.

---

## 📜 Logger

Middleware HTTP customizado, inspirado no Gin:

* Status colorido
* Método HTTP destacado
* Path
* Tempo de resposta

Exemplo de log:

```
[GIN] 200 | 1.23ms | GET    /users/1
```

---

## 🧪 Próximos passos sugeridos

* [✅] GET `/users/{id}`
* [🧑‍💻] Paginação (`limit`, `offset`)
* [ ] Seeds para popular BD
* [ ] Filtros de busca
* [ ] Middleware de erro padronizado em JSON
* [ ] Testes HTTP (`httptest`)
* [ ] Request ID
* [ ] Autenticação JWT
* [ ] Swagger / OpenAPI

---

## 👤 Autor

**WanKapef**
Projeto de estudo e base para APIs REST em Go.

