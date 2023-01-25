package evaluator

import (
	"github.com/NickDeChip/bottlebrush/pkg/ast"
	"github.com/NickDeChip/bottlebrush/pkg/object"
)

func evalContinue(node *ast.ContinueStatement, env *object.Environment) object.Object {
	return CONTINUE
}
