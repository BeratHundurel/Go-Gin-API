package database

import (
	"example/data-acces/internal/app/models"
	"fmt"
	"github.com/jmoiron/sqlx"
	_ "github.com/lib/pq"
	"log"
)

var DB *sqlx.DB // Export the database connection

func InitDB() {
	// Define connection parameters
	host := "localhost"
	port := 5432
	user := "postgres"
	password := "Berat9730"
	dbname := "GoLang"

	connStr := fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s sslmode=disable", host, port, user, password, dbname)

	// Establish a database connection
	db, err := sqlx.Connect("postgres", connStr)
	if err != nil {
		log.Fatal(err)
		return
	}

	DB = db // Set the DB variable to the opened database connection
}

// GetUsers fetches a list of users from the database and returns them.
func GetUsers() ([]models.User, error) {
	rows, err := DB.Queryx(`SELECT*FROM public."User"`)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var users []models.User

	for rows.Next() {
		var user models.User
		if err := rows.Scan(&user.Id, &user.Name, &user.Email, &user.Password); err != nil {
			return nil, err
		}
		users = append(users, user)
	}

	if err := rows.Err(); err != nil {
		return nil, err
	}

	return users, nil
}

func GetUser(id string) (models.User, error) {
	var user models.User
	err := DB.Get(&user, `SELECT*FROM public."user" WHERE id=$1`, id)
	if err != nil {
		return user, err
	}
	return user, nil
}

