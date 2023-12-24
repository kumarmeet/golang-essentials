package db

import (
	"database/sql"
	"log" // Import the log package for logging errors

	_ "github.com/go-sql-driver/mysql"
)

var DB *sql.DB

func InitDB() {
	// Use a single equal sign to assign to the global variable instead of :=
	var err error
	DB, err = sql.Open("mysql", "root:password@tcp(127.0.0.1:3306)/event")

	if err != nil {
		log.Fatal("DB not connected:", err) // Use log.Fatal to log the error and exit the program
	}

	DB.SetMaxOpenConns(10)
	DB.SetMaxIdleConns(5)

	createTables()
}

func createTables() {
	createEventsTable := `
	CREATE TABLE IF NOT EXISTS events (
		id INTEGER PRIMARY KEY AUTO_INCREMENT, -- Use AUTO_INCREMENT for MySQL
		name VARCHAR(255) NOT NULL,
		description TEXT NOT NULL, -- Use TEXT for unlimited length
		location VARCHAR(255) NOT NULL,
		created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP, -- Use DEFAULT CURRENT_TIMESTAMP for auto-setting timestamps
		updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
		user_id INTEGER
	);
`
	_, err := DB.Exec(createEventsTable)

	if err != nil {
		log.Fatal("Could not create events table:", err) // Use log.Fatal to log the error and exit the program
	}
}
