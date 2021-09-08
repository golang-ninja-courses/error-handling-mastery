package main

import (
	"errors"
	"fmt"

	"golang.org/x/sync/errgroup"
)

func main() {
	eg := errgroup.Group{}

	eg.Go(func() error {
		// Что-то делаем
		return errors.New("something bad has happened")
	})

	eg.Go(func() error {
		// Что-то делаем
		return nil
	})

	if err := eg.Wait(); err != nil {  // Дожидаемся окончания работ и проверяем, всё ли выполнилось без ошибок
		fmt.Println(err)
	}
}
