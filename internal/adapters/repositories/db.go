package repositories

import (
	"database/sql"
	"fmt"
	"log"

	_ "github.com/lib/pq" // Driver de PostgreSQL
)

// InitDB inicializa la conexión con la base de datos PostgreSQL
func InitDB(connStr string) (*sql.DB, error) {
	db, err := sql.Open("postgres", connStr)
	if err != nil {
		return nil, fmt.Errorf("error opening database: %w", err)
	}

	// Comprobar que la conexión  funciona
	if err = db.Ping(); err != nil {
		return nil, fmt.Errorf("error connecting to database: %w", err)
	}

	log.Println("Successfully connected to the database!")
	return db, nil
}
