package errors

type NotFoundError struct {
	message string
	err     error
}

func (e NotFoundError) Error() string {
	return e.message
}

func (e NotFoundError) Unwrap() error {
	return e.err
}

func NewNotFoundError(message string, err error) error {
	return NotFoundError{message: message, err: err}
}
