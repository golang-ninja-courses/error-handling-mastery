package pipeline

import (
	"errors"
	"strconv"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func Test_mergeErrors(t *testing.T) {
	t.Run("positive scenario", func(t *testing.T) {
		const n = 10
		const m = 100

		newError := func(i, j int) error {
			return errors.New(strconv.Itoa(i * j))
		}

		expectedErrors := make([]error, 0, n*m)
		for i := 0; i < n; i++ {
			for j := 0; j < m; j++ {
				expectedErrors = append(expectedErrors, newError(i, j))
			}
		}

		errcs := make([]<-chan error, n)
		for i := 0; i < n; i++ {
			i := i
			ch := make(chan error)
			go func() {
				defer close(ch)
				for j := 0; j < m; j++ {
					ch <- newError(i, j)
					time.Sleep(time.Millisecond)
				}
			}()
			errcs[i] = ch
		}

		errc, err := mergeErrors(errcs...)
		require.NoError(t, err)

		receivedErrors := make([]error, 0, len(expectedErrors))
		for err := range errc {
			receivedErrors = append(receivedErrors, err)
		}

		assert.ElementsMatch(t, expectedErrors, receivedErrors)
	})

	t.Run("empty input", func(t *testing.T) {
		errc, err := mergeErrors()
		require.ErrorIs(t, err, errEmptyInput)
		assert.Nil(t, errc)
	})

	t.Run("nil channel in input", func(t *testing.T) {
		errc, err := mergeErrors(
			make(chan error),
			nil,
			make(chan error),
		)
		require.ErrorIs(t, err, errNilChannel)
		assert.Nil(t, errc)
	})
}
