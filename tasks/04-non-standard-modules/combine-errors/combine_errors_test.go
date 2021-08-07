package combine_errors

import (
	"errors"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestCombineErrors(t *testing.T) {
	var (
		err              = errors.New("an error")
		err2             = errors.New("an error 2")
		err3             = errors.New("an error 3")
		errs             = []error{err2, err3}
		withSecondaryErr = &withSecondaryErrors{}
	)

	combinedErr := CombineErrors(err, errs...)

	assert.NotNil(t, combinedErr)
	assert.ErrorIs(t, combinedErr, err)
	assert.ErrorAs(t, combinedErr, &withSecondaryErr)
	assert.Len(t, withSecondaryErr.additionalErrs, 2)
}

func TestCombineErrors_ArgsAreNil(t *testing.T) {
	var (
		err  error
		errs []error
	)

	combinedErr := CombineErrors(err, errs...)

	assert.Nil(t, combinedErr)
}

func TestCombineErrors_OtherErrsAreNil(t *testing.T) {
	err := errors.New("an error")

	combinedErr := CombineErrors(err, nil)

	assert.NotNil(t, combinedErr)
	assert.ErrorIs(t, combinedErr, err)
}

func TestCombineErrors_ErrIsNil(t *testing.T) {
	var (
		err              = errors.New("an error")
		err2             = errors.New("an error 2")
		errs             = []error{err}
		withSecondaryErr = &withSecondaryErrors{}
	)

	combinedErr := CombineErrors(nil, errs...)

	assert.NotNil(t, combinedErr)
	assert.ErrorIs(t, combinedErr, err)
	assert.False(t, errors.As(combinedErr, &withSecondaryErr))

	errs = append(errs, err2)

	combinedErr = CombineErrors(nil, errs...)
	assert.NotNil(t, combinedErr)
	assert.ErrorIs(t, combinedErr, err)
	assert.ErrorAs(t, combinedErr, &withSecondaryErr)
	assert.Len(t, withSecondaryErr.additionalErrs, 1)
	assert.ErrorIs(t, withSecondaryErr.additionalErrs[0], err2)
}
