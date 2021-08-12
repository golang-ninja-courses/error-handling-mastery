package miniword

import (
	"io"
)

const maxPages = 3

var (
	errInvalidPageIndex error
	errNoMorePages      error
	errEmptyText        error
)

type Document struct { // Реализуй меня.
}

func NewDocument() *Document {
	// Реализуй меня.
	return nil
}

func (d *Document) AddPage() {
	// Реализуй меня.
}

func (d *Document) SetActivePage(index int) {
	// Реализуй меня.
}

func (d *Document) WriteText(s string) {
	// Реализуй меня.
}

func (d *Document) WriteTo(w io.Writer) (n int64, err error) {
	// Реализуй меня.
	return 0, nil
}
