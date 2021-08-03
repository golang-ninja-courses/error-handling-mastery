package main

import (
	"fmt"
	"net"

	"github.com/cockroachdb/errors"
)

type isTemporary interface {
	Temporary() bool
}

// Любая ошибка в цепочке, которая реализует isTemporary, приведет к тому,
// что вся цепочка будет признана временной.
func IsTemporaryOnce(err error) bool {
	for c := err; c != nil; c = errors.UnwrapOnce(c) {
		e, ok := c.(isTemporary)
		if ok && e.Temporary() {
			return true
		}
	}

	return false
}

// Если первая ошибка в цепочке реализует isTemporary, то считаем ошибку временной.
func IsTemporary(err error) bool {
	c := errors.UnwrapAll(err)
	e, ok := c.(isTemporary)
	if ok && e.Temporary() {
		return true
	}

	return false
}

func main() {
	var err error
	err = &net.DNSError{IsTemporary: true}   // Произошла сетевая ошибка
	err = fmt.Errorf("first wrap: %w", err)  // Оборачиваем раз
	err = fmt.Errorf("second wrap: %w", err) // Оборачиваем два

	ok := IsTemporaryOnce(err) // Проверяем временная ли ошибка

	fmt.Printf("is temporary: %t\n", ok)

	ok = IsTemporary(err)

	fmt.Printf("is temporary: %t\n", ok)
}
