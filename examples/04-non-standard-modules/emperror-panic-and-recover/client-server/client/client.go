package main

import (
	"bytes"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
)

const url = "http://127.0.0.1:8888/"

var (
	regularRequest = []byte("regular")
	errorRequest   = []byte("error")
	panicRequest   = []byte("panic")
)

func main() {
	for _, request := range [][]byte{regularRequest, errorRequest, panicRequest} {
		if err := makeRequest(request); err != nil {
			log.Println(err)
		}
	}
}

func makeRequest(request []byte) error {
	resp, err := http.Post(url, "application/json", bytes.NewBuffer(request)) //nolint:noctx
	if err != nil {
		return fmt.Errorf("do post: %v", err)
	}
	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return fmt.Errorf("read body: %v", err)
	}

	log.Printf("HTTP code: %d, Body: %s", resp.StatusCode, string(body))
	return resp.Body.Close()
}
