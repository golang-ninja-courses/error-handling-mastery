package main

import (
	"fmt"
	"net"
)

func IsTemporary(err error) bool {
	// Приводим ошибку к интерфейсу, тем самым проверяя поведение.
	e, ok := err.(interface{ Temporary() bool })
	return ok && e.Temporary()
}

func networkRequest() error {
	return &net.DNSError{ // У *DNSError есть метод Temporary() bool, загляните внутрь.
		IsTimeout:   true,
		IsTemporary: true,
	}
}

func main() {
	if err := networkRequest(); err != nil {
		fmt.Println("error is temporary:", IsTemporary(err)) // error is temporary: true
	}
}
