package db

import (
	"database/sql"
	"fmt"
)

func Connect() (*sql.DB, error) {
	// Replace these with your PostgreSQL database credentials
	connStr := "user=myuser dbname=mydb password=mypassword host=localhost port=5432 sslmode=disable"

	// Open a connection to the database
	db, err := sql.Open("postgres", connStr)
	if err != nil {
		return nil, err
	}
	defer db.Close()

	// Test the connection
	err = db.Ping()
	if err != nil {
		return nil, err
	}

	fmt.Println("Connected to the PostgreSQL database!")

	return db, nil
}
