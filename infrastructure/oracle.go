package infrastructure

import (
	"context"
	"database/sql"
	"fmt"
	"strings"
	"time"

	_ "github.com/godror/godror"
	"github.com/pressly/goose/v3"
)

// ConnectOracleDB opens a connection to the OracleDB and ensures that the db is reachable
func ConnectOracleDB(ctx context.Context, dsn string) (*sql.DB, error) {
	// TODO: convert dsn to godror format
	arguments := strings.Split(dsn, "@")
	arguments2 := strings.Split(arguments[0], ":")
	username := strings.Split(arguments2[1], "//")
	password := arguments2[2]
	connectString := arguments[1]

	dsnOracle := fmt.Sprintf(`user="%s" password="%s" connectString="%s"`, username[1], password, connectString)

	db, _ := sql.Open("godror", dsnOracle)
	return db, pingOracle(ctx, db)
}

func pingOracle(ctx context.Context, db *sql.DB) (err error) {
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

// MigrateOracleDB runs all migrations found in the given directory against the db
func MigrateOracleDB(db *sql.DB, migrationsDir string) error {
	goose.SetTableName("goose_db_version")
	return goose.Up(db, migrationsDir)
}

// OracleRepository struct of a oracle repository
// Needs a specific implementation of the Repository interface for every entity
type OracleRepository struct {
	DB *sql.DB
}
