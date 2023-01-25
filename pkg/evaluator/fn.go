package evaluator

import (
	"github.com/NickDeChip/bottlebrush/pkg/ast"
	"github.com/NickDeChip/bottlebrush/pkg/object"
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
