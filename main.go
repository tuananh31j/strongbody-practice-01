package main

import (
	"database/sql"
	"encoding/csv"
	"fmt"
	"log"
	"os"
	"time"

	"github.com/google/uuid"
	_ "github.com/lib/pq"
)

type Record struct {
	ID           int
	Name         string
	Email        string
	IsActive     bool
	CreatedAt    time.Time
	UpdatedAt    time.Time
	UUID         uuid.UUID
	Gender         string
	Age          int
	Salary       float64
	JoiningDate  time.Time
}

func main() {
	file, err := os.Open("data.csv")
	if err != nil {
		log.Fatal(err, "File not found")
	}
	defer file.Close()

	reader := csv.NewReader(file)
	records, err := reader.ReadAll()
	if err != nil {
		log.Fatal(err)
	}

	connStr := "host=localhost port=5432 user=postgres password=123 dbname=postgres sslmode=disable"
	db, err := sql.Open("postgres", connStr)
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	db.Exec("DELETE FROM users")


	for _, row := range records[1:] {
		record := Record{
			Name:         row[1],
			Email:        row[2],
			IsActive:     row[3] == "true",
			CreatedAt:    parseDate(row[4]),
			UpdatedAt:    parseDate(row[5]),
			UUID:         parseUUID(row[6]),
			Age:          parseInt(row[7]),
			Salary:       parseFloat(row[8]),
			JoiningDate:  parseDate(row[9]),
			Gender:  row[10],
		}

		_, err = db.Exec(`INSERT INTO users (name, email, is_active, created_at, updated_at, age, salary, joining_date, uuid, gender) 
		VALUES ($1, $2, $3, $4, $5, $6, $7, $8, $9, $10)`,
			record.Name, record.Email, record.IsActive, record.CreatedAt, record.UpdatedAt, record.Age, record.Salary, record.JoiningDate, record.UUID, record.Gender)
		if err != nil {
			log.Fatal(err)
		}
	}

	fmt.Println("Data inserted successfully.")
	fetchData(db)
}

func parseDate(dateStr string) time.Time {
	date, err := time.Parse("2006-01-02", dateStr)
	if err != nil {
		log.Fatal(err)
	}
	return date
}

func parseUUID(uuidStr string) uuid.UUID {
	parsedUUID, err := uuid.Parse(uuidStr)
	if err != nil {
		log.Fatal(err)
	}
	return parsedUUID
}

func parseInt(intStr string) int {
	var value int
	fmt.Sscanf(intStr, "%d", &value)
	return value
}

func parseFloat(floatStr string) float64 {
	var value float64
	fmt.Sscanf(floatStr, "%f", &value)
	return value
}
func fetchData(db *sql.DB) {
	rows, err := db.Query("SELECT * FROM users")
	if err != nil {
		log.Fatal(err)
	}
	defer rows.Close()

	for rows.Next() {
		var record Record
		err := rows.Scan(&record.ID, &record.Name, &record.Email, &record.IsActive, &record.CreatedAt, &record.UpdatedAt, &record.UUID, &record.Age, &record.Salary, &record.JoiningDate, &record.Gender)
		if err != nil {
			log.Fatal(err)
		}
		fmt.Printf("%d, %s, %s, %t, %s, %s, %s, %d, %.2f, %s, %v\n", record.ID, record.Name, record.Email, record.IsActive, record.CreatedAt, record.UpdatedAt, record.UUID, record.Age, record.Salary, record.JoiningDate, record.Gender)
	}
}

