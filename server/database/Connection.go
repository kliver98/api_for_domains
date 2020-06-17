package database

import (
    "database/sql"
    "log"

    _ "github.com/lib/pq"
)

var SQL string = "postgresql://root@localhost:26257/reto_prueba?sslmode=disable"
var HISTORY_TABLE string = "CREATE TABLE IF NOT EXISTS history (id INT PRIMARY KEY, domain STRING(200) NOT NULL, timestamptz TIMESTAMPTZ NOT NULL)"
var SEQUENCE_ID_PRUEBA = "CREATE SEQUENCE IF NOT EXISTS id_prueba START 1 INCREMENT 1"
var HISTORY_DATA_TEST string = "INSERT INTO history (id,domain,timestamptz) VALUES (nextval('id_prueba'),'google.com',clock_timestamp()) , (nextval('id_prueba'),'truora.com',clock_timestamp())"

/*func InitDatabase(db *sql.DB) (*[]error) {

    rows, err := Query(db, "SELECT * FROM history")
    if err != nil {
        log.Fatal(err)
    }
    defer rows.Close()
    fmt.Println("	>>>")
    for rows.Next() {
        var id int
        var domain string
        var timestamptz = time.Now()
        if err := rows.Scan(&id, &domain, &timestamptz); err != nil {
            log.Fatal(err)
        }
        fmt.Printf("%d %s %s\n", id, domain, timestamptz)
    }
    return nil
}*/

func GetConnection() (*sql.DB, error) {

	// Connect to the "reto_prueba" database.
    db, err := sql.Open("postgres", SQL)

    return db, err

}