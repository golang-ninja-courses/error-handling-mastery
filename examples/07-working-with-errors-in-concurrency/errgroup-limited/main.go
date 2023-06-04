package main

import (
	"log"
	"time"

	"golang.org/x/sync/errgroup"
)

func main() {
	var eg errgroup.Group

	eg.SetLimit(3)

	for i := 0; i < 11; i++ {
		i := i
		eg.Go(func() error {
			work(i)
			return nil
		})
	}

	_ = eg.Wait()
}

func work(i int) {
	log.Printf("worker %d: do hard work...\n", i)
	time.Sleep(3 * time.Second)
}

/*
2023/06/04 12:01:41 worker 1: do hard work...
2023/06/04 12:01:41 worker 2: do hard work...
2023/06/04 12:01:41 worker 0: do hard work...
2023/06/04 12:01:44 worker 3: do hard work...
2023/06/04 12:01:44 worker 4: do hard work...
2023/06/04 12:01:44 worker 5: do hard work...
2023/06/04 12:01:47 worker 8: do hard work...
2023/06/04 12:01:47 worker 7: do hard work...
2023/06/04 12:01:47 worker 6: do hard work...
2023/06/04 12:01:50 worker 10: do hard work...
2023/06/04 12:01:50 worker 9: do hard work...
*/
