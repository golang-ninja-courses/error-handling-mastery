package main

import (
	"fmt"
	"io"
	"log"
	"net/http"
)

func main() {
	index, err := httpGet("http://www.golang-ninja.ru/")
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println(string(index)[:300])
}

func httpGet(url string) ([]byte, error) {
	res, err := http.Get(url) //nolint:noctx
	if err != nil {
		return nil, fmt.Errorf("cannot do GET: %v", err)
	}
	defer res.Body.Close()

	body, err := io.ReadAll(res.Body)
	if err != nil {
		return nil, fmt.Errorf("cannot read body: %v", err)
	}

	if err := res.Body.Close(); err != nil {
		return nil, fmt.Errorf("cannot close body: %v", err)
	}
	return body, nil
}
