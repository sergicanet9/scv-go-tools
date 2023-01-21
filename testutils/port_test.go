package testutils

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestFreePort_Ok(t *testing.T) {
	// Act
	port := FreePort(t)

	// Assert
	assert.NotEmpty(t, port)

}
