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
	catsCh := make(chan Cat)

	var wg sync.WaitGroup
	wg.Add(len(catsJSONs))

	for _, catData := range catsJSONs {
		go func(catData string) { // Разбираем котиков в несколько горутин.
			defer wg.Done()

			var cat Cat
			if err := json.Unmarshal([]byte(catData), &cat); err != nil {
				// Случилась ошибка, что делать?
			}
			catsCh <- cat
		}(catData)
	}

	go func() {
		wg.Wait()
		close(catsCh)
	}()

	for cat := range catsCh {
		fmt.Println(cat)
	}
}
