package astpain

import (
	"go/ast"
	"go/parser"
	"go/token"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestGetDeferredFunctionName(t *testing.T) {
	cases := []struct {
		name          string
		src           string
		expectedFuncs []string
	}{
		{
			name: "deferred method",
			src: `
package main

import "sync"

func main() {
	var wg sync.WaitGroup
	wg.Add(1)

	go func() {
		defer wg.Done()
		bar()
	}()

	wg.Wait()
}`,
			expectedFuncs: []string{"wg.Done"},
		},
		{
			name: "deferred pkg func and anonymous func",
			src: `
package main

import (
	"fmt"
)

func bar() {
	defer fmt.Println("hello")
	defer func() {
		fmt.Println("world")
	}()
}
`,
			expectedFuncs: []string{"fmt.Println", "anonymous func"},
		},
		{
			name: "deferred func",
			src: `
package main

func bar() {
    defer foo()
}
`,
			expectedFuncs: []string{"foo"},
		},
	}

	for _, tt := range cases {
		t.Run(tt.name, func(t *testing.T) {
			assertDeferredFuncs(t, tt.src, tt.expectedFuncs)
		})
	}
}

func assertDeferredFuncs(t *testing.T, src string, expected []string) {
	t.Helper()

	f, err := parser.ParseFile(token.NewFileSet(), "", src, parser.AllErrors)
	require.NoError(t, err)

	var funcNames []string
	ast.Inspect(f, func(node ast.Node) bool {
		if n := GetDeferredFunctionName(node); n != "" {
			funcNames = append(funcNames, n)
		}
		return node != nil
	})
	assert.Equal(t, funcNames, expected)
}
