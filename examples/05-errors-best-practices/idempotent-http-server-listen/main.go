package main

import (
	"fmt"
	"net/http"
	"time"
)

func main() {
	var s http.Server

	go func() {
		time.Sleep(time.Second * 3)

		if err := s.Close(); err != nil {
			panic(err)
		}

		for i := 0; i < 5; i++ {
			if err := s.ListenAndServe(); err != nil {
				fmt.Println(err)
			}
		}
	}()

	if err := s.ListenAndServe(); err != nil {
		fmt.Println(err)
	}

	/*
		http: Server closed
		http: Server closed
		http: Server closed
		http: Server closed
		http: Server closed
		http: Server closed
	*/
}
