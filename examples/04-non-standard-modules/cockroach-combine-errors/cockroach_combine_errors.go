package main

import (
	"fmt"
	"io/ioutil"
	"net/http"

	"github.com/cockroachdb/errors"
)

const (
	google  = "https://www.google.com"
	invalid = "https://invalid"
)

func makeRequest(url string) error {
	var err error

	response, err := http.Get(url)
	if err != nil {
		return err
	}
	defer func() {
		// Дополняем основную ошибку второстепенной, если она возникнет.
		err = errors.WithSecondaryError(err, response.Body.Close())
	}()

	_, err = ioutil.ReadAll(response.Body)
	if err != nil {
		return err
	}

	return nil
}

func main() {
	err := errors.CombineErrors(makeRequest(google), makeRequest(invalid))
	if err != nil {
		fmt.Printf("%+v", err)
		return
	}

	fmt.Println("OK")
}
