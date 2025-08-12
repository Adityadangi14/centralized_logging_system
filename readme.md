


# 📝 Centralized Logging System using Golang Microservices & Docker

A **Centralized Logging System** built with **Golang microservices** and containerized using **Docker**.
This system simulates multiple log-producing microservices that send logs to a **Log Collector** service, which stores them in a central PostgreSQL database.
A **user-facing API** is provided for querying logs.

---

## 📂 Project Structure


```
.
├── api
│   ├── db
│   │   └── gen
│   │       ├── db.go
│   │       ├── models.go
│   │       └── query.sql.go
│   ├── Dockerfile
│   ├── go.mod
│   ├── go.sum
│   ├── handler
│   │   └── http_handler.go
│   ├── initilizers
│   │   ├── conn_db.go
│   │   └── load_env.go
│   ├── main.go
│   ├── query.sql
│   ├── schema.sql
│   └── sqlc.yml
├── docker-compose.yml
├── linux_sys_logs
│   ├── Dockerfile
│   ├── go.mod
│   ├── main.go
│   └── src
├── log_collector
│   ├── db
│   │   └── gen
│   │       ├── db.go
│   │       ├── models.go
│   │       └── query.sql.go
│   ├── Dockerfile
│   ├── go.mod
│   ├── go.sum
│   ├── main.go
│   ├── models
│   │   └── log_models.go
│   ├── query.sql
│   ├── schema.sql
│   ├── sqlc.yml
│   └── src
│       └── initializers
│           ├── db_conn.go
│           ├── load_env.go
│           └── tcp_conn.go
├── login_audit
│   ├── Dockerfile
│   ├── go.mod
│   ├── main.go
│   └── src
└── readme.md


```

---

## 🛠 Features

- **Multiple Microservices**
  - `linux_sys_logs` → Generates system logs
  - `login_audit` → Generates login audit logs
  - `log_collector` → Receives logs via TCP and stores them in DB
  - `api` → HTTP API to query logs

- **Tech Stack**
  - **Golang** for microservices
  - **PostgreSQL** for centralized storage
  - **SQLC** for type-safe DB queries
  - **Docker & Docker Compose** for containerization

- **Communication Flow**
  - Microservices → **TCP** → Log Collector → **PostgreSQL**
  - API → **PostgreSQL**

---

## 📦 Installation & Setup

### 1️⃣ Clone the Repository
```bash
git clone https://github.com/<your-username>/centralized_logging_system.git
cd centralized_logging_system
````

### 2️⃣ Create `.env` Files

Each service that loads environment variables must have a `.env` file.

Example `.env` for **log\_collector** and **api**:

```
DB_HOST=postgres
DB_PORT=5432
DB_USER=postgres
DB_PASSWORD=postgres
DB_NAME=logs_db
```

### 3️⃣ Build & Run with Docker Compose

```bash
docker-compose up --build
```

---

## 📡 API Endpoints

**Base URL:** `http://localhost:4000`

*Examples*
GET http://localhost:4000/logs?service=syslog
GET http://localhost:4000/logs?username=admin&is.blacklisted=false
GET http://localhost:4000/logs?level=error
GET http://localhost:4000/logs?service=syslog


---

## ⚙ How It Works

1. **Log Generators** (`linux_sys_logs`, `login_audit`)
   Send log entries via **TCP** to the **log\_collector**.

2. **Log Collector** (`log_collector`)

   * Receives TCP messages
   * Parses them
   * Stores them in **PostgreSQL**

3. **API Service** (`api`)

   * Exposes HTTP endpoints to retrieve and query logs

4. **PostgreSQL Database**

   * Central storage for all logs

---

## 🖼 Architecture Diagram

```
[linux_sys_logs]     [login_audit]
        \                  /
         \                /
           ----TCP----> [log_collector] ---> [PostgreSQL] <--- [api]
```

---

## 🧪 Development Notes

* **SQLC** generates Go code for database queries:

```bash
sqlc generate
```

* Rebuild services after making changes:

```bash
docker-compose up --build
```

---

## 🏆 Assignment Objectives Met

* ✅ **Microservice Communication** via TCP & HTTP
* ✅ **Go Concurrency** for log handling
* ✅ **REST API** for log retrieval
* ✅ **Containerization** with multi-stage Docker builds

---
