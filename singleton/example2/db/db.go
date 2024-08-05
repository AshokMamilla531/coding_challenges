package db

import (
	"database/sql"
	"log"
	"sync"

	_ "github.com/lib/pq"
)

var (
	instance *sql.DB
	once     sync.Once
)

// GetInstance returns the singleton instance of the database connection
func GetInstance() *sql.DB {
	once.Do(func() {
		var err error
		connStr := "user=yourusername password=yourpassword dbname=coding_challenges sslmode=disable"
		instance, err = sql.Open("postgres", connStr)
		if err != nil {
			log.Fatalf("Failed to connect to the database: %v", err)
		}

		if err = instance.Ping(); err != nil {
			log.Fatalf("Failed to ping the database: %v", err)
		}
	})

	return instance
}
