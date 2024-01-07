package multierror

import (
	"errors"
	"fmt"
	"testing"

	"github.com/hashicorp/go-multierror"
	"github.com/stretchr/testify/assert"
)

func ExampleMap() {
	err := Map(func(v int) error {
		if v%2 != 0 {
			return errors.New("unsupported value")
		}
		return nil
	}, []int{1, 2, 3})

	err.(*multierror.Error).ErrorFormat = StdFormat
	fmt.Println(err)

	// Output:
	// elem at 0: unsupported value
	// elem at 2: unsupported value
}

func ExampleMapV2() {
	err := MapV2(func(v int) error {
		if v%2 != 0 {
			return errors.New("unsupported value")
		}
		return nil
	}, []int{1, 2, 3})
	fmt.Println(err)

	// Output:
	// elem at 0: unsupported value
	// elem at 2: unsupported value
}

func TestMap(t *testing.T) {
	cases := []struct {
		name           string
		in             []int
		fn             func(v int) error
		expectedErrors []string
	}{
		{
			name: "no errors",
			in:   []int{1, 2, 3},
			fn: func(v int) error {
				return nil
			},
			expectedErrors: []string{},
		},
		{
			name: "single error",
			in:   []int{1, 2, 3},
			fn: func(v int) error {
				if v == 2 {
					return errors.New("unsupported value")
				}
				return nil
			},
			expectedErrors: []string{
				"elem at 1: unsupported value",
			},
		},
		{
			name: "errors everywhere",
			in:   []int{1, 2, 3},
			fn: func(v int) error {
				return errors.New("unsupported value")
			},
			expectedErrors: []string{
				"elem at 0: unsupported value",
				"elem at 1: unsupported value",
				"elem at 2: unsupported value",
			},
		},
	}

	for _, tt := range cases {
		t.Run(tt.name, func(t *testing.T) {
			t.Run("v1", func(t *testing.T) {
				err := Map(tt.fn, tt.in)
				assert.Equal(t, tt.expectedErrors, extractErrorsStrings(err))
			})

			t.Run("v2", func(t *testing.T) {
				err := MapV2(tt.fn, tt.in)
				assert.Equal(t, tt.expectedErrors, extractErrorsStrings(err))
			})
		})
	}
}

func extractErrorsStrings(err error) []string {
	errs := extractErrors(err)

	result := make([]string, 0, len(errs))
	for _, err := range errs {
		result = append(result, err.Error())
	}
	return result
}

func extractErrors(err error) []error {
	if err == nil {
		return nil
	}

	switch v := err.(type) {
	case interface{ WrappedErrors() []error }:
		return v.WrappedErrors()
	case interface{ Unwrap() []error }:
		return v.Unwrap()
	}
	return []error{err}
}

var StdFormat = multierror.ErrorFormatFunc(func(errs []error) string {
	var b []byte
	for i, err := range errs {
		if i > 0 {
			b = append(b, '\n')
		}
		b = append(b, err.Error()...)
	}
	return string(b)
})
