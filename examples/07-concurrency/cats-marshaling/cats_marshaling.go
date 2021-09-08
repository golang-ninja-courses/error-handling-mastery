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
	catsJson := []string{`{"name": "Bobby"}`, `"name": "Billy"`}
	catsChan := make(chan Cat, len(catsJson))
	wg := &sync.WaitGroup{}
	wg.Add(len(catsJson))

	for _, catJson := range catsJson {
		go func(catJson string) { // разбираем котиков в несколько горутин
			defer wg.Done()
			cat := Cat{}
			if err := json.Unmarshal([]byte(catJson), &cat); err != nil {
				// Случилась ошибка, что делать?
			}
			catsChan <- cat
		}(catJson)
	}

	wg.Wait()
	close(catsChan)

	for cat := range catsChan {
		fmt.Println(cat)
	}
}
