package infrastructure

import (
	"database/sql"
	"time"

	_ "github.com/lib/pq"
)

// ConnectPostgresDB connects to PostgresDB and ensures that the db is reachable
func ConnectPostgresDB(connection string) (*sql.DB, error) {
	db, err := openSQL("postgres", connection)
	return db, err
}

func openSQL(driver, connection string) (*sql.DB, error) {
	db, err := sql.Open(driver, connection)
	if err != nil {
		return nil, err
	}

	// wait until db is ready
	for start := time.Now(); time.Since(start) < (5 * time.Second); {
		err = db.Ping()
		if err == nil {
			break
		}

		time.Sleep(1 * time.Second)
		continue
	}

	return db, err
}

// PostgresRepository struct of a mongo repository
// Needs a specific implementation of the Repository interface for every entity
type PostgresRepository struct {
	DB *sql.DB
}
