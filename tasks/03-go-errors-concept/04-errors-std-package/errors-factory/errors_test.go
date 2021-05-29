package errors

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestNewError(t *testing.T) {
	for _, s := range []string{
		"context canceled",
		"end of file",
		"invalid token",
	} {
		err := NewError(s)
		require.Implements(t, (*error)(nil), err)
		require.Equal(t, s, err.Error())
	}
}

func TestNewError_equality(t *testing.T) {
	lhs := NewError("invalid token")
	rhs := NewError("invalid token")
	require.False(t, lhs == rhs, "different errors with the same text must not be equal")
}
