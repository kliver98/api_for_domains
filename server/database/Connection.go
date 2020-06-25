package database

import (
    "database/sql"
    "net/http"
    "time"

    _ "github.com/lib/pq"
    errors "../errors"
)

const SQL string = "postgresql://root@localhost:26257/reto_prueba?sslmode=disable"
const PING_DATABASE = "http://localhost:5001"


func GetConnection() (*sql.DB, error) {

	// Connect to the "reto_prueba" database.
    db, err := sql.Open("postgres", SQL)

    if isDown(PING_DATABASE) {
    	err = &errors.NoPingError{Message: "Unreachable to database "+PING_DATABASE}
    }

    return db, err

}

func isDown(domain string) bool {
	timeout := time.Duration(1 * time.Second)
	client := http.Client{
	    Timeout: timeout,
	}
	_, err := client.Get(domain)
	if err != nil {
	    return true
	} else {
	    return false
	}
}