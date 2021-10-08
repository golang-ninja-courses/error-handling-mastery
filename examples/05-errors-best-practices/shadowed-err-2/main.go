package main

import (
	"fmt"
	"log"
)

func Handle() (err error) {
	if err = handleConn(); err != nil {
		// Новая ошибка не перетирает возвращаемую.
		// Попробуйте заменить `:=` на `=`.
		if err := closeConn(); err != nil {
			log.Println(err)
		}
	}
	return
}

func handleConn() error {
	return fmt.Errorf("handle conn error")
}

func closeConn() error {
	return fmt.Errorf("close error")
}

func main() {
	fmt.Println(Handle())
}

/*
2021/07/10 11:58:53 close error
handle conn error
*/
