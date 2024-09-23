package bytebuffer

import (
	"fmt"
)

const bufferMaxSize = 1024
type MaxSizeExceededError struct {
	desiredLen int
}
type EndOfBufferError struct{}
type ByteReader interface {
    ReadByte() (byte, error)
}
type ByteWriter interface {
    WriteByte(c byte) error
}
type ByteBuffer struct {
	buffer []byte
	offset int
}
func (e *MaxSizeExceededError) Error() string {
    return fmt.Sprintf("buffer max size exceeded: %d > %d", e.desiredLen, bufferMaxSize)
}
func (b *EndOfBufferError) Error() string {
    return "end of buffer"
}
func (b *ByteBuffer) WriteByte(p byte) error {
    if len(b.buffer) >= bufferMaxSize {
        return &MaxSizeExceededError{desiredLen: len(b.buffer)}
    }
    b.buffer = append(b.buffer, p)
    return nil
}
func (b *ByteBuffer) ReadByte() (byte, error) {
    if len(b.buffer) <= 0 {
        return 0, new(EndOfBufferError)
    }
    res := b.buffer[0]
    b.buffer = b.buffer[1:]
    return res, nil
}
// Необходимо сделать так, чтобы тип *ByteBuffer реализовывал интерфейсы io.ByteWriter и io.ByteReader.
// Метод WriteByte должен возвращать ошибку *MaxSizeExceededError при попытке записи в буфер,
// если в нём уже больше bufferMaxSize байт.
// Метод ReadByte должен возвращать ошибку *EndOfBufferError при попытке чтения из буфера,
// если ранее буфер уже был вычитан полностью.
// Необходимо сделать так, чтобы тип *ByteBuffer реализовывал интерфейсы io.ByteWriter и io.ByteReader.
// Метод WriteByte должен возвращать ошибку *MaxSizeExceededError при попытке записи в буфер,
// если в нём уже больше bufferMaxSize байт.
// Метод ReadByte должен возвращать ошибку *EndOfBufferError при попытке чтения из буфера,
// если ранее буфер уже был вычитан полностью.