package main

import "fmt"

type ByteBuffer struct {
	// buffer представляем собой непосредственно буфер: содержит какие-то данные.
	buffer []byte
	// offset представляет собой смещение, указывающее на первый непрочитанный байт.
	offset int
}

func (b *ByteBuffer) Write(p []byte) int {
	b.buffer = append(b.buffer, p...)
	return len(p)
}

func (b *ByteBuffer) Read(p []byte) int {
	if b.offset >= len(b.buffer) {
		return 0
	}

	n := copy(p, b.buffer[b.offset:])
	b.offset += n
	return n
}

func main() {
	var b ByteBuffer
	b.Write([]byte("hello hello hello"))

	p := make([]byte, 3)
	for {
		n := b.Read(p)
		if n == 0 {
			break
		}
		fmt.Print(string(p[:n])) // hello hello hello
	}
}
