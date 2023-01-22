package testutils

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func testFunction() bool {
	return true
}

// TestFunctionName_Ok checks that FunctionName returns the expected value
func TestFunctionName_Ok(t *testing.T) {
	// Arrange
	expectedFunctionName := "testFunction"

	// Act
	functionName := FunctionName(t, testFunction)

	// Assert
	assert.Equal(t, expectedFunctionName, functionName)
}
