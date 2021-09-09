//nolint:staticcheck
package main

import (
	"fmt"

	"github.com/labstack/echo/v4"
)

func Handle() error {
	var err *echo.HTTPError

	if err2 := usefulWork(); err2 != nil {
		err = echo.ErrInternalServerError
	}
	return err
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
