package main

import (
	"errors"
	"fmt"
	"io"
	"os"

	"github.com/jackc/pgx"
)

func main() {
	exist := errors.New("file already exists")
	fmt.Println(errors.Is(exist, os.ErrExist)) // false

	eof := errors.New("EOF")
	fmt.Println(errors.Is(eof, io.EOF)) // false

	noRows := errors.New("no rows in result set")
	fmt.Println(errors.Is(noRows, pgx.ErrNoRows)) // false
}
