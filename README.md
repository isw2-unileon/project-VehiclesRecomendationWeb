# Vehicles Recommendation Web
 
A full-stack web application with a smart vehicle recommendation engine powered by Artificial Intelligence (Groq Llama 3.1), an advanced manual search engine, full-text search, and a user favorites system with JWT-based authentication.
 
**Live:** [https://project-vehiclesrecomendationweb.onrender.com]
 
---
 
## Description
 
Vehicles Recommendation Web allows users to register, log in, and search for vehicles from a real dataset of 2025 cars. Users can get AI-powered recommendations by describing their needs in natural language, search the database by brand, fuel type, price and seats, perform full-text model searches, and save vehicles to a personal favorites list. For every car, users can view photos, details, and save it to their favorites. 
 
---
 
## Prerequisites
 
- [Go 1.24+](https://golang.org/dl/)
- [Docker & Docker Compose](https://docs.docker.com/get-docker/)
- A free API key from [Groq Cloud](https://console.groq.com/)
---
 
## Setup and Run Locally
 
### 1. Clone the repository
 
```bash
git clone https://github.com/isw2-unileon/project-VehiclesRecomendationWeb.git
cd project-VehiclesRecomendationWeb
```
 
### 2. Configure environment variables
 
Create a `.env` file in the project root. This file is ignored by Git for security.
 
```env
GROQ_API_KEY=gsk_your_key_here
JWT_SECRET=your_secret_key_here
```
 
### 3. Download Go dependencies
 
```bash
go mod tidy
```
 
### 4. Start the database
 
```bash
docker compose up -d
```
 
### 5. Load schema and seed data
 
```bash
# Create the database tables
docker exec -i vehicles_db psql -U postgres -d cars < resources/schema.sql
 
# Convert encoding and seed the database with car data
iconv -f UTF-16 -t UTF-8 cars_seeder.sql > cars_seeder_utf8.sql
docker exec -i vehicles_db psql -U postgres -d cars < cars_seeder_utf8.sql
```
 
### 6. Run the server
 
```bash
go run cmd/api/main.go
```
 
The application will be available at [http://localhost:8080](http://localhost:8080).  
Health check: [http://localhost:8080/api/health](http://localhost:8080/api/health)
 
---
 
## Daily Development
 
```bash
# Start the database (if not running)
docker compose up -d
 
# Run the server
go run cmd/api/main.go
 
# Stop the server
Ctrl+C
 
# Stop the database (data is preserved)
docker compose stop
```
 
---
 
## Running Tests
 
```bash
# Run all tests
go test ./...
 
# Run tests with coverage
go test -cover ./...
 
# Run tests for a specific package
go test ./internal/core/services/...
go test ./internal/adapters/handlers/...
```
 
---
 
## How to Contribute
 
This project follows **Trunk Based Development (TBD)**. Each task is developed in a short-lived branch created from `main`, integrated via Pull Request, and deleted after merging. The `main` branch always contains the latest stable and deployed version.
 
### Branch naming
 
```
feature/short-description
fix/short-description
refactor/short-description
docs/short-description
test/short-description
```
 
### Commit convention
 
```
feat: add new feature
fix: correct a bug
refactor: improve code structure
test: add or update tests
docs: update documentation
chore: maintenance tasks
```
 
### Pull Request process
 
1. Create a short-lived branch from `main`:
   ```bash
   git checkout -b feature/your-feature
   ```
2. Make changes and commit following the convention above.
3. Push and open a Pull Request to `main`.
4. Request a review from a team member.
5. Once approved and CI passes, merge the PR and delete the branch immediately:
   ```bash
   git checkout main
   git pull origin main
   git branch -d feature/your-feature
   git push origin --delete feature/your-feature
   ```
 
---
 
## Technical Documentation
 
Full technical documentation (architecture, data models, design decisions, API reference) is available in [`/docs`](./docs).