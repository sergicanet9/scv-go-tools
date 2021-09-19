package infrastructure

import (
	"database/sql"

	_ "github.com/lib/pq" // postgres driver, needs to be imported
)

//ConnectPostgresDB connect to PostgresDB
func ConnectPostgresDB(connection string) (*sql.DB, error) {
	db, err := sql.Open("postgres", connection)
	if err != nil {
		return nil, err
	}
	return db, nil
}
