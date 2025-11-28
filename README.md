# Go REST API (Gin + GORM + PostgreSQL + Viper + Zap)

A clean and scalable **Go REST API starter** built using:

- **Gin** – HTTP router & middleware
- **GORM** – ORM for PostgreSQL
- **Viper** – Configuration & environment management
- **Zap** – Structured logging (Uber)
- **Clean Architecture** with `cmd/`, `internal/`, `pkg/`

This repository provides a production-ready base for building APIs with proper structure, configuration management, logging, services, and repository layer.

---

### 🧩 What this layout means

| Folder                  | Purpose                                 |
| ----------------------- | --------------------------------------- |
| `cmd/api`               | Application entry point                 |
| `internal/config`       | Loads environment variables using Viper |
| `internal/models`       | GORM models                             |
| `internal/repositories` | DB connection + repository layer        |
| `internal/services`     | Business logic                          |
| `internal/handlers`     | HTTP handlers + routing                 |
| `pkg/utils`             | Shared utilities like response helpers  |

---

## ⚙️ Environment Configuration

Create a `.env` file from .env.example

These values are loaded through `config.go` using **Viper**.

---

## 🛠️ Running the Project

go run ./cmd/api
