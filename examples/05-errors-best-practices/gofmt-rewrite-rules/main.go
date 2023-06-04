package main

import (
	"fmt"
	"io"
)

func main() {
	if err := work(); nil != err {
		fmt.Println("error")
	}

	if err := work(); err == nil {
		fmt.Println("no error")
	}
}

func work() error {
	return io.EOF
}
