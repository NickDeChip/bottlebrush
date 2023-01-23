package evaluator

import (
	"github.com/NickDeChip/bottle-brush/pkg/ast"
	"github.com/NickDeChip/bottle-brush/pkg/object"
)

func evalFn(node *ast.FunctionLiteral, env *object.Environment) object.Object {
	params := node.Parameters
	body := node.Body
	return &object.Fn{
		Parameters: params,
		Env:        env,
		Body:       body,
	}
}
