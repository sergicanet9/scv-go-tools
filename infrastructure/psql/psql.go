package infrastructure

import (
	"database/sql"

	_ "github.com/lib/pq" // postgres driver, needs to be imported
)

//ConnectPsqlDB connect to PostgresDB
func ConnectPsqlDB(connection string) (*sql.DB, error) {
	db, err := sql.Open("postgres", connection)
	if err != nil {
		return nil, err
	}
	return db, nil
}
