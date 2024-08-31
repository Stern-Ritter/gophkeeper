package errors

type ForbiddenError struct {
	message string
	err     error
}

func (e ForbiddenError) Error() string {
	return e.message
}

func (e ForbiddenError) Unwrap() error {
	return e.err
}

func NewForbiddenError(message string, err error) error {
	return ForbiddenError{message: message, err: err}
}
