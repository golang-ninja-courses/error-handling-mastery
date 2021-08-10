package main

import (
	"bufio"
	"errors"
	"fmt"
	"os"

	"go.uber.org/multierr"
)

var (
	errCloserMock = errors.New("close error")
)

type CloserMock struct{}

func (c *CloserMock) Close() error {
	return errCloserMock
}

// processFile – пример штатного использования multierr.AppendInvoke().
// Обратите внимание, что multierr.AppendInvoke() используется вкупе с именованным возвращаемым аргументом err.
func processFile(path string) (err error) {
	f, err := os.Open(path)
	if err != nil {
		return err
	}
	defer multierr.AppendInvoke(&err, multierr.Close(f))

	scanner := bufio.NewScanner(f)
	defer multierr.AppendInvoke(&err, multierr.Invoke(scanner.Err))

	for scanner.Scan() {
		fmt.Println(string(scanner.Bytes()))
	}

	return err
}

func processCloserMock() (err error) {
	mock := CloserMock{}
	defer multierr.AppendInvoke(&err, multierr.Close(&mock))

	return err
}

func processCloserMockWithError() (err error) {
	mock := CloserMock{}
	defer multierr.AppendInvoke(&err, multierr.Close(&mock))

	err = errors.New("an error")

	return err
}

func main() {
	err := processFile("./examples/04-non-standard-modules/multierr-append-invoke/test_file.txt")
	if err != nil {
		fmt.Printf("%v\n", err) // ничего не выведет
	}

	err = processCloserMock()
	if err != nil {
		fmt.Printf("%v\n", err) // выведет "close error"
	}

	err = processCloserMockWithError()
	if err != nil {
		fmt.Printf("%v\n", err) // выведет an error; close error
	}
}
