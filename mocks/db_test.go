package mocks

import (
	"database/sql"
	"testing"

	"github.com/DATA-DOG/go-sqlmock"
	"github.com/stretchr/testify/assert"
	"go.mongodb.org/mongo-driver/mongo/integration/mtest"
)

// TestNewMongoDB_Ok checks that NewMongoDB returns a non empty object of the expected type
func TestNewMongoDB_Ok(t *testing.T) {
	// Arrange
	expectedType := &mtest.T{}

	// Act
	gotDB := NewMongoDB(t)

	// Assert
	assert.NotEmpty(t, gotDB)
	assert.IsType(t, expectedType, gotDB)
}

// TestNewSqlDB_Ok checks that NewSqlDB returns non empty objects of the expected types
func TestNewSqlDB_Ok(t *testing.T) {
	// Arrange
	expectedMockInterface := (*sqlmock.Sqlmock)(nil)
	expectedDBType := &sql.DB{}

	// Act
	gotMock, gotDB := NewSqlDB(t)

	// Assert
	assert.NotEmpty(t, gotMock)
	assert.NotEmpty(t, gotDB)
	assert.Implements(t, expectedMockInterface, gotMock)
	assert.IsType(t, expectedDBType, gotDB)
}
