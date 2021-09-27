package main

import (
	"context"
	"fmt"

	"golang.org/x/sync/errgroup"
)

const (
	workersCount = 2

	urlsCount  = 100
	googleURL  = "https://www.google.com"
	invalidURL = "https://invalid_url"
)

// Решается следующая задача:
//  Необходимо сходить по urlsCount ссылкам в workersCount горутинах и получить данные по ним, используя networkRequest.
//  При этом, если случается хотя бы одна ошибка, то необходимо прервать выполнение программы, а не продолжать
//  обрабатывать оставшиеся ссылки.
func main() {
	ctx := context.Background()

	eg, ctx := errgroup.WithContext(ctx)
	// Попробуйте раскомментировать строчку ниже и закомментировать строчку выше,
	// и посмотреть на вывод в консоли.
	// var eg errgroup.Group

	urls := make(chan string)

	eg.Go(func() error {
		defer close(urls)

		for i := 0; i < urlsCount; i++ {
			url := googleURL
			if i == 2 {
				url = invalidURL
			}

			select {
			case <-ctx.Done():
				return ctx.Err()
			case urls <- url:
			}
		}

		return nil
	})

	for i := 0; i < workersCount; i++ {
		eg.Go(func() error {
			for {
				select {
				case <-ctx.Done():
					return ctx.Err()

				case url := <-urls:
					if _, err := networkRequest(ctx, url); err != nil {
						return fmt.Errorf("network request %s error: %v", url, err)
					}
				}
			}
		})
	}

	if err := eg.Wait(); err != nil {
		fmt.Print(err)
	}
}
