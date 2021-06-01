package main

import (
	"io"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestByteBuffer_ImplementsNecessary(t *testing.T) {
	assert.Implements(t, (*io.ByteWriter)(nil), new(ByteBuffer))
	assert.Implements(t, (*io.ByteReader)(nil), new(ByteBuffer))
}

func TestByteBuffer_IoReader(t *testing.T) {
	b := ByteBuffer{}

	expected := "TestByteBuffer_IoReader"
	actual := ""

	b.buffer = []byte(expected)

	for {
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

	for {
		err := b.WriteByte('1')
		if err != nil {
			assert.EqualError(t, err, "max allowed length 256 is less than desired 257")
			break
		}
	}

	assert.Len(t, b.buffer, 256)
}
