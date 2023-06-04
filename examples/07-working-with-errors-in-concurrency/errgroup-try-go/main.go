package main

import (
	"log"
	"time"

	"golang.org/x/sync/errgroup"
)

func main() {
	var eg errgroup.Group
	eg.SetLimit(1)

	ready := make(chan struct{})
	eg.Go(func() error {
		defer close(ready)
		time.Sleep(3 * time.Second)
		return nil
	})

	f := func() error {
		log.Println("SUCCESS")
		return nil
	}

	log.Println(eg.TryGo(f))
	log.Println(eg.TryGo(f))
	<-ready
	log.Println(eg.TryGo(f))

	_ = eg.Wait()
}

/*
2023/06/04 12:19:57 false
2023/06/04 12:19:57 false
2023/06/04 12:20:00 true
2023/06/04 12:20:00 SUCCESS
*/
