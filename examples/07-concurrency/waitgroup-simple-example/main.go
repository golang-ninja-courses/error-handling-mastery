package main

import (
	"errors"
	"fmt"
	"sync"
)

func main() {
	// Будем выполнять два параллельных действия
	wg := &sync.WaitGroup{}
	wg.Add(2)
	errChan := make(chan error, 2)

	go func() {
		defer wg.Done()
		// Что-то делаем
		errChan <- errors.New("something bad has happened")
	}()

	go func() {
		defer wg.Done()
		// Что-то делаем
		errChan <- nil
	}()

	wg.Wait() // Дожидаемся окончания работ

	close(errChan)

	for err := range errChan { // Проверяем, всё ли выполнилось без ошибок
		if err != nil {
			fmt.Println(err)
		}
	}
}
