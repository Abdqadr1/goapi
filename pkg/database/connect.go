package database

import (
	"database/sql"
	"fmt"

	_ "github.com/go-sql-driver/mysql"
)

func ConnectDB() (*sql.DB, error) {
	// Replace with your credentials
	dsn := "root:@tcp(localhost:3306)/go_todos?parseTime=true"

	// Open a connection
	db, err := sql.Open("mysql", dsn)
	if err != nil {
		return nil, fmt.Errorf("failed to connect to database: %w", err)
	}

	// Verify the connection
	if err := db.Ping(); err != nil {
		return nil, fmt.Errorf("failed to verify connection: %w", err)
	}

	fmt.Println("Database connection established!")
	return db, nil
}
