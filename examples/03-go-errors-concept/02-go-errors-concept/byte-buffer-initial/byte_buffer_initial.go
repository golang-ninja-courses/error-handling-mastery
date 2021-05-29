package main

import "fmt"

type ByteBuffer struct {
	buffer []byte // сам буфер, содержит какие-то данные
	offset int    // смещение, указывающее на первый непрочитанный байт
}

func (b *ByteBuffer) Write(p []byte) int {
	b.buffer = append(b.buffer, p...)
	return len(p)
}

func (b *ByteBuffer) Read(p []byte) int {
	n := copy(p, b.buffer[b.offset:])
	b.offset += n
	return n
}

func main() {
	b := ByteBuffer{}
	b.Write([]byte("hello"))

	p := make([]byte, 1)
	for b.Read(p) != 0 {
		fmt.Print(string(p))
	}
}
