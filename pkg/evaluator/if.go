package evaluator

import (
	"github.com/NickDeChip/bottle-brush/pkg/ast"
	"github.com/NickDeChip/bottle-brush/pkg/object"
)

func evalIf(node *ast.IfExspression, env *object.Environment) object.Object {
	condition := Eval(node.Condition, env)
	if isError(condition) {
		return condition
	}

	if condition.Type() != object.BOOL {
		return newError("expected condition type %s, got: %s; line=%d; col=%d ", object.BOOL, condition.Type(), node.Line(), node.Col())
	}

	if condition == TRUE {
		ifEnv := object.NewEncolsedEnvironment(env)
		return Eval(node.Consequence, ifEnv)
	} else {
		return NULL
	}
}
