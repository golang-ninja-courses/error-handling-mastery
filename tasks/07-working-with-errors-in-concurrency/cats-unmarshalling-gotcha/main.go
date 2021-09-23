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
	catsCh := make(chan Cat, len(catsJSONs))

	var wg sync.WaitGroup
	wg.Add(len(catsJSONs))

	var err error // Заводим специальную переменную для хранения ошибки.

	for _, catData := range catsJSONs {
		go func(catData string) {
			defer wg.Done()

			var cat Cat
			err = json.Unmarshal([]byte(catData), &cat)
			catsCh <- cat
		}(catData)
	}

	wg.Wait()
	fmt.Println(err) // Выводим ошибку на экран.
}
