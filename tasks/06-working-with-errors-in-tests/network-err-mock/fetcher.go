package fetcher

import (
	"context"
	"errors"
	"fmt"
	"io/ioutil"
	"net/http"
)

var (
	ErrFetchTimeout = errors.New("fetch url timeout")
	ErrFetchTmp     = errors.New("fetch url temporary error")
)

type Client interface {
	Do(req *http.Request) (*http.Response, error)
}

func FetchURL(ctx context.Context, client Client, url string) ([]byte, error) {
	req, err := http.NewRequestWithContext(ctx, http.MethodGet, url, nil)
	if err != nil {
		return nil, fmt.Errorf("create request: %v", err)
	}

	resp, err := client.Do(req)
	if err != nil {
		if isTimeout(err) {
			return nil, fmt.Errorf("%w: %v", ErrFetchTimeout, err)
		}
		if isTemporary(err) {
			return nil, fmt.Errorf("%w: %v", ErrFetchTmp, err)
		}
		return nil, fmt.Errorf("do request: %w", err)
	}
	defer resp.Body.Close()

	data, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, fmt.Errorf("read all body: %v", err)
	}

	if err := resp.Body.Close(); err != nil {
		return nil, fmt.Errorf("close body: %v", err)
	}
	return data, nil
}

func isTimeout(err error) bool {
	var i interface{ Timeout() bool }
	return errors.As(err, &i) && i.Timeout()
}

func isTemporary(err error) bool {
	var i interface{ Temporary() bool }
	return errors.As(err, &i) && i.Temporary()
}
