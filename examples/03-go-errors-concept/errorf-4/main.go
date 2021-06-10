package main

import (
	"fmt"
)

func main() {
	err := fmt.Errorf("cannot do operation: %w", nil)
	fmt.Println(err)        // cannot do operation: %!w(<nil>)
	fmt.Println(err == nil) // false
}
