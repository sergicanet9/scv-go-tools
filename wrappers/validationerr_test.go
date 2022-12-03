package wrappers

import (
	"errors"
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

// TestNewValidationErr_Ok checks that NewValidationErr returns the expected error type when receives an error
func TestNewValidationErr_Ok(t *testing.T) {
	// Arrange
	err := fmt.Errorf("test error")

	// Act
	gotErr := NewValidationErr(err)

	// Assert
	assert.NotEmpty(t, gotErr)
	assert.IsType(t, ValidationErr, gotErr)
}

// TestNewValidationErr_NilErr checks that NewValidationErr returns nil when receives a nil error
func TestNewValidationErr_NilErr(t *testing.T) {
	// Arrange
	var err error

	// Act
	gotErr := NewValidationErr(err)

	// Assert
	assert.Nil(t, gotErr)
}

// TestValidationErrError_Ok checks that Error returns the expected error message of the receiver
func TestValidationErrError_Ok(t *testing.T) {
	// Arrange
	expectedMsg := "test error"
	err := NewValidationErr(fmt.Errorf(expectedMsg))

	// Act
	gotMsg := err.Error()

	// Assert
	assert.Equal(t, expectedMsg, gotMsg)
}

// TestValidationErrIs_True checks that Is returns true when the receiver is a validationError
func TestValidationErrIs_True(t *testing.T) {
	// Arrange
	err := NewValidationErr(fmt.Errorf("test"))

	// Act
	isValidationErr := errors.Is(err, ValidationErr)

	// Assert
	assert.True(t, isValidationErr)
}

// TestValidationErrIs_False checks that Is returns false when the receiver is not a validationError
func TestValidationErrIs_False(t *testing.T) {
	// Arrange
	err := fmt.Errorf("test")

	// Act
	isValidationErr := errors.Is(err, ValidationErr)

	// Assert
	assert.False(t, isValidationErr)
}
