package errors

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestNewError(t *testing.T) {
	for _, s := range []string{
		"context canceled",
		"end of file",
		"invalid token",
	} {
		err := NewError(s)
		assert.Equal(t, s, err.Error())
	}
}

func TestNewErrorEquality(t *testing.T) {
	lhs := NewError("invalid token")
	rhs := NewError("invalid token")
	assert.False(t, lhs == rhs, "different errors with the same text must not be equal")
}
