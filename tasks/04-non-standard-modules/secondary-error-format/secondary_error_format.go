package secondary_error_format

type withSecondaryError struct {
	cause        error
	secondaryErr error
}

func (e *withSecondaryError) Error() string { return e.cause.Error() }

// TODO реализовать метод/методы форматирования для withSecondaryError

func CombineErrors(err, otherErr error) error {
	// TODO реализовать
	return nil
}
