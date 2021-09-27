package main

import (
	"context"
	"fmt"
	"io"
	"net/http"
)

func networkRequest(ctx context.Context, url string) ([]byte, error) {
	req, err := http.NewRequestWithContext(ctx, http.MethodGet, url, nil)
	if err != nil {
		return nil, fmt.Errorf("build request: %v", err)
	}

	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		return nil, fmt.Errorf("do request: %v", err)
	}
	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, fmt.Errorf("read body: %v", err)
	}

	if err = resp.Body.Close(); err != nil {
		return nil, fmt.Errorf("close body: %v", err)
	}

	fmt.Printf("requesting %q - OK\n", url)
	return body, nil
}
