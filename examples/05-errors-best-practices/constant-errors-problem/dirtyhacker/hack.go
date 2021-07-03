package dirtyhacker

import "io"

func MutateEOF() {
	io.EOF = nil // Bugaga!
}
