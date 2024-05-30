// main_test.go
package main

import (
	"database/sql"
	"testing"

	_ "modernc.org/sqlite"
)

// TestDatabaseConnection tests the connection to the SQLite database.
func TestDatabaseConnection(t *testing.T) {
	db, err := sql.Open("sqlite", "imdb_database.db")
	if err != nil {
		t.Fatalf("Failed to connect to database: %v", err)
	}
	defer db.Close()

	err = db.Ping()
	if err != nil {
		t.Fatalf("Failed to ping database: %v", err)
	}
}

// TestInsertData tests the insertion of data into the SQLite database.
func TestInsertData(t *testing.T) {
	db, err := sql.Open("sqlite", "imdb_database.db")
	if err != nil {
		t.Fatalf("Failed to connect to database: %v", err)
	}
	defer db.Close()

	// Adjusted insert statement based on actual schema
	_, err = db.Exec(`INSERT INTO actors (id, first_name, last_name, gender) VALUES (?, ?, ?, ?)`, 1, "Test", "Actor", "M")
	if err != nil {
		t.Fatalf("Failed to insert data: %v", err)
	}
}
