package evaluator

import (
	"github.com/NickDeChip/bottle-brush/pkg/ast"
	"github.com/NickDeChip/bottle-brush/pkg/object"
)

func evalVar(node *ast.VarStatement, env *object.Environment) object.Object {
	val := Eval(node.Value, env)
	if isError(val) {
		return val
	}
	if _, ok := env.GetVar(node.Name.Value); ok {
		return newError("the identifier %s is already declared; line=%d; col=%d", node.Name.Value, node.Name.Token.Line, node.Name.Token.Col)
	}
	env.SetVar(node.Name.Value, object.Var{
		Object: val,
		Mut:    node.Mut,
	})
	return nil
}
