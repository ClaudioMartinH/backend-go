package connection

import (
	"database/sql"
	"fmt"
	"os"
)

func ConnectToDatabase() (*sql.DB, error) {
	dbUser := os.Getenv("DB_USER")
	dbPassword := os.Getenv("DB_PASSWORD")
	dbName := os.Getenv("DB_NAME")
	dbHost := os.Getenv("DB_HOST")
	dbPort := os.Getenv("DB_PORT")

	// Construct the connection string
	connectionString := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s", dbUser, dbPassword, dbHost, dbPort, dbName)

	// Open the database connection
	db, err := sql.Open("mysql", connectionString)
	if err != nil {
		return nil, err
	}

	// Test the connection
	err = db.Ping()
	if err != nil {
		return nil, err
	}

	fmt.Println("Connected to database successfully!")
	return db, nil
}
