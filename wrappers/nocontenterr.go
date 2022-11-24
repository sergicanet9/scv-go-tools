package wrappers

// NonExistentErr is an error of type nonExistentError with the underlying error message
var NonExistentErr error = nonExistentError{msg: "non existent resource"}

// nonExistentError is an implementation of error interface
type nonExistentError struct {
	msg string
}

// NewNonExistentErr wraps the given error in a nonExistentError
func NewNonExistentErr(err error) error {
	if err == nil {
		return nil
	}

	return nonExistentError{
		msg: err.Error(),
	}
}

// Error returns the error message
func (e nonExistentError) Error() string {
	return e.msg
}

// Is returns true if the target error is a nonExistentError
func (e nonExistentError) Is(tgt error) bool {
	_, ok := tgt.(nonExistentError)
	return ok
}
