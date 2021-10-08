package ptrerror

import (
	"errors"
	"testing"

	"github.com/stretchr/testify/require"
)

func TestPointerError(t *testing.T) {
	p1 := NewPointerError("sky is falling")
	p2 := NewPointerError("sky is falling")
	require.False(t, errors.Is(p1, p2))
}
