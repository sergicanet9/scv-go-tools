package infrastructure

import "database/sql"

//ConnectPostgresDB connect to PostgresDB
func ConnectPostgresDB(connection string) (*sql.DB, error) {
	db, err := sql.Open("postgres", connection)
	if err != nil {
		return nil, err
	}
	return db, nil
}

// PostgresRepository struct of a mongo repository. Needs a specific implementation for every Repository to be used as an adapter.
type PostgresRepository struct {
	DB *sql.DB
}
