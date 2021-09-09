//nolint:govet
package main

import (
	"fmt"

	"github.com/labstack/echo/v4"
)

func Handle() error {
	if err := usefulWork; err != nil {
		return fmt.Errorf("%w: %v", echo.ErrInternalServerError, err)
	}
	return nil
}

func usefulWork() error {
	return nil
}

func main() {
	if err := Handle(); err != nil {
		fmt.Println("handle err:", err)
	} else {
		fmt.Println("no handle err")
	}
}
