Developer Commands Reference

Prerequisites

Go 1.24+
Docker & Docker Compose

-------------------------------------------------------------------------------------
Initial Setup, first time only !!

# 1. Clone the repository
git clone https://github.com/isw2-unileon/project-VehiclesRecomendationWeb.git
cd project-VehiclesRecomendationWeb

Create a `.env` file in the root folder of the project. This file is ignored by Git for security.
Inside the file, paste your personal Groq API Key (get it for free at console.groq.com):
GROQ_API_KEY=gsk_your_key_here

# 2. Download Go dependencies
go mod tidy

# 3. Start the database
docker compose up -d

# 4. Create the cars table
docker exec -i vehicles_db psql -U postgres -d cars < resources/schema.sql

# 5. Seed the database with car data
iconv -f UTF-16 -t UTF-8 cars_seeder.sql > cars_seeder_utf8.sql
docker exec -i vehicles_db psql -U postgres -d cars < cars_seeder_utf8.sql

# 6. Run the server
go run cmd/api/main.go

-------------------------------------------------------------------------------------

Daily

# Start the database (if not running)
docker compose up -d

# Run the server
go run cmd/api/main.go

# Stop the server
Ctrl+C

# Stop the database (keeps data)
docker compose stop

# Stop and remove the database container (data is safe in Docker volume)
docker compose down

-------------------------------------------------------------------------------------

Build 

# Compile the project (checks for errors)
go build -v ./...

-------------------------------------------------------------------------------------

Test

# Run all tests
go test ./...

# Run tests with coverage
go test -cover ./...

# Run tests for a specific package
go test ./internal/core/services/...

-------------------------------------------------------------------------------------

# Start the database container
docker compose up -d

# Connect to the database manually
docker exec -it vehicles_db psql -U postgres -d cars

# Run a SQL file
docker exec -i vehicles_db psql -U postgres -d cars < path/to/file.sql

# Check how many cars are loaded
docker exec -i vehicles_db psql -U postgres -d cars -c "SELECT COUNT(*) FROM cars;"

# Reset the database (drop and recreate table)
docker exec -i vehicles_db psql -U postgres -d cars -c "DROP TABLE IF EXISTS cars;"
docker exec -i vehicles_db psql -U postgres -d cars < resources/schema.sql
docker exec -i vehicles_db psql -U postgres -d cars < cars_seeder_utf8.sql

# Stop the database
docker compose stop

# Remove everything (containers + networks, data volume is kept)
docker compose down

-------------------------------------------------------------------------------------

API Endpints

# Health check
curl http://localhost:8080/api/health

# Get all cars
curl http://localhost:8080/api/cars

# Get car by ID
curl http://localhost:8080/api/cars/1

# Search by brand
curl "http://localhost:8080/api/cars/search?brand=Ferrari"

# Search by fuel type
curl "http://localhost:8080/api/cars/search?fuel_type=Petrol"

# Search by price range
curl "http://localhost:8080/api/cars/search?min_price=10000&max_price=50000"

# Combined filters
curl "http://localhost:8080/api/cars/search?brand=BMW&fuel_type=Petrol&max_price=80000"

# Get AI Car Recommendation (Requires .env configured)
curl -X POST http://localhost:8080/api/recommend \
  -H "Content-Type: application/json" \
  -d '{
    "preferences": "I am looking for a fast and reliable car for my family.",
    "filters": {
      "FuelType": "Petrol",
      "MinSeats": 4
    }
  }'

-------------------------------------------------------------------------------------

Seeder (regenerate SQL from CSV)

# Regenerate cars_seeder.sql from the CSV file
go run cmd/seeder/mainCSV.go > cars_seeder.sql

# Convert encoding and load into DB
iconv -f UTF-16 -t UTF-8 cars_seeder.sql > cars_seeder_utf8.sql
docker exec -i vehicles_db psql -U postgres -d cars < cars_seeder_utf8.sql


