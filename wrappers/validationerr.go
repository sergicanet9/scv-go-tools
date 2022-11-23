package wrappers

// ValidationErr is an error of type validationError with the underlying error message
var ValidationErr error = validationError{msg: "validation failed"}

// validationError is an implementation of error interface
type validationError struct {
	msg string
}

// NewValidationErr wraps the given error in a validationError
func NewValidationErr(err error) error {
	if err == nil {
		return nil
	}

	return validationError{
		msg: err.Error(),
	}
}

// Error returns the error message
func (e validationError) Error() string {
	return e.msg
}

// Is returns true if the target error is a validationError
func (e validationError) Is(tgt error) bool {
	_, ok := tgt.(validationError)
	return ok
}
