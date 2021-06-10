package main

import (
	"fmt"
	"net"
)

type temporary interface {
	Temporary() bool
}

func IsTemporary(err error) bool {
	e, ok := err.(temporary) // Приводим ошибку к интерфейсу temporary, тем самым проверяя поведение
	return ok && e.Temporary()
}

func networkRequest() error {
	return &net.DNSError{ // У DNSError есть метод Temporary() bool, загляните внутрь
		IsTimeout:   true,
		IsTemporary: true,
	}
}

func main() {
	if err := networkRequest(); err != nil {
		if IsTemporary(err) {
			fmt.Printf("temporary error detected") // Будет выведено
		}
	}
}
