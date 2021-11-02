package rest

import (
	"errors"
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

func ExampleHandler() {
	if err := Handle(); err != nil {
		fmt.Println("handle err:", err)
	} else {
		fmt.Println("no handle err")
	}

	// Output:
	// no handle err
}

func TestHandle(t *testing.T) {
	t.Run("no error", func(t *testing.T) {
		assert.NoError(t, Handle())
	})

	t.Run("internal server error", func(t *testing.T) {
		defer func() {
			usefulWork = func() error { return nil }
		}()

		usefulWork = func() error {
			return errors.New("something wrong")
		}
		assert.ErrorIs(t, Handle(), ErrInternalServerError)
	})
}
