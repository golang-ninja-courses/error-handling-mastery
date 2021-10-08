package errors

import (
	"bytes"
	"fmt"
	"io"
	"testing"

	"github.com/pkg/errors"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func ExampleTrimStackTrace() {
	err := errors.Wrap(errors.WithStack(io.EOF), "read file")
	fmt.Printf("%+v", TrimStackTrace(err))

	// Output:
	// read file: EOF
}

func TestTrimStackTrace(t *testing.T) {
	cases := []struct {
		name string
		err  error
	}{
		{
			name: "pkg/errors",
			err:  errors.Wrapf(errors.Wrap(errors.WithStack(io.EOF), "message 1"), "message %d", 2),
		},
		{
			name: "std errors",
			err:  fmt.Errorf("message %d: %w", 2, fmt.Errorf("message 1: %w", io.EOF)),
		},
	}

	for _, tt := range cases {
		t.Run(tt.name, func(t *testing.T) {
			trimmedErr := TrimStackTrace(tt.err)

			b := bytes.NewBuffer(nil)
			_, err := fmt.Fprintf(b, "%+v", trimmedErr)
			require.NoError(t, err)
			assert.ErrorIs(t, trimmedErr, io.EOF)
			assert.Equal(t, "message 2: message 1: EOF", b.String())
		})
	}
}

func TestTrimStackTrace_Nil(t *testing.T) {
	err := TrimStackTrace(nil)
	assert.Nil(t, err)
}
