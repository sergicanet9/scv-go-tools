package infrastructure

import (
	"context"
	"database/sql"
	"fmt"
	"time"

	_ "github.com/lib/pq"
	"github.com/pressly/goose/v3"
)

// ConnectPostgresDB opens a connection to the PostgresDB and ensures that the db is reachable
func ConnectPostgresDB(ctx context.Context, dsn string) (*sql.DB, error) {
	db, err := sql.Open("postgres", dsn)
	return db, pingSql(ctx, db, err)
}

func pingSql(ctx context.Context, db *sql.DB, err error) error {
	if db == nil || err != nil {
		return fmt.Errorf("an unexpected error happened while opening the connection: %s", err)
	}

	// wait until db is ready
	for start := time.Now(); time.Since(start) < (5 * time.Second); {
		err = db.PingContext(ctx)
		if err == nil {
			break
		}

		time.Sleep(1 * time.Second)
	}
	return err
}

// MigratePostgresDB runs all migrations found in the given directory against the db
func MigratePostgresDB(db *sql.DB, migrationsDir string) error {
	goose.SetTableName("public.goose_db_version")
	return goose.Up(db, migrationsDir)
}

// PostgresRepository struct of a mongo repository
// Needs a specific implementation of the Repository interface for every entity
type PostgresRepository struct {
	DB *sql.DB
}
