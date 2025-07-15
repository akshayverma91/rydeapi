## Criteria to select programming language: I chose Go as language.

## Criteria to select framework: use Gin since fast and minimalistic 

# Ryde API — Golang Backend with JWT, MongoDB, and Docker

## Features

JWT-based Auth (Login/Register)
MongoDB-based user management
Versioned APIs (`/api/v1`, `/api/v2`)
Follows/Followers relationship (via separate collection)
Geo-query: Find users nearby (MongoDB 2dsphere)
Docker and Docker Compose support
Swagger UI
Postman collection for testing

---

## 🧾 Project Structure

config/           # Mongo + env config
controllers/      # All API handlers
middleware/       # JWT auth middleware
models/           # Data models
repositories/     # DB logic (users, follows)
routes/           # Router groups (v1/v2)
utils/            # Helpers, JWT logic
docs/             # Swagger docs
tests/            # Unit/integration tests
main.go           # Entry point
.env              # Local environment vars
.env.example      # Template env file
docker-compose.yml
Dockerfile
go.mod
Readme.md

## How to Run the API
git clone https://github.com/your-username/rydeapi.git
cd rydeapi

## Setup Environment
Create your .env file:

## Run Docker Compose
docker build -t rydeapi .
docker run -p 8080:8080 rydeapi

for docker compose
to build: docker-compose build
to start: docker-compose up
to stop: docker-compose down

## Initialize Swagger
use command to generate swagger using swag init
add or update controller run again swag init
Run project and go to http://localhost:8080/swagger/index.html 

## Run Locally (without Docker)
Start MongoDB
go run main.go


## To run test
go test ./tests