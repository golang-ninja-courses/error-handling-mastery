package main

import (
	"fmt"
	"io"
	"net/http"
)

func main() {
	res, err := http.Get("http://www.golang-courses.ru/") //nolint:noctx
	if err != nil {
		panic(err)
	}
	defer res.Body.Close()
	/*
		defer func() {
			if err := res.Body.Close(); err != nil {
				log.Println("cannot close response body: " + err.Err())
			}
		}()
	*/

	index, err := io.ReadAll(res.Body)
	if err != nil {
		panic(err)
	}

	fmt.Println(string(index)[:300])
}
