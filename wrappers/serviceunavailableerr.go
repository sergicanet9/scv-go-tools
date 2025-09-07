package wrappers

// ServiceUnavailableErr is an error of type serviceUnavailableError with the underlying error message
var ServiceUnavailableErr error = serviceUnavailableError{msg: "service unavailable"}

// serviceUnavailableError is an implementation of error interface
type serviceUnavailableError struct {
	msg string
}

// NewServiceUnavailableErr wraps the given error in a serviceUnavailableError
func NewServiceUnavailableErr(err error) error {
	if err == nil {
		return nil
	}

	return serviceUnavailableError{
		msg: err.Error(),
	}
}

// Error returns the error message
func (e serviceUnavailableError) Error() string {
	return e.msg
}

// Is returns true if the target error is a serviceUnavailableError
func (e serviceUnavailableError) Is(tgt error) bool {
	_, ok := tgt.(serviceUnavailableError)
	return ok
}
