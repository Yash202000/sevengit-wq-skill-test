# 📄 PDF Report Generation Microservice

This is a standalone Go microservice that generates student report PDFs by consuming the Node.js backend API.

---

## 🚀 Features

- Connects to the existing `/api/v1/students/:id` endpoint in the Node.js backend
- Authenticates using login credentials and handles CSRF + cookie-based sessions
- Generates downloadable PDF reports using [`johnfercher/maroto`](https://github.com/johnfercher/maroto)
- Exposes a REST API: `GET /api/v1/students/:id/report`

---

## 🛠️ Requirements

- Go 1.18+
- Node.js backend running on port `5007`
- PostgreSQL database seeded and running
- `.env` file in the root

---

## 🧪 Installation

```bash
# Install dependencies
go mod tidy

# Run the service
go run main.go
```
---

## 📦 .env Format
Make a .env file based on the .env.example:
```
API_BASE_URL=http://localhost:5007/api/v1
ADMIN_USERNAME=admin@school-admin.com
ADMIN_PASSWORD=3OU4zn3q6Zh9

PDF_SERVICE_PORT=8081
```
---
## 🔥 API Endpoint
Generate PDF
```
GET /api/v1/students/:id/report
```