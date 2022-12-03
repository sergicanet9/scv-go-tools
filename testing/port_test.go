package testing

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestFreePort_Ok(t *testing.T) {
	// Arrange

	// Act
	port := FreePort(t)

	// Assert
	assert.NotEmpty(t, port)

}
