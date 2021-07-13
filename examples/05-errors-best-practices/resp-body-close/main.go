package main

import (
	"fmt"
	"io"
	"log"
	"net/http"
)

func main() {
	res, err := http.Get("http://www.golang-courses.ru/")
	if err != nil {
		log.Fatal(err)
	}
	defer res.Body.Close()
	/*
		defer func() {
			if err := res.Body.Close(); err != nil {
				log.Println("cannot close response body: " + err.Error())
			}
		}()
	*/

	index, err := io.ReadAll(res.Body)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println(string(index)[:300])
}
