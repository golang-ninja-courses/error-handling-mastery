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

	var err error // Заводим специальную переменную err для хранения ошибки

	for _, catData := range catsJSON {
		go func(catData string) {
			defer wg.Done()
			cat := Cat{}
			err = json.Unmarshal([]byte(catData), &cat)
			catsChan <- cat
		}(catData)
	}

	wg.Wait()

	fmt.Println(err)
}
