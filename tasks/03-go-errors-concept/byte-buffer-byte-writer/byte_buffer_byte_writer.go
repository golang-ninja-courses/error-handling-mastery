package main

import "io"

const (
	bufferMaxSize = 256
)

type BufferIsEmptyError struct {
	error
}

type MaxSizeExceededError struct {}

type ByteBuffer struct {
	io.ByteReader
	io.ByteWriter
	buffer []byte // сам буфер, содержит какие-то данные
	offset int    // смещение, указывающее на первый непрочитанный байт
}

// Реализовать интерфейсы io.ByteWriter и io.ByteReader в структуре ByteBuffer.
//
// Метод WriteByte() должен возвращать ошибку *MaxSizeExceededError при попытке записать в буфер,
// если там уже больше 256 байт. Метод *MaxSizeExceededError Error() должен выводить
// "max allowed length %d is less than desired %d", где первый %d – bufferMaxSize, второй желаемое число байт в буфере.
//
// Метод ReadByte() должен возвращать ошибку при попытке прочитать из буфера, если он пуст.
// Тип возвращаемой ошибки должен быть *BufferIsEmptyError.
