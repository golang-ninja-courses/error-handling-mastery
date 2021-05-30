package main

type ByteBuffer struct {
	buffer []byte // сам буфер, содержит какие-то данные
	offset int    // смещение, указывающее на первый непрочитанный байт
}

// Реализовать интерфейсы io.ByteWriter и io.ByteReader в структуре ByteBuffer.
//
// Метод WriteByte() должен возвращать ошибку при попытке записать в буфер, если там уже больше 256 байт.
// Метод Error() возвращаемой ошибки должен выводить "max allowed length 256 is less than desired %d".
//
// Метод ReadByte() должен возвращать ошибку при попытке прочитать из буфера, если он пуст.
// Тип возвращаемой ошибки должен быть BufferIsEmptyError.
