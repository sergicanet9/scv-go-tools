package observability

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"go.mongodb.org/mongo-driver/event"
)

// TestSetupNewRelic_InvalidKey checks that an error is returned when an invalid New Relic key is provided
func TestSetupNewRelic_InvalidKey(t *testing.T) {
	// Arrange
	appName := "test-app"
	invalidKey := "invalid-key"

	// Act
	_, err := SetupNewRelic(appName, invalidKey)

	// Assert
	assert.NotEmpty(t, err)
}

func TestNewRelicMongoMonitor(t *testing.T) {
	// Arrange
	var expectedType *event.CommandMonitor

	// Act
	monitor := NewRelicMongoMonitor()

	// Assert
	assert.NotNil(t, monitor)
	assert.IsType(t, expectedType, monitor)
}
