package is_any

import (
	"errors"
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

var errReference = errors.New("reference Is() error")

type withMessageError struct {
	message string
	err     error
}

func (e *withMessageError) Error() string {
	return e.message
}

func (e *withMessageError) Unwrap() error {
	return e.err
}

func (e *withMessageError) Is(target error) bool {
	return target == errReference
}

func TestIsAny(t *testing.T) {
	err := errors.New("just an error")
	errSecond := errors.New("second error")
	errThird := errors.New("third error")

	assert.False(t, IsAny(err, errSecond, errThird))
	assert.True(t, IsAny(err, errSecond, errThird, err))
}

func TestIsAny_Wrap(t *testing.T) {
	err := errors.New("just an error")
	errFirstWrap := fmt.Errorf("first wrap: %w", err)
	errSecondWrap := fmt.Errorf("second wrap: %w", errFirstWrap)
	errWithMessage := &withMessageError{
		message: "any message",
		err:     errSecondWrap,
	}

	for _, e := range []error{errSecondWrap, errFirstWrap, err} {
		assert.True(t, IsAny(errWithMessage, e))
	}

	assert.False(t, IsAny(err, errWithMessage))
}

func TestIsAny_Is(t *testing.T) {
	err := errors.New("just an error")
	errFirstWrap := fmt.Errorf("first wrap: %w", err)
	errSecondWrap := fmt.Errorf("second wrap: %w", errFirstWrap)
	errWithMessage := &withMessageError{
		message: "any message",
		err:     errSecondWrap,
	}

	assert.True(t, IsAny(errWithMessage, errReference))
}
