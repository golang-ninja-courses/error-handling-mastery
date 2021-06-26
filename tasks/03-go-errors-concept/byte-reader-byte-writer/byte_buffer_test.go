package bytebuffer

import (
	"io"
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestByteBuffer_implementsNecessary(t *testing.T) {
	assert.Implements(t, (*io.ByteWriter)(nil), new(ByteBuffer))
	assert.Implements(t, (*io.ByteReader)(nil), new(ByteBuffer))
}

func TestByteBuffer_ioReader(t *testing.T) {
	var b ByteBuffer

	expected := "TestByteBuffer_IoReader"
	var actual strings.Builder

	for _, c := range []byte(expected) {
		err := b.WriteByte(c)
		require.NoError(t, err)
	}

	for i := 0; i < len(expected)+1; i++ {
		bb, err := b.ReadByte()
		if err != nil {
			if isEndOfBuffer(err) {
				break
			}
			require.NoError(t, err)
		}

		actual.WriteByte(bb)
	}

	assert.Equal(t, expected, actual.String())
}

func TestByteBuffer_ioWriter(t *testing.T) {
	var b ByteBuffer

	for i := 0; i < bufferMaxSize+1; i++ {
		err := b.WriteByte('1')
		if err != nil {
			if isMaxSizeExceededError(err) {
				break
			}
			require.NoError(t, err)
		}
	}

	assert.Len(t, b.buffer, bufferMaxSize)
}

func TestByteBuffer_readFromEmptyBuffer(t *testing.T) {
	var b ByteBuffer
	n, err := b.ReadByte()
	assert.Equal(t, byte(0), n)
	assert.True(t, isEndOfBuffer(err))
}

func TestEndOfBuffer_Error(t *testing.T) {
	assert.NotEmpty(t, new(EndOfBuffer).Error())
}

func TestMaxSizeExceededError_Error(t *testing.T) {
	assert.NotEmpty(t, new(MaxSizeExceededError).Error())
}

func isEndOfBuffer(err error) bool {
	_, ok := err.(*EndOfBuffer)
	return ok
}

func isMaxSizeExceededError(err error) bool {
	_, ok := err.(*MaxSizeExceededError)
	return ok
}
