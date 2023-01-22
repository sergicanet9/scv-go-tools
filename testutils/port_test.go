package testutils

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

// TestFreePort_Ok checks that FreePort returns an available port
func TestFreePort_Ok(t *testing.T) {
	// Act
	port := FreePort(t)

	// Assert
	assert.NotEmpty(t, port)
}
