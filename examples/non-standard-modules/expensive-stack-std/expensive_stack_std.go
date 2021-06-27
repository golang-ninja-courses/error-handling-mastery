package main

import (
	"errors"
	"fmt"
)

func GimmeError() error {
	return errors.New("ooops, an error") // просто возвращает ошибку с текстом
}

func main() {
	fmt.Printf("%+v", GimmeError())
}
