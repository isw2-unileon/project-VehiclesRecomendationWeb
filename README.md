# Vehicles Recommendation Web

A full-stack web application with a smart vehicle recommendation engine powered by Artificial Intelligence (Groq Llama 3.1), an advanced manual search engine. 

## Technologies Used
- **Backend:** Go 1.24 (Golang)
- **Architecture:** Strict Hexagonal Architecture & SOLID principles
- **Database:** PostgreSQL (Containerized with Docker)
- **Frontend:** Vanilla JS + HTML5 + TailwindCSS (Served via Go)
- **AI Integration:** Groq API (Llama 3.1 8B Instant)
- **Security:** JWT (JSON Web Tokens) & Bcrypt for password hashing
- **Concurrency:** Go Routines (Background Price Simulator)

## Architecture Design 
This project strictly follows **Hexagonal Architecture** to decouple the business logic from external frameworks:
- `cmd/api`: Application entry point and server configuration.
- `internal/core/domain`: Core business models (`Car`, `User`).
- `internal/core/ports`: Input/Output interfaces.
- `internal/core/services`: Business logic, Auth, and AI prompt engineering.
- `internal/adapters`: External implementations (HTTP Handlers, Groq Client, PostgreSQL Repositories, Background Simulator).

## How to run locally
1. Clone the repository: `git clone https://github.com/isw2-unileon/project-VehiclesRecomendationWeb.git`
2. Run the backend: `go run cmd/api/main.go`


## How to Run Locally

### Prerequisites
- Go 1.24+ installed
- Docker & Docker Compose installed
- A free API Key from [Groq Cloud](https://console.groq.com/)

### Commands
1. **Clone the repository:**
   `git clone https://github.com/isw2-unileon/project-VehiclesRecomendationWeb.git`
2. **Download dependencies:**
   `go mod tidy`
3. **Compile the project:**
   `go build -v ./...`
4. **Run the server:**
   `go run cmd/api/main.go`

   http://localhost:8080/api/health

## Database Seeder (CSV to SQL)
This project includes a custom Go script to automate the database seeding process. The script reads the raw data from our `resources/Cars Datasets 2025.csv` file, processes the columns, and generates ready-to-use SQL `INSERT` statements.

### Features
* **Data Cleaning:** Automatically removes unnecessary text characters, currency symbols (`$`), and units (`cc`, `hp`, `km/h`) from numeric fields.
* **Range Parsing:** Detects numerical ranges (e.g., `150-200`) and dynamically generates a random integer within that range to ensure valid SQL numeric types.
* **SQL Injection Prevention:** Escapes single quotes in text fields to prevent syntax errors during insertion.

**Usage:**
`go run cmd/seeder/mainCSV.go > coches_seeder.sql`

## 📌 Roadmap & Next Steps

- **Authentication:** Upgrade the current login system to integrate third-party secure authentication services (e.g., OAuth, Auth0) as per project requirements.
- **Real-time Market Simulation:** Maintain and expand the background API simulator that automatically updates vehicle data and prices periodically using Go Routines.
- **Version Control & CI/CD:** Review pending code changes and accept the open Pull Requests into the main branch to proceed with the frontend integration.
