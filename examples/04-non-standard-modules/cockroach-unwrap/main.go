package main

import (
	"fmt"
	"net"

	"github.com/cockroachdb/errors"
)

type isTemporary interface {
	Temporary() bool
}

// IsTemporaryOnce считает цепочку ошибок временной, если хотя бы одна
// из ошибок в ней была временной.
func IsTemporaryOnce(err error) bool {
	for c := err; c != nil; c = errors.UnwrapOnce(c) {
		e, ok := c.(isTemporary)
		if ok && e.Temporary() {
			return true
		}
	}
	return false
}

// IsTemporary считает цепочку ошибок временной, только если
// корневая ошибка в ней была временной.
func IsTemporary(err error) bool {
	c := errors.UnwrapAll(err)
	e, ok := c.(isTemporary)
	return ok && e.Temporary()
}

func main() {
	dnsErr := &net.DNSError{IsTemporary: true} // Произошла сетевая ошибка.

	err := fmt.Errorf("second wrap: %w",
		fmt.Errorf("first wrap: %w", dnsErr))

	fmt.Printf("is temporary: %t\n", IsTemporaryOnce(err))     // true
	fmt.Printf("is temporary at root: %t\n", IsTemporary(err)) // true
}
