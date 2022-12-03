package wrappers

import (
	"errors"
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

// TestNewUnauthorizedErr_Ok checks that NewUnauthorizedErr returns the expected error type when receives an error
func TestNewUnauthorizedErr_Ok(t *testing.T) {
	// Arrange
	err := fmt.Errorf("test error")

	// Act
	gotErr := NewUnauthorizedErr(err)

	// Assert
	assert.NotEmpty(t, gotErr)
	assert.IsType(t, UnauthorizedErr, gotErr)
}

// TestNewUnauthorizedErr_NilErr checks that NewUnauthorizedErr returns nil when receives a nil error
func TestNewUnauthorizedErr_NilErr(t *testing.T) {
	// Arrange
	var err error

	// Act
	gotErr := NewUnauthorizedErr(err)

	// Assert
	assert.Nil(t, gotErr)
}

// TestUnauthorizedErrError_Ok checks that Error returns the expected error message of the receiver
func TestUnauthorizedErrError_Ok(t *testing.T) {
	// Arrange
	expectedMsg := "test error"
	err := NewUnauthorizedErr(fmt.Errorf(expectedMsg))

	// Act
	gotMsg := err.Error()

	// Assert
	assert.Equal(t, expectedMsg, gotMsg)
}

// TestUnauthorizedErrIs_True checks that Is returns true when the receiver is an unauthorizedError
func TestUnauthorizedErrIs_True(t *testing.T) {
	// Arrange
	err := NewUnauthorizedErr(fmt.Errorf("test"))

	// Act
	isUnauthorizedErr := errors.Is(err, UnauthorizedErr)

	// Assert
	assert.True(t, isUnauthorizedErr)
}

// TestUnauthorizedErrIs_False checks that Is returns false when the receiver is not an unauthorizedError
func TestUnauthorizedErrIs_False(t *testing.T) {
	// Arrange
	err := fmt.Errorf("test")

	// Act
	isUnauthorizedErr := errors.Is(err, UnauthorizedErr)

	// Assert
	assert.False(t, isUnauthorizedErr)
}
