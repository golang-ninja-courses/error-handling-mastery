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
	catsJSON := []string{`{"name": "Bobby"}`, `"name": "Billy"`}
	catsChan := make(chan Cat, len(catsJSON))
	wg := &sync.WaitGroup{}
	wg.Add(len(catsJSON))

	for _, catData := range catsJSON {
		go func(catData string) { // разбираем котиков в несколько горутин
			defer wg.Done()
			cat := Cat{}
			if err := json.Unmarshal([]byte(catData), &cat); err != nil {
				// Случилась ошибка, что делать?
			}
			catsChan <- cat
		}(catData)
	}

	wg.Wait()
	close(catsChan)

	for cat := range catsChan {
		fmt.Println(cat)
	}
}
