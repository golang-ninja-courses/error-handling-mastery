package main

import "fmt"

const bufferMaxSize = 1024

type MaxSizeExceededError struct {
	desiredLen int
}

func (e *MaxSizeExceededError) Error() string {
	return fmt.Sprintf("buffer max size exceeded: %d > %d", e.desiredLen, bufferMaxSize)
}

type EndOfBuffer struct{}

func (b *EndOfBuffer) Error() string {
	return "end of buffer"
}

type ByteBuffer struct {
	// buffer представляем собой непосредственно буфер: содержит какие-то данные.
	buffer []byte
	// offset представляет собой смещение, указывающее на первый непрочитанный байт.
	offset int
}

func (b *ByteBuffer) Write(p []byte) (int, error) {
	if len(b.buffer)+len(p) > bufferMaxSize {
		return 0, &MaxSizeExceededError{desiredLen: len(b.buffer) + len(p)}
	}

	b.buffer = append(b.buffer, p...)
	return len(p), nil
}

func (b *ByteBuffer) Read(p []byte) (int, error) {
	if b.offset >= len(b.buffer) {
		return 0, new(EndOfBuffer)
	}

	n := copy(p, b.buffer[b.offset:])
	b.offset += n
	return n, nil
}

func main() {
	var b ByteBuffer
	if _, err := b.Write([]byte("hello hello hello")); err != nil {
		panic(err)
	}

	p := make([]byte, 3)
	for {
		n, err := b.Read(p)
		if _, ok := err.(*EndOfBuffer); ok {
			break
		}
		fmt.Print(string(p[:n])) // hello hello hello
	}
}
