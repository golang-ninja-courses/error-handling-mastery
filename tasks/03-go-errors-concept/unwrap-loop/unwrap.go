package errs

type Unwrapper interface {
	Unwrap() error
}

func Unwrap(err error) error {
	// Реализуй меня.
	return nil
}
