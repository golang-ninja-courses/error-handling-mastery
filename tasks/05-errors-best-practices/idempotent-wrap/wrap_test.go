package errors

import (
	"fmt"
	"io"
	"testing"

	"github.com/pkg/errors"
	"github.com/stretchr/testify/assert"
)

func ExampleWrap() {
	err := Wrap(Wrap(Wrap(io.EOF, "wrap 0"), "wrap 1"), "wrap 2")
	fmt.Printf("%+v", err)
}

func TestWrap(t *testing.T) {
	const wraps = 100

	t.Run("single wrap", func(t *testing.T) {
		err := Wrap(io.EOF, "wrap 0")
		traces := getTracesCount(err)
		assert.Equal(t, 1, traces)
	})

	t.Run("use wrap only", func(t *testing.T) {
		err := io.EOF
		for i := 0; i < wraps; i++ {
			err = Wrap(err, fmt.Sprintf("wrap %d", i))
			if i == 0 {
				assert.Implements(t, (*stackTracer)(nil), err)
			}
		}

		traces := getTracesCount(err)
		assert.Equal(t, 1, traces)
	})

	t.Run("wrapped error already has stack", func(t *testing.T) {
		err := errors.WithStack(io.EOF)
		for i := 0; i < wraps; i++ {
			err = Wrap(err, fmt.Sprintf("wrap %d", i))
		}

		traces := getTracesCount(err)
		assert.Equal(t, 1, traces)
	})
}

func getTracesCount(err error) (n int) {
	for e := err; e != nil; e = errors.Unwrap(e) {
		if _, ok := e.(stackTracer); ok {
			n++
		}
	}
	return n
}
