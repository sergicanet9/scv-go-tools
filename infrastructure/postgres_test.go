package infrastructure

import (
	"testing"

	"github.com/sergicanet9/scv-go-tools/v3/test/mocks"
	"github.com/stretchr/testify/assert"
)

// TestConnectPostgresDB_InvalidConnection checks that ConnectPostgresDB returns an error when an invalid connection string is provided
func TestConnectPostgresDB_InvalidConnection(t *testing.T) {
	// Arrange
	expectedError := "missing \"=\" after \"invalid-connection\" in connection info string\""

	// Act
	_, err := ConnectPostgresDB("invalid-connection")

	// Assert
	assert.Equal(t, expectedError, err.Error())
}

// TestPingSql_Ok checks that pingSql does not return an error when a valid db is received
func TestPingSql_Ok(t *testing.T) {
	// Arrange
	_, db := mocks.NewSqlDB(t)
	// Act
	err := pingSql(db)

	// Assert
	assert.Equal(t, nil, err)
}
