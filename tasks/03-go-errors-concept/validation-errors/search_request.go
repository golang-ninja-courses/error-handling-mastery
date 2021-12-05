package requests

import (
	// Доступные пакеты, _ для сохранения импортов.
	_ "errors"
	_ "fmt"
	_ "regexp"
	_ "strings"
)

const maxPageSize = 100

// Реализуй нас.
var (
	errIsNotRegexp     error
	errInvalidPage     error
	errInvalidPageSize error
)

// Реализуй мои методы.
type ValidationErrors []error

type SearchRequest struct {
	Exp      string
	Page     int
	PageSize int
}

func (r SearchRequest) Validate() error {
	// Реализуй меня.
	return nil
}
