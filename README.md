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