package errors

import (
	"io"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestWrapf(t *testing.T) {
	t.Run("positive", func(t *testing.T) {
		err := Wrapf(io.EOF, "user %d: cannot read file %q", 42, "/etc/hosts")
		require.Error(t, err)
		assert.EqualError(t, err, `user 42: cannot read file "/etc/hosts": EOF`)
		assert.ErrorIs(t, err, io.EOF)
	})

	t.Run("nil wrapping", func(t *testing.T) {
		err := Wrapf(nil, "some message")
		assert.Nil(t, err)
	})
}
