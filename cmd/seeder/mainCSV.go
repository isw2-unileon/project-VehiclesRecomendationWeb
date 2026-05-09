package main

import (
	"encoding/csv"
	"fmt"
	"math/rand"
	"os"
	"strconv"
	"strings"
	"time"
)

// escapeSQL escapes single quotes to prevent SQL syntax errors
func escapeSQL(s string) string {
	return strings.ReplaceAll(s, "'", "''")
}

// processNumericField cleans text, handles ranges, and returns a numeric type (int or float64)
func processNumericField(val string, r *rand.Rand) interface{} {
	// 1. Clean visual characters, currency symbols, and units
	replacer := strings.NewReplacer(
		"$", "",
		",", "",
		" cc", "",
		" hp", "",
		" km/h", "",
		" sec", "",
		" Nm", "",
		" ", "", // Remove spaces
	)
	cleaned := replacer.Replace(val)

	// If the field is empty after cleaning, return SQL NULL
	if cleaned == "" {
		return "NULL"
	}

	// 2. Range handling: if it contains a hyphen, pick a random integer within that range
	if strings.Contains(cleaned, "-") {
		parts := strings.Split(cleaned, "-")
		if len(parts) == 2 {
			minF, err1 := strconv.ParseFloat(parts[0], 64)
			maxF, err2 := strconv.ParseFloat(parts[1], 64)

			if err1 == nil && err2 == nil {
				min := int(minF)
				max := int(maxF)

				if min > max {
					min, max = max, min
				}

				rangeSize := max - min + 1
				if rangeSize > 0 {
					randomValue := min + r.Intn(rangeSize)
					return randomValue
				}
			}
		}
	}

	// 3. Try to parse as a single number (float/int)
	if num, err := strconv.ParseFloat(cleaned, 64); err == nil {
		return num
	}

	// 4. Fallback: return formatted text for SQL
	return fmt.Sprintf("'%s'", escapeSQL(val))
}

func main() {
	// Initialize random seed
	source := rand.NewSource(time.Now().UnixNano())
	r := rand.New(source)

	// Open the CSV file from the resources folder
	file, err := os.Open("resources/Cars Datasets 2025.csv")
	if err != nil {
		fmt.Printf("Error opening file: %v\n", err)
		return
	}
	defer file.Close()

	reader := csv.NewReader(file)

	// Read all records
	records, err := reader.ReadAll()
	if err != nil {
		fmt.Printf("Error reading CSV: %v\n", err)
		return
	}

	// Check if there is data
	if len(records) < 2 {
		fmt.Println("The file is empty or contains only headers.")
		return
	}

	// Iterate over records (skip header row)
	for i, row := range records {
		if i == 0 {
			continue
		}

		// Ensure enough columns exist (adjust to 11 based on your CSV)
		if len(row) < 11 {
			continue
		}

		// Text fields
		company := fmt.Sprintf("'%s'", escapeSQL(row[0]))
		carName := fmt.Sprintf("'%s'", escapeSQL(row[1]))
		engine := fmt.Sprintf("'%s'", escapeSQL(row[2]))
		fuel := fmt.Sprintf("'%s'", escapeSQL(row[8]))

		// Numeric fields (processed via helper function)
		cc := processNumericField(row[3], r)
		hp := processNumericField(row[4], r)
		speed := processNumericField(row[5], r)
		acceleration := processNumericField(row[6], r)
		price := processNumericField(row[7], r)
		seats := processNumericField(row[9], r)
		torque := processNumericField(row[10], r)

		// SQL Query in English
		query := fmt.Sprintf("INSERT INTO cars (company, car_name, engine, capacity_cc, power_hp, max_speed_kmh, acceleration_0_100_sec, price, fuel_type, seats, torque_nm) VALUES (%v, %v, %v, %v, %v, %v, %v, %v, %v, %v, %v);",
			company, carName, engine, cc, hp, speed, acceleration, price, fuel, seats, torque)

		fmt.Println(query)
	}
}
