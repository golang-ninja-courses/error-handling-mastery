package main

import (
	"context"
	"errors"
	"fmt"
)

type FileLoadError struct {
	URL string
	Err error // Для хранения "родительской" ошибки.
}

func (p *FileLoadError) Error() string {
	// Текст "родительской ошибки" фигурирует в тексте этой ошибки.
	return fmt.Sprintf("%q: %v", p.URL, p.Err)
}

func (p *FileLoadError) Unwrap() error {
	return p.Err
}

func main() {
	fmt.Println(
		errors.Unwrap(&FileLoadError{Err: context.Canceled}))
}
