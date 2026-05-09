# project-VehiclesRecomendationWeb
# Vehicles Recommendation Web

Web application with a vehicle recommendation and comparison system using AI (Groq) and a manual search engine.

## Technologies Used
- **Backend:** Golang (Hexagonal Architecture)
- **Database:** SQL (PostgreSQL/MySQL)
- **Frontend:** [React / Next.js / Angular / Vue]
- **AI Integration:** Groq API (LLM for recommendations)

## Architecture
This project follows **Hexagonal Architecture** and **SOLID** principles:
- `cmd/api`: Application entry point.
- `internal/core/domain`: Business models (Car, User).
- `internal/core/ports`: Interfaces.
- `internal/core/services`: Business logic and AI integration.
- `internal/adapters`: Handlers (HTTP) and Repositories (SQL DB).

## How to run locally
1. Clone the repository: `git clone https://github.com/isw2-unileon/project-VehiclesRecomendationWeb.git`
2. Run the backend: `go run cmd/api/main.go`


## How to Run Locally

### Prerequisites
- Go 1.24+
- PostgreSQL database

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



5. **Database Seeder (CSV to SQL)**
This project includes a custom Go script to automate the database seeding process. The script reads the raw data from our `resources/Cars Datasets 2025.csv` file, processes the columns, and generates ready-to-use SQL `INSERT` statements.
### Features
* **Data Cleaning:** Automatically removes unnecessary text characters, currency symbols (`$`), and units (`cc`, `hp`, `km/h`) from numeric fields.
* **Range Parsing:** Detects numerical ranges (e.g., `150-200`) and dynamically generates a random integer within that range to ensure valid SQL numeric types.
* **SQL Injection Prevention:** Escapes single quotes in text fields to prevent syntax errors during insertion.

`go run cmd/seeder/mainCSV.go > coches_seeder.sql`

   


login ( mejor buscar servicios de terceros)

base de datos , simular una especie de api que se vaya actualizado de vez en cuando , 

vemos si puede aceptar el pull request 



