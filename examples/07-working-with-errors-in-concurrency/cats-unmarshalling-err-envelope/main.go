package main

import (
	"encoding/json"
	"errors"
	"fmt"
	"sync"
)

type Cat struct {
	Name string `json:"name"`
}

type CatEnvelope struct {
	cat Cat
	err error // Сюда складываем ошибку, если что-то пошло не так.
}

func NewCatEnvelope(c Cat) CatEnvelope {
	return CatEnvelope{cat: c}
}

func NewCatEnvelopeWithErr(err error) CatEnvelope {
	return CatEnvelope{err: err}
}

func (c CatEnvelope) Unpack() (Cat, error) {
	return c.cat, c.err
}

func main() {
	catsJSONs := []string{`{"name": "Bobby"}`, `"name": "Billy"`, `{"name": "Васёк"}`}
	catsCh := make(chan CatEnvelope, len(catsJSONs))

	var wg sync.WaitGroup
	wg.Add(len(catsJSONs))

	for _, catData := range catsJSONs {
		go func(catData string) {
			defer wg.Done()

			var cat Cat
			if err := json.Unmarshal([]byte(catData), &cat); err != nil {
				catsCh <- NewCatEnvelopeWithErr(err) // Случилась ошибка.
				return
			}
			catsCh <- NewCatEnvelope(cat) // Всё прошло хорошо.
		}(catData)
	}

	go func() {
		wg.Wait()
		close(catsCh)
	}()

	for catEnvelope := range catsCh {
		// Мы можем как угодно менять копии, полученные после распаковки,
		// не влияя на оригинальный контейнер.
		c, err := catEnvelope.Unpack()
		c = Cat{Name: "Hacked"}
		err = errors.New("hacked")

		c, err = catEnvelope.Unpack()
		if err != nil {
			fmt.Println("ERROR:", err)
			continue
		}
		fmt.Println(c)
	}
}

/*
ERROR: invalid character ':' after top-level value
{Васёк}
{Bobby}
*/
