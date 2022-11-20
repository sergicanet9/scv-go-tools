package infrastructure

import (
	"database/sql"
	"fmt"
	"time"

	_ "github.com/lib/pq"
)

// ConnectPostgresDB connects to PostgresDB and ensures that the db is reachable
func ConnectPostgresDB(connection string) (*sql.DB, error) {
	db, err := sql.Open("postgres", connection)
	return db, pingSql(db, err)
}

func pingSql(db *sql.DB, err error) error {
	if db == nil || err != nil {
		return fmt.Errorf("an unexpected error happened while opening the connection: %s", err)
	}

	// wait until db is ready
	for start := time.Now(); time.Since(start) < (5 * time.Second); {
		err = db.Ping()
		if err == nil {
			break
		}

		time.Sleep(1 * time.Second)
	}
	return err
}

// PostgresRepository struct of a mongo repository
// Needs a specific implementation of the Repository interface for every entity
type PostgresRepository struct {
	DB *sql.DB
}
