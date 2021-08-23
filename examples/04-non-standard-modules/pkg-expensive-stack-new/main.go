package main

import (
	stderrors "errors"
	"fmt"

	"github.com/pkg/errors"
)

func GimmeError() error {
	return stderrors.New("ooops, an error") // Просто возвращает ошибку с текстом.
}

func GimmePkgError() error {
	return errors.New("ooops, an error") // Возвращает ошибку вместе с текстом и стектрейсом.
}

func main() {
	fmt.Printf("%+v", GimmePkgError())
	fmt.Println("\n---")
	fmt.Printf("%+v", GimmeError())
}
