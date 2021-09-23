package main

import (
	"encoding/json"
	"fmt"
	"sync"
)

type Cat struct {
	Name string `json:"name"`
}

type CatContainer struct {
	Cat Cat
	Err error // Сюда складываем ошибку, если что-то пошло не так.
}

func main() {
	catsJSONs := []string{`{"name": "Bobby"}`, `"name": "Billy"`, `{"name": "Васёк"}`}
	catsCh := make(chan CatContainer, len(catsJSONs))

	var wg sync.WaitGroup
	wg.Add(len(catsJSONs))

	for _, catData := range catsJSONs {
		go func(catData string) {
			defer wg.Done()

			var cat Cat
			if err := json.Unmarshal([]byte(catData), &cat); err != nil {
				catsCh <- CatContainer{Err: err} // Случилась ошибка.
				return
			}
			catsCh <- CatContainer{Cat: cat} // Всё прошло хорошо.
		}(catData)
	}

	go func() {
		wg.Wait()
		close(catsCh)
	}()

	for catContainer := range catsCh {
		if catContainer.Err != nil {
			fmt.Println("ERROR:", catContainer.Err)
			continue
		}
		fmt.Println(catContainer.Cat)
	}
}

/*
ERROR: invalid character ':' after top-level value
{Васёк}
{Bobby}
*/
