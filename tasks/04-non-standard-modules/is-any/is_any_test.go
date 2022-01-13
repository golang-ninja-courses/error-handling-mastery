package errors

import (
	"errors"
	"fmt"
	"io"
	"testing"

	"github.com/stretchr/testify/assert"
)

var errReference = errors.New("reference Is() error")

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
	errWithMessage := newWithMessageError("any message", errSecondWrap)

	for _, e := range []error{errSecondWrap, errFirstWrap, err} {
		assert.True(t, IsAny(errWithMessage, e))
	}

	for _, e := range []error{err, errFirstWrap, errSecondWrap} {
		assert.False(t, IsAny(e, errWithMessage))
	}
}

func TestIsAny_Is(t *testing.T) {
	err := errors.New("just an error")
	errFirstWrap := fmt.Errorf("first wrap: %w", err)
	errSecondWrap := fmt.Errorf("second wrap: %w", errFirstWrap)
	errWithMessage := newWithMessageError("any message", errSecondWrap)

	assert.True(t, IsAny(errWithMessage, io.EOF, errReference))
}

func newWithMessageError(m string, err error) *withMessageError {
	return &withMessageError{message: m, err: err}
}

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
