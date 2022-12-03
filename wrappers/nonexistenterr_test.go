package wrappers

import (
	"errors"
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

// TestNewNonExistentErr_Ok checks that NewNonExistentErr returns the expected error type when receives an error
func TestNewNonExistentErr_Ok(t *testing.T) {
	// Arrange
	err := fmt.Errorf("test error")

	// Act
	gotErr := NewNonExistentErr(err)

	// Assert
	assert.NotEmpty(t, gotErr)
	assert.IsType(t, NonExistentErr, gotErr)
}

// TestNewNonExistentErr_NilErr checks that NewNonExistentErr returns nil when receives a nil error
func TestNewNonExistentErr_NilErr(t *testing.T) {
	// Arrange
	var err error

	// Act
	gotErr := NewNonExistentErr(err)

	// Assert
	assert.Nil(t, gotErr)
}

// TestUNonExistentErrError_Ok checks that Error returns the expected error message of the receiver
func TestUNonExistentErrError_Ok(t *testing.T) {
	// Arrange
	expectedMsg := "test error"
	err := NewNonExistentErr(fmt.Errorf(expectedMsg))

	// Act
	gotMsg := err.Error()

	// Assert
	assert.Equal(t, expectedMsg, gotMsg)
}

// TestNonExistentErrIs_True checks that Is returns true when the receiver is a nonExistentError
func TestNonExistentErrIs_True(t *testing.T) {
	// Arrange
	err := NewNonExistentErr(fmt.Errorf("test"))

	// Act
	isNonExistentErr := errors.Is(err, NonExistentErr)

	// Assert
	assert.True(t, isNonExistentErr)
}

// TestNonExistentErrIs_False checks that Is returns false when the receiver is not a nonExistentError
func TestNonExistentErrIs_False(t *testing.T) {
	// Arrange
	err := fmt.Errorf("test")

	// Act
	isNonExistentErr := errors.Is(err, NonExistentErr)

	// Assert
	assert.False(t, isNonExistentErr)
}
