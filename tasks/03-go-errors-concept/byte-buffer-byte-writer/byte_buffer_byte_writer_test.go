package main

import (
	"io"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestByteBuffer_ImplementsNecessary(t *testing.T) {
	assert.Implements(t, (*io.ByteWriter)(nil), new(ByteBuffer))
	assert.Implements(t, (*io.ByteReader)(nil), new(ByteBuffer))
}

func TestByteBuffer_IoReader(t *testing.T) {
	b := ByteBuffer{}

	expected := "TestByteBuffer_IoReader"
	actual := ""

	for _, c := range []byte(expected) {
		err := b.WriteByte(c)
		require.NoError(t, err)
	}

	for i := 0; i < len(expected)+1; i++ {
		bb, err := b.ReadByte()
		if err != nil {
			assert.ErrorIs(t, err, new(BufferIsEmptyError))
			break
		}

		actual += string(bb)
	}

	assert.Equal(t, expected, actual)
}

func TestByteBuffer_IoWriter(t *testing.T) {
	b := ByteBuffer{}

	errMaxSizeExceeded := new(MaxSizeExceededError)

	for i := 0; i < bufferMaxSize+1; i++ {
		err := b.WriteByte('1')
		if err != nil {
			assert.ErrorAs(t, err, &errMaxSizeExceeded)
			break
		}
	}

	assert.Len(t, b.buffer, bufferMaxSize)
}
