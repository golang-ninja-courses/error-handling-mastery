package main

import (
	"errors"
	"fmt"

	"golang.org/x/sync/errgroup"
)

func main() {
	if err := work(); err != nil {
		fmt.Println(err) // something bad has happened
	}
}

func work() error {
	var eg errgroup.Group

	eg.Go(func() error {
		// Выполняем какую-то операцию, завершившуюся с ошибкой.
		// ...
		return errors.New("something bad has happened")
	})

	eg.Go(func() error {
		// Выполняем какую-то операцию, завершившуюся без ошибки.
		// ...
		return nil
	})

	// Дожидаемся окончания работ.
	// Возвращаем ошибку от любой из операций (если ошибка произошла).
	return eg.Wait()
}
