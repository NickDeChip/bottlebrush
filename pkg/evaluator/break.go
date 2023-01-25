package evaluator

import (
	"github.com/NickDeChip/bottlebrush/pkg/ast"
	"github.com/NickDeChip/bottlebrush/pkg/object"
)

func evalBreak(node *ast.BreakStatement, env *object.Environment) object.Object {
	return BREAK
}
