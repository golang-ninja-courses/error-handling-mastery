package main

import (
	"fmt"
	"io"
)

func main() {
	for _, err := range []error{
		fmt.Errorf("request failed: %v", io.EOF),
		fmt.Errorf("request failed: %v", io.EOF.Error()),
		fmt.Errorf("request failed: %s", io.EOF),
		fmt.Errorf("request failed: %s", io.EOF.Error()),
	} {
		fmt.Println(err)
	}
}
