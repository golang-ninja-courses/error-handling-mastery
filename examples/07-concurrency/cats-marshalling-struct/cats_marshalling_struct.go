package main

import (
	"encoding/json"
	"fmt"
	"sync"
)

type Cat struct {
	Name string `json:"name"`
}

type CatResult struct {
	Cat Cat
	Err error // сюда складываем ошибку, если что-то пошло не так
}

func main() {
	catsJson := []string{`{"name": "Bobby"}`, `"name": "Billy"`}
	catsChan := make(chan CatResult, len(catsJson))
	wg := &sync.WaitGroup{}
	wg.Add(len(catsJson))

	for _, catJson := range catsJson {
		go func(catJson string) {
			defer wg.Done()
			cat := Cat{}
			if err := json.Unmarshal([]byte(catJson), &cat); err != nil {
				catsChan <- CatResult{Err: err} // случилась ошибка
				return
			}
			catsChan <- CatResult{Cat: cat} // всё прошло хорошо
		}(catJson)
	}

	wg.Wait()
	close(catsChan)

	for catResult := range catsChan {
		if catResult.Err != nil {
			fmt.Println(catResult.Err)
			continue
		}
		fmt.Println(catResult.Cat)
	}
}
