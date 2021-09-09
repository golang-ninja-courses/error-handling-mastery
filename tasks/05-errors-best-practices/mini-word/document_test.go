package miniword

import (
	"bytes"
	"fmt"
	"os"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func ExampleDocument() {
	d := NewDocument()
	d.WriteText("Hello")

	if _, err := d.WriteTo(os.Stdout); err != nil {
		panic(err)
	}

	// Output:
	// --- Page 1 ---
	// Hello
	//
}

func TestDoc(t *testing.T) {
	d := NewDocument()
	d.WriteText("Text for first page!")

	d.AddPage()
	d.AddPage()

	d.SetActivePage(3)
	d.WriteText("Text for third page!")

	d.SetActivePage(2)
	d.WriteText("\t\tText for second page!\n")

	b := bytes.NewBuffer(nil)
	n, err := d.WriteTo(b)
	require.NoError(t, err)

	expectedText := `--- Page 1 ---
Text for first page!
--- Page 2 ---
		Text for second page!

--- Page 3 ---
Text for third page!
`
	assert.Equal(t, int64(len(expectedText)), n)
	assert.Equal(t, expectedText, b.String())
}

func TestDoc_AddPage(t *testing.T) {
	d := NewDocument()
	d.WriteText("Hello!")
	for i := 0; i < 100; i++ {
		d.AddPage()
	}
	d.WriteText("Hello!")

	b := bytes.NewBuffer(nil)
	n, err := d.WriteTo(b)
	require.ErrorIs(t, err, errNoMorePages)
	assert.Equal(t, int64(0), n)
	assert.Empty(t, b.String())
}

func TestDoc_SetActivePage(t *testing.T) {
	for _, i := range []int{-1, 100} {
		t.Run(fmt.Sprintf("page_num_%d", i), func(t *testing.T) {
			d := NewDocument()
			d.WriteText("Hello!")
			d.SetActivePage(i)
			d.AddPage()
			d.WriteText("Hello!")

			b := bytes.NewBuffer(nil)
			n, err := d.WriteTo(b)
			require.ErrorIs(t, err, errInvalidPageNumber)
			assert.Equal(t, int64(0), n)
			assert.Empty(t, b.String())
		})
	}
}

func TestDoc_WriteText(t *testing.T) {
	d := NewDocument()
	d.WriteText("Hello!")
	d.WriteText("")
	d.AddPage()
	d.SetActivePage(1)
	d.WriteText("Hello!")

	b := bytes.NewBuffer(nil)
	n, err := d.WriteTo(b)
	require.ErrorIs(t, err, errEmptyText)
	assert.Equal(t, int64(0), n)
	assert.Empty(t, b.String())
}

func TestDoc_MultipleErrors(t *testing.T) {
	d := NewDocument()
	d.WriteText("Hello!")
	d.WriteText("")
	d.SetActivePage(-1)
	for i := 0; i < 100; i++ {
		d.AddPage()
	}

	b := bytes.NewBuffer(nil)
	n, err := d.WriteTo(b)
	require.ErrorIs(t, err, errEmptyText)
	assert.Equal(t, int64(0), n)
	assert.Empty(t, b.String())
}

func TestErrorsAreFilled(t *testing.T) {
	assert.NotNil(t, errInvalidPageNumber)
	assert.NotNil(t, errNoMorePages)
	assert.NotNil(t, errEmptyText)
}
