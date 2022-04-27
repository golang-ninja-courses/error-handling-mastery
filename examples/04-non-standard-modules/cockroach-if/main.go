package main

import (
	"fmt"
	"net"

	"github.com/cockroachdb/errors"
)

type isTemporary interface {
	Temporary() bool
}

func IsTemporary(err error) (any, bool) {
	e, ok := err.(isTemporary)
	return e, ok && e.Temporary() // Возвращаем не только флаг, но и ошибку.
}

func main() {
	dnsErr := &net.DNSError{IsTemporary: true} // Произошла сетевая ошибка.

	err := fmt.Errorf("second wrap: %w",
		fmt.Errorf("first wrap: %w", dnsErr))

	e, ok := errors.If(err, IsTemporary)
	fmt.Printf("%T is temporary: %t", e, ok) // *net.DNSError is temporary: true
}
