package main

import (
	"context"
	"errors"
	"fmt"
)

func main() {
	err := fmt.Errorf("cannot do operation: %w", context.Canceled.Error())
	fmt.Println(err)                              // cannot do operation: %!w(string=context canceled)
	fmt.Println(errors.Is(err, context.Canceled)) // false
}
