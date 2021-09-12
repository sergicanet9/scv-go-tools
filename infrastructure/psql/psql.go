package infrastructure

import (
	"database/sql"
)

//ConnectPsqlDB connect to PostgresDB
func ConnectPsqlDB(connection string) (*sql.DB, error) {
	db, err := sql.Open("postgres", connection)
	if err != nil {
		return nil, err
	}
	return db, nil
}
