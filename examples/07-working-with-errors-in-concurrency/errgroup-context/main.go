package main

import (
	"bytes"
	"context"
	"fmt"
	"io"
	"net/http"

	"golang.org/x/sync/errgroup"
)

const (
	workersCount = 2

	urlsCount  = 100
	googleUrl  = "https://www.google.com"
	invalidUrl = "https://invalid_url"
)

func networkRequest(ctx context.Context, url string) ([]byte, error) {
	req, err := http.NewRequestWithContext(ctx, http.MethodGet, url, &bytes.Buffer{})
	if err != nil {
		return nil, fmt.Errorf("cannont build request: %w", err)
	}

	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		return nil, fmt.Errorf("cannont do request: %w", err)
	}
	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, fmt.Errorf("cannot read body: %v", err)
	}

	if err = resp.Body.Close(); err != nil {
		return nil, fmt.Errorf("cannot close body: %v", err)
	}

	fmt.Printf("requesting %q OK\n", url)

	return body, nil
}

// Тут решаем задачу – сходить по N ссылкам в workersCount горутинах, что-то по ним запросить в networkRequest().
// При этом, если случается хотя бы одна ошибка, то хотим прервать выполнение, а не ходить по оставшимся ссылкам.
func main() {
	urls := make(chan string)

	go func() {
		defer close(urls)

		for i := 0; i < urlsCount; i++ {
			url := googleUrl
			if i == 2 {
				url = invalidUrl
			}

			urls <- url
		}
	}()

	ctx := context.Background()
	eg, ctx := errgroup.WithContext(ctx)
	// eg := errgroup.Group{} // попробуйте раскомментировать эту строчку и закомментировать строчку выше и посмотреть на вывод в консоль

	for i := 0; i < workersCount; i++ {
		eg.Go(func() error {
			for url := range urls {
				select {
				case <-ctx.Done():
					return ctx.Err()
				default:
				}

				_, err := networkRequest(ctx, url)
				if err != nil {
					return fmt.Errorf("network request %s error: %w", url, err)
				}
			}

			return nil
		})
	}

	if err := eg.Wait(); err != nil {
		fmt.Print(err)
	}
}
