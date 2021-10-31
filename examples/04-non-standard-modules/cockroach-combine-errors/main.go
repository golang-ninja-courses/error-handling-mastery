package main

import (
	"context"
	"fmt"
	"io/ioutil"
	"net/http"

	"github.com/cockroachdb/errors"
)

const google = "https://www.google.com"

func GetPage(ctx context.Context, url string) (data []byte, err error) {
	req, err := http.NewRequestWithContext(ctx, http.MethodGet, url, nil)
	if err != nil {
		return nil, errors.WithMessage(err, "create req")
	}

	response, err := http.DefaultClient.Do(req)
	if err != nil {
		return nil, errors.WithMessage(err, "do request")
	}
	defer func() {
		// Дополняем основную ошибку второстепенной, если она возникнет.
		// При этом, если есть только второстепенная ошибка, то она и будет возвращена.
		//
		// ВАЖНО: использование именованного возвращаемого значения.
		err = errors.CombineErrors(err, response.Body.Close())
	}()

	data, err = ioutil.ReadAll(response.Body)
	if err != nil {
		return nil, errors.WithMessage(err, "read body")
	}
	return data, nil
}

func main() {
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	d, err := GetPage(ctx, google)
	if err != nil {
		fmt.Printf("%+v", err)
		return
	}
	fmt.Println(string(d))
}
