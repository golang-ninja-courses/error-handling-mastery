package combine_errors

type withSecondaryErrors struct {
	cause          error
	additionalErrs []error
}

func (e *withSecondaryErrors) Error() string { return e.cause.Error() }

func (e *withSecondaryErrors) Unwrap() error { return e.cause }

func WithSecondaryErrors(err error, additionalErrs ...error) error {
	if err == nil || len(additionalErrs) == 0 {
		return err
	}

	return &withSecondaryErrors{cause: err, additionalErrs: additionalErrs}
}

func CombineErrors(err error, otherErrs ...error) error {
	// TODO реализовать
	return nil
}
