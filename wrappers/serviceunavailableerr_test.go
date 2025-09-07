package wrappers

import (
	"errors"
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

// TestNewServiceUnavailableErr_Ok checks that NewServiceUnavailableErr returns the expected error type when receives an error
func TestNewServiceUnavailableErr_Ok(t *testing.T) {
	// Arrange
	err := fmt.Errorf("test error")

	// Act
	gotErr := NewServiceUnavailableErr(err)

	// Assert
	assert.NotEmpty(t, gotErr)
	assert.IsType(t, ServiceUnavailableErr, gotErr)
}

// TestNewServiceUnavailableErr_NilErr checks that NewServiceUnavailableErr returns nil when receives a nil error
func TestNewServiceUnavailableErr_NilErr(t *testing.T) {
	// Arrange
	var err error

	// Act
	gotErr := NewServiceUnavailableErr(err)

	// Assert
	assert.Nil(t, gotErr)
}

// TestServiceUnavailableErrError_Ok checks that Error returns the expected error message of the receiver
func TestServiceUnavailableErrError_Ok(t *testing.T) {
	// Arrange
	expectedMsg := "test error"
	err := NewServiceUnavailableErr(errors.New(expectedMsg))

	// Act
	gotMsg := err.Error()

	// Assert
	assert.Equal(t, expectedMsg, gotMsg)
}

// TestServiceUnavailableErrIs_True checks that Is returns true when the receiver is a serviceUnavailableError
func TestServiceUnavailableErrIs_True(t *testing.T) {
	// Arrange
	err := NewServiceUnavailableErr(fmt.Errorf("test"))

	// Act
	isServiceUnavailableErr := errors.Is(err, ServiceUnavailableErr)

	// Assert
	assert.True(t, isServiceUnavailableErr)
}

// TestServiceUnavailableErrIs_False checks that Is returns false when the receiver is not a serviceUnavailableError
func TestServiceUnavailableErrIs_False(t *testing.T) {
	// Arrange
	err := fmt.Errorf("test")

	// Act
	isServiceUnavailableErr := errors.Is(err, ServiceUnavailableErr)

	// Assert
	assert.False(t, isServiceUnavailableErr)
}
