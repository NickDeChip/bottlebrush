package evaluator

import (
	"github.com/NickDeChip/bottle-brush/pkg/ast"
	"github.com/NickDeChip/bottle-brush/pkg/object"
)

func evalAssignment(node *ast.AssignmentStatement, env *object.Environment) object.Object {
	val := Eval(node.Value, env)
	if isError(val) {
		return val
	}
	oldVal, ok := env.Get(node.Name.Value)
	if !ok {
		return newError("the identifier %s is not declared; line=%d; col=%d", node.Name.Value, node.Name.Token.Line, node.Name.Token.Col)
	}
	if !oldVal.Mut {
		return newError("The identifier %s is a constant value and can not be chaged; line=%d; col=%d", node.Name.Value, node.Name.Token.Line, node.Name.Token.Col)
	}
	env.Set(node.Name.Value, object.Var{
		Object: val,
		Mut:    true,
	})
	return nil
}
