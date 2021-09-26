package main

import (
	"errors"
	"fmt"
	"sync"
)

func main() {
	if err := work(); err != nil {
		fmt.Println(err) // something bad has happened
	}
}

func work() error {
	// Будем выполнять два параллельных действия.
	var wg sync.WaitGroup
	errsCh := make(chan error, 2)

	wg.Add(1)
	go func() {
		defer wg.Done()
		// Выполняем какую-то операцию, завершившуюся с ошибкой.
		// ...
		errsCh <- errors.New("something bad has happened")
	}()

	wg.Add(1)
	go func() {
		defer wg.Done()
		// Выполняем какую-то операцию, завершившуюся без ошибки.
		// ...
	}()

	wg.Wait() // Дожидаемся окончания работ.

	// Возвращаем ошибку от любой из операций (если ошибка произошла).
	select {
	case err := <-errsCh:
		return err
	default:
		return nil
	}
}
