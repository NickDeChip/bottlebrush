package evaluator

import (
	"github.com/NickDeChip/bottlebrush/pkg/ast"
	"github.com/NickDeChip/bottlebrush/pkg/object"
)

func evalFor(node *ast.ForStatement, env *object.Environment) object.Object {
	for {
		condition := Eval(node.Condition, env)
		if isError(condition) {
			return condition
		}

		if condition.Type() != object.BOOL {
			return newError("expected condition type %s, got: %s; line=%d; col=%d ", object.BOOL, condition.Type(), node.Line(), node.Col())
		}

		if condition == TRUE {
			forEnv := object.NewEncolsedEnvironment(env)
			res := Eval(node.Consequence, forEnv)
			if res == nil {
				continue
			}
			if res.Type() == object.ERROR || res.Type() == object.RETURN {
				return res
			}
		} else {
			return NULL
		}
	}
}
