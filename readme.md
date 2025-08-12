Here’s a fresh, polished `README.md` you can drop into your project:

```markdown
# 📝 Centralized Logging System using Golang Microservices & Docker

A **Centralized Logging System** built with **Golang microservices** and containerized using **Docker**.
This system simulates multiple log-producing microservices that send logs to a **Log Collector** service, which then stores them in a central database.
A **user-facing API** is provided for querying logs.

---

## 📂 Project Structure

```

.
├── api                     # API service to query logs
│   ├── db                  # SQLC generated DB code
│   ├── handler             # HTTP handlers
│   ├── initilizers         # Database & env initialization
│   ├── query.sql           # SQL queries for logs
│   ├── schema.sql          # Database schema
│   ├── sqlc.yml            # SQLC configuration
│   ├── Dockerfile
│   ├── go.mod / go.sum
│   └── main.go
│
├── linux\_sys\_logs          # Simulated Linux system logs generator
│   ├── Dockerfile
│   ├── go.mod
│   └── main.go
│
├── log\_collector           # Central log collector service
│   ├── db                  # SQLC generated DB code
│   ├── models              # Log models
│   ├── src/initializers    # TCP server, DB & env loader
│   ├── query.sql           # SQL queries
│   ├── schema.sql          # Database schema
│   ├── sqlc.yml
│   ├── Dockerfile
│   ├── go.mod / go.sum
│   └── main.go
│
├── login\_audit             # Simulated login audit log generator
│   ├── Dockerfile
│   ├── go.mod
│   └── main.go
│
├── docker-compose.yml      # Docker orchestration file
└── readme.md               # Project documentation

````

---

## 🛠 Features

- **Multiple Microservices**
  - `linux_sys_logs` → Generates system logs
  - `login_audit` → Generates login audit logs
  - `log_collector` → Receives logs from microservices via TCP and stores them in DB
  - `api` → Provides HTTP API to query logs

- **Tech Stack**
  - **Go** for microservices
  - **PostgreSQL** for storage
  - **SQLC** for type-safe database queries
  - **Docker & Docker Compose** for containerization

- **Communication**
  - Microservices → TCP → Log Collector → PostgreSQL
  - API → PostgreSQL

---

## 📦 Installation & Setup

### 1️⃣ Clone the Repository
```bash
git clone https://github.com/<your-username>/centralized_logging_system.git
cd centralized_logging_system
````

### 2️⃣ Create `.env` Files

Each service that loads environment variables should have a `.env` file.
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

**Base URL:** `http://localhost:8080`

| Method | Endpoint      | Description                 |
| ------ | ------------- | --------------------------- |
| GET    | `/logs`       | Fetch all logs              |
| GET    | `/logs/:id`   | Fetch a log by ID           |
| POST   | `/logs/query` | Query logs based on filters |

---

## ⚙ How It Works

1. **Log Generators (`linux_sys_logs`, `login_audit`)**
   Periodically send log entries via TCP to the **log\_collector**.

2. **Log Collector (`log_collector`)**

   * Receives TCP log messages
   * Parses and stores them in **PostgreSQL**

3. **API (`api`)**

   * Provides HTTP endpoints to query and retrieve logs

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

* Rebuild services after making code changes:

```bash
docker-compose up --build
```

---

## 🏆 Assignment Objectives Met

* ✅ **Microservice Communication** using TCP and HTTP
* ✅ **Go Concurrency** in log handling
* ✅ **API Design** with REST endpoints
* ✅ **Containerization** with multi-stage Docker builds

---
