package errs

import (
	"time"
)

type WithTimeError struct { // Реализуй меня.
}

var NewWithTimeError = func(err error) error {
	return newWithTimeError(err, time.Now)
}

func newWithTimeError(err error, timeFunc func() time.Time) error {
	// Реализуй меня.
	return nil
}
