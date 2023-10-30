package database

import (
	"database/sql"
	"fmt"
	"log"
	_ "github.com/lib/pq"
)

func main() {
	// Define connection parameters
	host := "localhost"     // Replace with the PostgreSQL server's host or IP address
	port := 5432            // PostgreSQL default port
	user := "postgres"      // Your PostgreSQL username
	password := "Berat9730" // Your PostgreSQL password
	dbname := "GoLang"      // Your PostgreSQL database name
	// Connection string
	connStr := fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s sslmode=disable", host, port, user, password, dbname)
	// Establish a database connection
	db, err := sql.Open("postgres", connStr)
	if err != nil {
		log.Fatal(err)
		return
	}
	defer db.Close()
	// Ping the database to test the connection
	if err = db.Ping(); err != nil {
		log.Fatal(err)
		return
	}
	fmt.Printf("Connection succesfull")
}
