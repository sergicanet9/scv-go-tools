package infrastructure

import (
	"context"
	"testing"

	"github.com/sergicanet9/scv-go-tools/v3/mocks"
	"github.com/stretchr/testify/assert"
)

// TestConnectOracleDB_InvalidDSN checks that ConnectOracleDB returns an error when an invalid DSN is provided
//func TestConnectOracleDB_InvalidDSN(t *testing.T) {
//	// Arrange
//	expectedError := "missing \"=\" after \"invalid-dsn\" in connection info string\""
//
//	// Act
//	_, err := ConnectOracleDB(context.Background(), "invalid-dsn")
//
//	// Assert
//	assert.Equal(t, expectedError, err.Error())
//}

// TestPingOracle_Ok checks that pingOracle does not return an error when a valid db is received
func TestPingOracle_Ok(t *testing.T) {
	// Arrange
	_, db := mocks.NewSqlDB(t)

	// Act
	err := pingOracle(context.Background(), db)

	// Assert
	assert.Nil(t, err)
}

// TestMigrateOracleDB_NotValidDirectory checks that MigrateOracleDB retuns an error when the given directory does not exist
func TestMigrateOracleDB_NotValidDirectory(t *testing.T) {
	// Arrange
	_, db := mocks.NewSqlDB(t)
	expectedError := "invalid-directory directory does not exist"

	// Act
	err := MigrateOracleDB(db, "invalid-directory")

	// Assert
	assert.Equal(t, expectedError, err.Error())
}
