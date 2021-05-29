package main

import "fmt"

type MaxSizeExceededError struct {
	desiredLen int
}

func (e *MaxSizeExceededError) Error() string {
	return fmt.Sprintf("max allowed length 1024 is less than desired %d", e.desiredLen)
}

type BufferIsEmptyError struct {}

func (b *BufferIsEmptyError) Error() string {
	return "buffer is empty"
}

type ByteBuffer struct {
	buffer []byte // сам буфер, содержит какие-то данные
	offset int    // смещение, указывающее на первый непрочитанный байт
}

func (b *ByteBuffer) Write(p []byte) (int, error) {
	if len(b.buffer) + len(p) > 1024 {
		return 0, &MaxSizeExceededError{desiredLen: len(b.buffer) + len(p)}
	}

	b.buffer = append(b.buffer, p...)
	return len(p), nil
}

func (b *ByteBuffer) Read(p []byte) (int, error) {
	if len(b.buffer) == b.offset {
		return 0, &BufferIsEmptyError{}
	}

	n := copy(p, b.buffer[b.offset:])
	b.offset += n
	return n, nil
}

func main() {
	b := ByteBuffer{}
	b.Write([]byte("hello"))

	p := make([]byte, 1)
	for {
		b.Read(p)
		fmt.Print(string(p))
	}
}
