package infrastructure

import (
	"context"
	"testing"

	"github.com/sergicanet9/scv-go-tools/v3/mocks"
	"github.com/stretchr/testify/assert"
	"go.mongodb.org/mongo-driver/mongo/integration/mtest"
)

// TestConnectMongoDB_InvalidConnection checks that ConnectMongoDB returns an error when an invalid connection string is provided
func TestConnectMongoDB_InvalidConnection(t *testing.T) {
	// Arrange
	expectedError := "an unexpected error happened while opening the connection: error parsing uri: scheme must be \"mongodb\" or \"mongodb+srv\""

	// Act
	_, err := ConnectMongoDB(context.Background(), "test", "invalid-connection")

	// Assert
	assert.Equal(t, expectedError, err.Error())
}

// TestPingMongo_Ok checks that pingMongo does not return an error when a valid db is received
func TestPingMongo_Ok(t *testing.T) {
	mt := mocks.NewMongoDB(t)
	mt.Run("", func(mt *mtest.T) {
		// Arrange
		mt.AddMockResponses(mtest.CreateSuccessResponse())

		// Act
		_, err := pingMongo(mt.Client, "test", nil)

		// Assert
		assert.Nil(mt, err, "Ping error: %v", err)
		// assert.Equal(t, nil, err)
	})
}

// TestPingMongo_NilDB checks that pingMongo returns an error when a nil db is received
func TestPingMongo_NilDB(t *testing.T) {
	// Arrange
	expectedError := "an unexpected error happened while opening the connection: %!s(<nil>)"
	// Act
	_, err := pingMongo(nil, "", nil)

	// Assert
	assert.Equal(t, expectedError, err.Error())
}
