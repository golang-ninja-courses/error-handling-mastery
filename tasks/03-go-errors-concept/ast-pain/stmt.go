package astpain

import (
	"go/ast"
)

// GetDeferredFunctionName возвращает имя функции, чей вызов был отложен через defer,
// если входящий node является *ast.DeferStmt.
func GetDeferredFunctionName(node ast.Node) string {
	// Реализуй меня.
	return ""
}
