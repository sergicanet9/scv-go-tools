package infrastructure

import (
	"context"
	"testing"

	"github.com/sergicanet9/scv-go-tools/v3/mocks"
	"github.com/stretchr/testify/assert"
)

// TestConnectPostgresDB_InvalidDSN checks that ConnectPostgresDB returns an error when an invalid DSN is provided
func TestConnectPostgresDB_InvalidDSN(t *testing.T) {
	// Arrange
	expectedError := "missing \"=\" after \"invalid-dsn\" in connection info string\""

	// Act
	_, err := ConnectPostgresDB(context.Background(), "invalid-dsn")

	// Assert
	assert.Equal(t, expectedError, err.Error())
}

// TestPingSql_Ok checks that pingSql does not return an error when a valid db is received
func TestPingSql_Ok(t *testing.T) {
	// Arrange
	_, db := mocks.NewSqlDB(t)

	// Act
	err := pingSql(context.Background(), db)

	// Assert
	assert.Nil(t, err)
}

// TestMigratePostgresDB_NotValidDirectory checks that MigratePostgresDB retuns an error when the given directory does not exist
func TestMigratePostgresDB_NotValidDirectory(t *testing.T) {
	// Arrange
	_, db := mocks.NewSqlDB(t)
	expectedError := "invalid-directory directory does not exist"

	// Act
	err := MigratePostgresDB(db, "invalid-directory")

	// Assert
	assert.Equal(t, expectedError, err.Error())
}
