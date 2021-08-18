package main

import (
	"context"
	"fmt"

	"github.com/pkg/errors"
)

type FileLoadError struct {
	URL string
	Err error // Для хранения "родительской" ошибки.
}

func (p *FileLoadError) Error() string {
	// Текст "родительской ошибки" фигурирует в тексте этой ошибки.
	return fmt.Sprintf("%q: %v", p.URL, p.Err)
}

type File struct{}

func getFile(u string) (File, error) {
	return File{}, context.Canceled
}

func loadFiles(urls ...string) ([]File, error) {
	files := make([]File, len(urls))
	for i, url := range urls {
		f, err := getFile(url)
		if err != nil {
			return nil, errors.WithStack(&FileLoadError{url, err})
		}
		files[i] = f
	}
	return files, nil
}

func transfer() error {
	_, err := loadFiles("www.golang-courses.ru")
	if err != nil {
		return errors.WithMessage(err, "cannot load files")
	}

	// ...
	return nil
}

func handle() error {
	if err := transfer(); err != nil {
		return errors.WithMessage(err, "cannot transfer files")
	}

	// ...
	return nil
}

func main() {
	err := handle()
	fmt.Printf("%+v\n\n", err)

	if f, ok := errors.Cause(err).(*FileLoadError); ok {
		fmt.Println(f.URL)
	}

	var fileLoadErr *FileLoadError
	if err := handle(); errors.As(err, &fileLoadErr) {
		fmt.Println(fileLoadErr.URL) // www.golang-courses.ru
	}
}
