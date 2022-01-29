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
			// If some data is available but not len(p) bytes, Read conventionally
			// returns what is available instead of waiting for more.
			name:      "greedy reader giving out by 5 bytes to 7 bytes chunk",
			r:         newGreedyReader(s, 5),
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

func Test_greedyReader(t *testing.T) {
	const (
		wantToRead      = 30
		maxBytesPerRead = 10
	)

	r := newGreedyReader("Errors are values.", maxBytesPerRead)
	{
		data := make([]byte, wantToRead)
		n, err := r.Read(data)
		assert.NoError(t, err)
		assert.Equal(t, 10, n)
		assert.Equal(t, "Errors are", string(data[:n]))
	}
	{
		data := make([]byte, wantToRead)
		n, err := r.Read(data)
		assert.NoError(t, err)
		assert.Equal(t, 8, n)
		assert.Equal(t, " values.", string(data[:n]))
	}
	{
		data := make([]byte, wantToRead)
		n, err := r.Read(data)
		assert.ErrorIs(t, err, io.EOF)
		assert.Equal(t, 0, n)
		assert.Equal(t, make([]byte, wantToRead), data)
	}
}

type errReader struct {
	err error
}

func (r errReader) Read([]byte) (int, error) {
	return 0, r.err
}

// greedyReader представляет собой жадного читателя,
// т.е. выдающего не более чем maxBytesPerRead байт за раз.
type greedyReader struct {
	data            []byte
	maxBytesPerRead int
	pos             int
}

func newGreedyReader(str string, maxPerRead int) *greedyReader {
	return &greedyReader{data: []byte(str), maxBytesPerRead: maxPerRead}
}

func (s *greedyReader) Read(p []byte) (int, error) {
	if s.pos >= len(s.data) {
		return 0, io.EOF
	}

	right := s.pos + s.maxBytesPerRead
	if right > len(s.data) {
		right = len(s.data)
	}

	n := copy(p, s.data[s.pos:right])
	s.pos += n

	return n, nil
}

func debugChunks(chunks [][]byte) string {
	return string(bytes.Join(chunks, []byte{'_'}))
}
