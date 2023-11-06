package errors

import (
	"errors"
	"fmt"
	"io"
	"os"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func ExampleCombine() {
	err := Combine(os.ErrExist, errors.New("ah shit, here we go again"))
	fmt.Println(err)
	fmt.Println()
	fmt.Printf("%+v", err)

	// Output:
	// file already exists
	//
	// file already exists
	//   - ah shit, here we go again
	//
}

func TestCombine(t *testing.T) {
	var (
		err  = errors.New("an error")
		err2 = errors.New("an error 2")
		err3 = errors.New("an error 3")
		err4 = errors.New("an error 4")
	)

	combinedErr := Combine(err, err2, err3, err4)
	require.NotNil(t, combinedErr)
	assert.ErrorIs(t, combinedErr, err)
	assert.NotErrorIs(t, combinedErr, err2)
	assert.NotErrorIs(t, combinedErr, err3)
	assert.NotErrorIs(t, combinedErr, err4)
}

func TestCombine_Format(t *testing.T) {
	var err error

	combinedErr := Combine(err, os.ErrExist)
	assertErrFormat(t, combinedErr, "file already exists")

	combinedErr = Combine(os.ErrExist, errors.New("ah shit, here we go again"))
	assertErrFormat(t, combinedErr, `file already exists
  - ah shit, here we go again
`)

	combinedErr = Combine(errors.New("an error"),
		errors.New("an error 2"),
		errors.New("an error 3"),
		errors.New("an error 4"))
	assertErrFormat(t, combinedErr, `an error
  - an error 2
  - an error 3
  - an error 4
`)

	combinedErr = Combine(combinedErr, combinedErr)
	assertErrFormat(t, combinedErr, `an error
  - an error
`)
}

func TestCombine_ArgsAreNil(t *testing.T) {
	var (
		err  error
		errs []error
	)

	combinedErr := Combine(err, errs...)
	assert.Nil(t, combinedErr)
}

func TestCombine_NoOtherErrors(t *testing.T) {
	err := errors.New("an error")

	combinedErr := Combine(err)
	require.NotNil(t, combinedErr)
	assert.ErrorIs(t, combinedErr, err)
}

func TestCombine_NilAmongOther(t *testing.T) {
	var (
		err  = errors.New("an error")
		err2 error
		err3 = errors.New("an error 3")
	)

	combinedErr := Combine(err, err2, err3)
	assert.NotNil(t, combinedErr)
	assert.ErrorIs(t, combinedErr, err)
	assert.NotErrorIs(t, combinedErr, err2)
	assert.NotErrorIs(t, combinedErr, err3)

	assertErrFormat(t, combinedErr, `an error
  - an error 3
`)
}

func TestCombine_ErrIsNil(t *testing.T) {
	var (
		err  error
		err2 = errors.New("an error 2")
		err3 = errors.New("an error 3")
		err4 = errors.New("an error 4")
	)

	combinedErr := Combine(err)
	assert.Nil(t, combinedErr)

	combinedErr = Combine(err, err2, err3, err4)
	assert.NotNil(t, combinedErr)
	assert.ErrorIs(t, combinedErr, err2)
	assert.NotErrorIs(t, combinedErr, err3)
	assert.NotErrorIs(t, combinedErr, err4)
}

func TestCombine_Immutable(t *testing.T) {
	errs := make([]error, 0, 10)
	errs = append(errs,
		errors.New("an error 2"),
		errors.New("an error 3"),
	)

	combinedErr := Combine(io.EOF, errs...)
	assertErrFormat(t, combinedErr, `EOF
  - an error 2
  - an error 3
`)

	errs = append(errs, errors.New("an error 4")) //nolint:ineffassign,staticcheck,wastedassign
	assertErrFormat(t, combinedErr, `EOF
  - an error 2
  - an error 3
`)
}

func assertErrFormat(t *testing.T, err error, expected string) {
	t.Helper()
	assert.Equal(t, expected, fmt.Sprintf("%+v", err))
}
