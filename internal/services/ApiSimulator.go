package services

import (
	"database/sql"
	"log"
	"math/rand"
	"time"
)

// Global variables for easy adjustment
var (
	VariationPercent  float64 = 5.0
	TargetCarID       int     = 0
	ExecutionInterval int     = 10
)

type ApiSimulator struct {
	DB *sql.DB
}

func (s *ApiSimulator) Start() {
	log.Printf("[SIMULATOR] Running in background every %ds", ExecutionInterval)
	for {
		time.Sleep(time.Duration(ExecutionInterval) * time.Second)

		var idToUpdate int
		if TargetCarID == 0 {
			idToUpdate = s.getRandomID()
		} else {
			idToUpdate = TargetCarID
		}

		if idToUpdate > 0 {
			randomRange := (rand.Float64() * 2) - 1
			factor := 1.0 + (randomRange * (VariationPercent / 100.0))
			s.applyChange(idToUpdate, factor)
		}
	}
}

func (s *ApiSimulator) getRandomID() int {
	var id int
	// Adjusted for your cars table
	query := "SELECT id FROM cars ORDER BY RANDOM() LIMIT 1"
	err := s.DB.QueryRow(query).Scan(&id)
	if err != nil {
		return 0
	}
	return id
}

func (s *ApiSimulator) applyChange(id int, factor float64) {
	query := "UPDATE cars SET price = price * $1 WHERE id = $2"
	_, err := s.DB.Exec(query, factor, id)
	if err != nil {
		log.Printf("[SIM_ERROR] ID %d: %v", id, err)
		return
	}
	log.Printf("[SIM_SUCCESS] ID %d updated by factor %.4f", id, factor)
}
