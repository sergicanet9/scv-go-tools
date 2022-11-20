package infrastructure

import (
	"context"
	"testing"

	"github.com/stretchr/testify/assert"
)

// TestConnectMongoDB_InvalidConnection checks that ConnectMongoDB returns an error when an invalid connection string is provided
func TestConnectMongoDB_InvalidConnection(t *testing.T) {
	// Arrange
	expectedError := "error parsing uri: scheme must be \"mongodb\" or \"mongodb+srv\""

	// Act
	_, err := ConnectMongoDB(context.Background(), "test", "invalid-connection")

	// Assert
	assert.Equal(t, expectedError, err.Error())
}

// // TestPingMongo_Ok checks that pingMongo does not return an error when a valid db is received
// func TestPingMongo_Ok(t *testing.T) {
// 	mt := mocks.NewMongoDB(t)
// 	mt.Run("", func(mt *mtest.T) {
// 		// Arrange
// 		ping := bson.D{{"ping", "true"}}
// 		mt.AddMockResponses(ping)
// 		err2 := mt.Client.Ping(context.Background(), mtest.PrimaryRp)
// 		fmt.Print(err2)
// 		// Act
// 		err := pingMongo(mt.DB)

// 		// Assert
// 		assert.Nil(mt, err, "Ping error: %v", err)
// 		// assert.Equal(t, nil, err)
// 	})
// }
