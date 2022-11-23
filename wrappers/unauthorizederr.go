package wrappers

// UnauthorizedErr is an error of type unauthorizedError with the underlying error message
var UnauthorizedErr error = unauthorizedError{msg: "unauthorized"}

// unauthorizedError is an implementation of error interface
type unauthorizedError struct {
	msg string
}

// NewUnauthorizedErr wraps the given error in an unauthorizedError
func NewUnauthorizedErr(err error) error {
	if err == nil {
		return nil
	}

	return unauthorizedError{
		msg: err.Error(),
	}
}

// Error returns the error message
func (e unauthorizedError) Error() string {
	return e.msg
}

// Is returns true if the target error is an unauthorizedError
func (e unauthorizedError) Is(tgt error) bool {
	_, ok := tgt.(unauthorizedError)
	return ok
}
