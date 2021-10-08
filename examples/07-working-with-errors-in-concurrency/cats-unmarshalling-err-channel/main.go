package main

import (
	"encoding/json"
	"fmt"
	"sync"
)

type Cat struct {
	Name string `json:"name"`
}

func main() {
	catsJSONs := []string{`{"name": "Bobby"}`, `"name": "Billy"`, `{"name": "Васёк"}`}

	done := make(chan struct{})
	catsCh := make(chan Cat)
	errsCh := make(chan error)

	for _, catData := range catsJSONs {
		go func(catData string) { // Разбираем котиков в несколько горутин.
			var cat Cat
			if err := json.Unmarshal([]byte(catData), &cat); err != nil {
				errsCh <- err // Случилась ошибка.
				return
			}
			catsCh <- cat // Всё прошло хорошо.
		}(catData)
	}

	var wg sync.WaitGroup
	wg.Add(len(catsJSONs))

	go func() {
		wg.Wait()
		close(done)
	}()

	for {
		select {
		case <-done:
			return

		case c := <-catsCh:
			wg.Done()
			fmt.Println(c)

		case err := <-errsCh:
			wg.Done()
			fmt.Println(err)
		}
	}
}
