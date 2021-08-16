package main

import (
	"bytes"
	"fmt"
	"io/ioutil"
	"net/http"
)

const (
	url = "http://127.0.0.1:8888/"
)

var (
	regularRequest = []byte("regular")
	errorRequest   = []byte("error")
	panicRequest   = []byte("panic")
)

func makeRequest(request []byte) {
	resp, err := http.Post(url, "application/json", bytes.NewBuffer(request))
	if err != nil {
		fmt.Printf("request error: %v", err)
		return
	}

	defer func() {
		_ = resp.Body.Close()
	}()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		fmt.Printf("read body error: %v", err)
		return
	}

	fmt.Printf("HTTP code: %d, Body: %s\n", resp.StatusCode, string(body))
}

func main() {
	for _, request := range [][]byte{regularRequest, errorRequest, panicRequest} {
		makeRequest(request)
	}
}
