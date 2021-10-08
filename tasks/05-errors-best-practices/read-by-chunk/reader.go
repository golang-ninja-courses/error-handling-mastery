package reader

import "io"

var ErrInvalidChunkSize error

func ReadByChunk(r io.Reader, chunkSize int) ([][]byte, error) {
	return nil, nil
}
