package main

import (
	"log"
	"net/http"
	"time"
)

func main() {
	http.HandleFunc("/", func(w http.ResponseWriter, req *http.Request) {
		time.Sleep(3 * time.Second)

		data := make([]byte, 1<<20) // 1Mb
		logWriteErr(w.Write(data))  // write tcp [::1]:8080->[::1]:60756: write: broken pipe
	})

	go func() {
		time.Sleep(3 * time.Second) // Wait for server start up.

		client := &http.Client{Timeout: time.Second}
		_, err := client.Get("http://localhost:8080")
		if err != nil {
			// context deadline exceeded (Client.Timeout exceeded while awaiting headers)
			log.Println("cannot do GET: " + err.Error())
		}
	}()

	if err := http.ListenAndServe(":8080", http.DefaultServeMux); err != nil {
		log.Fatal(err)
	}
}

func logWriteErr(_ int, err error) {
	if err != nil {
		log.Println("cannot write response: " + err.Error())
	}
}
