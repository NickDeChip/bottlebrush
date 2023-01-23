package evaluator

import (
	"github.com/NickDeChip/bottle-brush/pkg/ast"
	"github.com/NickDeChip/bottle-brush/pkg/object"
)

func evalReturn(node *ast.ReturnStatement, env *object.Environment) object.Object {
	val := Eval(node.ReturnValue, env)
	if isError(val) {
		return val
	}
	return &object.ReturnValue{
		Value: val,
	}
}
