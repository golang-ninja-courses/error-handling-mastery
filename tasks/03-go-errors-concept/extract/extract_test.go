package errs

import (
	"context"
	"errors"
	"fmt"
	"io"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestExtract(t *testing.T) {
	errInvalidCursor := errors.New("invalid cursor")
	vWrapped := fmt.Errorf("parse request: %v", errInvalidCursor)

	cases := []struct {
		in       error
		expected []error
	}{
		{
			in:       errInvalidCursor,
			expected: []error{errInvalidCursor},
		},
		{
			in:       vWrapped,
			expected: []error{vWrapped},
		},
		{
			in:       fmt.Errorf("parse request: %w", errInvalidCursor),
			expected: []error{errInvalidCursor},
		},
		{
			in:       fmt.Errorf("parse request: %w", fmt.Errorf("decode cursor: %w", errInvalidCursor)),
			expected: []error{errInvalidCursor},
		},
		{
			in:       fmt.Errorf("parse request: %w: %w", errInvalidCursor, io.EOF),
			expected: []error{errInvalidCursor, io.EOF},
		},
		{
			in: fmt.Errorf("parse request: %w: %w: %w: %w",
				context.Canceled,
				fmt.Errorf("%w", errInvalidCursor),
				fmt.Errorf("another yet error: %w: %w", io.EOF, context.DeadlineExceeded),
				nil,
			),
			expected: []error{
				context.Canceled,
				errInvalidCursor,
				io.EOF,
				context.DeadlineExceeded,
			},
		},
	}

	for _, tt := range cases {
		t.Run("", func(t *testing.T) {
			chain := Extract(tt.in)
			assert.Equal(t, tt.expected, chain)
		})
	}
}
