package errs

import (
	"time"
)

type ErrorWithTime struct { // Реализуй меня.
}

var NewErrorWithTime = func(err error) error {
	return newErrorWithTime(err, time.Now)
}

func newErrorWithTime(err error, timeFunc func() time.Time) error {
	// Реализуй меня.
	return nil
}
