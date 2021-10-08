package main

import (
	"context"
	"fmt"
)

func main() {
	fmt.Printf("cannot do operation: %w", context.Canceled)
	// cannot do operation: %!w(*errors.errorString=&{context canceled})
}
