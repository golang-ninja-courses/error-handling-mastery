package main

import (
	"errors"
	"fmt"
)

func GimmeDeepError(depth int) error {
	var err error
	if depth != 0 {
		err = GimmeDeepError(depth - 1)
		return fmt.Errorf("error happened on level %d: %w", depth-1, err) // по-другому оборачиваем ошибку
	}
	return errors.New("ooops, an error")
}

func main() {
	fmt.Printf("%+v", GimmeDeepError(2))
}