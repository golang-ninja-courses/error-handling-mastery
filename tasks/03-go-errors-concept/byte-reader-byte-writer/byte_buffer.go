package bytebuffer

const bufferMaxSize = 1024

type MaxSizeExceededError struct {
	desiredLen int
}

type EndOfBuffer struct{}

type ByteBuffer struct {
	buffer []byte
	offset int
}

// Необходимо сделать так, чтобы тип *ByteBuffer реализовывал интерфейсы io.ByteWriter и io.ByteReader.
//
// Метод WriteByte должен возвращать ошибку *MaxSizeExceededError при попытке записи в буфер,
// если в нём уже больше bufferMaxSize байт.
//
// Метод ReadByte должен возвращать ошибку *EndOfBuffer при попытке чтения из буфера,
// если ранее буфер уже был вычитан полностью.
