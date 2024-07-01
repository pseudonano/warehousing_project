package configuration

import (
	"database/sql"
	"fmt"
	"log"

	_ "github.com/lib/pq" // PostgreSQL driver
)

// DB is a global database connection
type DB struct {
	*sql.DB
}

func (d *DB) Connect() {
	config := GetConfig()
	dsn := config.DSN()

	var err error
	d.DB, err = sql.Open("postgres", dsn)
	if err != nil {
		log.Fatalf("Error opening database: %s\n", err)
	}

	if err = d.Ping(); err != nil {
		log.Fatalf("Error pinging database: %s\n", err)
	}

	fmt.Println("Successfully connected to the database")
}
