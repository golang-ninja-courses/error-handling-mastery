package main

import (
	"fmt"
	"net"

	"github.com/cockroachdb/errors"
)

type isTemporary interface {
	Temporary() bool
}

func IsTemporary(err error) (interface{}, bool) {
	e, ok := err.(isTemporary)
	return e, ok && e.Temporary()
}

func main() {
	var err error
	err = &net.DNSError{IsTemporary: true}   // Произошла сетевая ошибка
	err = fmt.Errorf("first wrap: %w", err)  // Оборачиваем раз
	err = fmt.Errorf("second wrap: %w", err) // Оборачиваем два

	e, ok := errors.If(err, IsTemporary) // Проверяем временная ли ошибка

	fmt.Printf("%T is temporary: %t", e, ok)
}
