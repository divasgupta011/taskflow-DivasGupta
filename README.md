# TaskFlow Backend

A minimal task management backend built with Go, PostgreSQL, and Docker.
Supports authentication, project management, and task tracking with proper authorization.

---

## 🚀 Tech Stack

* **Language:** Go
* **Database:** PostgreSQL
* **Auth:** JWT (stateless)
* **Migrations:** golang-migrate
* **Containerization:** Docker + docker-compose

---

## 🧱 Architecture

The application follows a layered architecture:

* **Handler** → HTTP layer (request/response)
* **Service** → Business logic & authorization
* **Repository** → Database access (SQL)

### Key Decisions

* Used **raw SQL + migrations** for control and simplicity
* Implemented **JWT-based stateless authentication**
* Enforced **ownership-based authorization** at service layer
* Automated **migrations on startup** for zero manual setup

---

## ⚙️ Running Locally

```bash
git clone https://github.com/divasgupta011/taskflow-DivasGupta.git
cd taskflow-DivasGupta
cp .env.example .env
docker compose up --build
```

Server runs at:

```text
http://localhost:8080
```

---

## 🔑 Test Credentials (Seed Data)

```text
Email: test@example.com
Password: password123
```

Seed data includes:

* 1 user
* 1 project
* 3 tasks

---

## 📦 API Overview

### Auth

* `POST /auth/register`
* `POST /auth/login`

### Projects

* `GET /projects`
* `POST /projects`
* `GET /projects/{id}`
* `PATCH /projects/{id}`
* `DELETE /projects/{id}`

### Tasks

* `GET /projects/{id}/tasks?status=&assignee=`
* `POST /projects/{id}/tasks`
* `PATCH /tasks/{id}`
* `DELETE /tasks/{id}`

---

## 🧪 API Testing

Import the included Postman collection:

```text
taskflow.postman_collection.json
```

Steps:

1. Login → get token
2. Create project
3. Create tasks
4. Test filters and updates

---

## ✅ Features

* JWT authentication (24h expiry)
* Secure password hashing (bcrypt)
* Ownership-based access control
* Task filtering (status, assignee)
* Automatic DB migrations
* Seed data for instant testing
* Dockerized setup (single command)

---

## 📌 Notes

* All endpoints return **JSON responses**
* Proper HTTP status codes are used (401, 403, 404)
* No sensitive data (e.g., passwords) is exposed

---

## 🏁 Run Summary

```bash
docker compose up
```

→ DB starts
→ migrations run
→ seed data inserted
→ backend ready

No manual steps required.

---
