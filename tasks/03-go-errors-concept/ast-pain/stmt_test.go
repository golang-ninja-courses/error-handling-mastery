package astpain

import (
	"go/ast"
	"go/parser"
	"go/token"
	"testing"

	"github.com/stretchr/testify/require"
)

func TestGetDeferredFunctionName(t *testing.T) {
	src := `
package main

import (
	"fmt"
	"sync"
)

func main() {
	var wg sync.WaitGroup
	wg.Add(1)

	go func() {
		defer wg.Done()
		bar()
	}()

	wg.Wait()
}

func bar() {
	defer fmt.Println("hello")
	defer func() {
		fmt.Println("world")
	}()
	defer foo()
}

func foo() {
}
`
	f, err := parser.ParseFile(token.NewFileSet(), "", src, parser.AllErrors)
	require.NoError(t, err)

	var funcNames []string
	ast.Inspect(f, func(node ast.Node) bool {
		if n := GetDeferredFunctionName(node); n != "" {
			funcNames = append(funcNames, n)
		}
		return node != nil
	})
	require.ElementsMatch(t, funcNames, []string{
		"wg.Done",
		"fmt.Println",
		"anonymous func",
		"foo",
	})
}
