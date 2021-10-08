package factorial

import (
	"strconv"
	"testing"

	"github.com/stretchr/testify/require"
)

func TestCalculate(t *testing.T) {
	t.Run("negative n", func(t *testing.T) {
		f, err := Calculate(-1)
		require.ErrorIs(t, err, ErrNegativeN)
		require.Equal(t, 0, f)
	})

	t.Run("too deep", func(t *testing.T) {
		f, err := Calculate(1e3)
		require.ErrorIs(t, err, ErrTooDeep)
		require.Contains(t, err.Error(), strconv.Itoa(maxDepth))
		require.Equal(t, 0, f)
	})

	t.Run("factorial of 0 is 1", func(t *testing.T) {
		f, err := Calculate(0)
		require.NoError(t, err)
		require.Equal(t, 1, f)
	})

	t.Run("factorial of 10 is 3628800", func(t *testing.T) {
		f, err := Calculate(10)
		require.NoError(t, err)
		require.Equal(t, 3628800, f)
	})
}
