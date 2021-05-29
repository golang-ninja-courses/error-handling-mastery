package main

import (
	"context"
	"errors"
	"fmt"
	"io"
	"net"
	"os"
)

type FileLoadError struct {
	URL string
	Err error
}

func (p FileLoadError) Error() string {
	return fmt.Sprintf("%q: %v", p.URL, p.Err)
}

func (p FileLoadError) Unwrap() error {
	return p.Err
}

func (p FileLoadError) Timeout() bool {
	return true
}

type File struct {
}

func getFile(u string) (File, error) {
	return File{}, context.Canceled
}

func loadFiles(urls ...string) ([]File, error) {
	for _, url := range urls {
		_, err := getFile(url)
		if err != nil {
			return nil, FileLoadError{url, err}
		}
	}
	return nil, nil
}

func transfer() error {
	urls := []string{"www.yandex.ru"}

	files, err := loadFiles(urls...)
	if err != nil {
		if lfe, ok := err.(FileLoadError); ok { // Здесь мы ещё можем проверить тип ошибки.
			switch lfe.Err.(type) { // И даже поработать с нижележащей ошибкой.
			case *net.AddrError:
			case net.UnknownNetworkError:
			}
		}
		return fmt.Errorf("cannot load files: %v", err)
	}

	_ = files
	return nil
}

func handle() error {
	// ...
	if err := transfer(); err != nil {
		if _, ok := err.(FileLoadError); ok { // Здесь уже не можем написать подобное.
			// Никогда не случится, так как после fmt.Errorf в transfer
			// мы получим *errors.errorString вместо FileLoadError.
		}
		fmt.Printf("%T\n", err)
		return fmt.Errorf("cannot tranfer files: %v", err)
	}
	// ...
	return nil
}

func check() {
	var err error = FileLoadError{Err: context.Canceled}
	var i interface {
		Timeout() string
	}
	if errors.As(err, i) {
		fmt.Println(i.Timeout())
	}

	if lfe, ok := err.(FileLoadError); ok && lfe.Err == context.Canceled {
		fmt.Println("ok")
	}
	if errors.Is(err, context.Canceled) {
		fmt.Println("ok")
	}
	switch {
	case errors.Is(err, context.Canceled):
		// ...
	case errors.Is(err, io.EOF):
		// ...
	}
	if errors.Is(err, context.Canceled) {
		fmt.Println("ok")
	}

	err = FileLoadError{URL: "www.yandex.ru"}

	//var e FileLoadError
	//if errors.As(err, &e) {
	//	fmt.Println(e.URL)
	//}

	var e *os.PathError
	if errors.As(err, &e) {
		fmt.Println(e.Op, e.Path)
	}

	//var err2 interface{
	//	Timeout() bool
	//}
	//if errors.As(err, &err2) {
	//	fmt.Println(err2.Timeout())
	//}

	err = &net.AddrError{Addr: "0.0.0.0"}

	var n *net.AddrError
	switch {
	case errors.As(err, &n):
		// ...
	case errors.As(err, new(net.UnknownNetworkError)):
		// ..
	case errors.Is(err, context.Canceled):
		// ...
	case errors.Is(err, io.EOF):
		// ...
	}
	if lfe, ok := err.(FileLoadError); ok { // Здесь мы ещё можем проверить тип ошибки.
		switch lfe.Err.(type) { // И даже поработать с нижележащей ошибкой.
		case *net.AddrError:
		case net.UnknownNetworkError:
		}
	}
}

func main() {
	check()
}
