package reader

import (
	"bytes"
	"errors"
	"io"
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestReadByChunk(t *testing.T) {
	s := "Errors are values."
	errSkyIsFalling := errors.New("sky is falling")

	cases := []struct {
		name           string
		r              io.Reader
		chunkSize      int
		expectedErr    error
		expectedChunks [][]byte
	}{
		{
			name:           "nil reader",
			r:              nil,
			chunkSize:      10,
			expectedErr:    nil,
			expectedChunks: nil,
		},
		{
			name:        "zero chunk size",
			r:           strings.NewReader(s),
			chunkSize:   0,
			expectedErr: ErrInvalidChunkSize,
		},
		{
			name:        "negative chunk size",
			r:           strings.NewReader(s),
			chunkSize:   -1,
			expectedErr: ErrInvalidChunkSize,
		},
		{
			name:      "by byte",
			r:         strings.NewReader(s),
			chunkSize: 1,
			expectedChunks: [][]byte{
				{'E'}, {'r'}, {'r'}, {'o'}, {'r'}, {'s'}, {' '}, {'a'}, {'r'}, {'e'}, {' '}, {'v'}, {'a'}, {'l'}, {'u'}, {'e'}, {'s'}, {'.'},
			},
		},
		{
			name:      "by 3 bytes",
			r:         strings.NewReader(s),
			chunkSize: 3,
			expectedChunks: [][]byte{
				{'E', 'r', 'r'}, {'o', 'r', 's'}, {' ', 'a', 'r'}, {'e', ' ', 'v'}, {'a', 'l', 'u'}, {'e', 's', '.'},
			},
		},
		{
			name:      "by 5 bytes",
			r:         strings.NewReader(s),
			chunkSize: 5,
			expectedChunks: [][]byte{
				{'E', 'r', 'r', 'o', 'r'}, {'s', ' ', 'a', 'r', 'e'}, {' ', 'v', 'a', 'l', 'u'}, {'e', 's', '.'},
			},
		},
		{
			name:      "by 7 bytes",
			r:         strings.NewReader(s),
			chunkSize: 7,
			expectedChunks: [][]byte{
				{'E', 'r', 'r', 'o', 'r', 's', ' '}, {'a', 'r', 'e', ' ', 'v', 'a', 'l'}, {'u', 'e', 's', '.'},
			},
		},
		{
			name:      "one chunk equals to source",
			r:         strings.NewReader(s),
			chunkSize: len(s),
			expectedChunks: [][]byte{
				[]byte(s),
			},
		},
		{
			name:      "one chunk bigger than source",
			r:         strings.NewReader(s),
			chunkSize: len(s) * 2,
			expectedChunks: [][]byte{
				[]byte(s),
			},
		},
		{
			name:           "unexpected io error",
			r:              errReader{errSkyIsFalling},
			chunkSize:      10,
			expectedErr:    errSkyIsFalling,
			expectedChunks: nil,
		},
	}

	for _, tt := range cases {
		t.Run(tt.name, func(t *testing.T) {
			chunks, err := ReadByChunk(tt.r, tt.chunkSize)
			require.ErrorIs(t, err, tt.expectedErr)
			assert.Equal(t, tt.expectedChunks, chunks, debugChunks(chunks))
		})
	}
}

type errReader struct {
	err error
}

func (r errReader) Read([]byte) (int, error) {
	return 0, r.err
}

func debugChunks(chunks [][]byte) string {
	return string(bytes.Join(chunks, []byte{'_'}))
}
