package main

import (
	"fmt"

	"github.com/pkg/errors"
)

func GimmeError() error {
	return errors.New("ooops, an error") // возвращает ошибку с текстом и стектрейсом
}

func main() {
	fmt.Printf("%+v", GimmeError())
}
