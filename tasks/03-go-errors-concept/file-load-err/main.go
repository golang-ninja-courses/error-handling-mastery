package main

import (
	"context"
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

type File struct{}

func getFile(u string) (File, error) {
	return File{}, context.Canceled
}

func loadFiles(urls ...string) ([]File, error) {
	files := make([]File, len(urls))
	for i, url := range urls {
		f, err := getFile(url)
		if err != nil {
			return nil, &FileLoadError{url, err} // <- Врапим ошибку загрузки в *FileLoadError.
		}
		files[i] = f
	}
	return files, nil
}

func transfer() error {
	_, err := loadFiles("www.golang-ninja.ru")
	if err != nil {
		return fmt.Errorf("cannot load files: %v", err)
	}

	// ...
	return nil
}

func main() {
	if err := transfer(); err != nil {
		if _, ok := err.(*FileLoadError); ok {
			fmt.Println("file load err received")
		} else {
			fmt.Println("unexpected error")
		}
	}
}
