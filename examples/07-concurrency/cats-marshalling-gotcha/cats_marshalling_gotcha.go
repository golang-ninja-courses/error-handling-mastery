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

	var err error // Заводим специальную переменную err для хранения ошибки

	for _, catJson := range catsJson {
		go func(catJson string) {
			defer wg.Done()
			cat := Cat{}
			err = json.Unmarshal([]byte(catJson), &cat)
			catsChan <- cat
		}(catJson)
	}

	wg.Wait()

	fmt.Println(err)
}
