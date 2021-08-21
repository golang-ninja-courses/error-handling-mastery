package errors

// Wrapf работает аналогично fmt.Errorf, только поддерживает nil-ошибки.
func Wrapf(err error, f string, v ...interface{}) error {
	return nil
}
