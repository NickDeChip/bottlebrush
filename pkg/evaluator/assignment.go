package evaluator

import (
	"github.com/NickDeChip/bottlebrush/pkg/ast"
	"github.com/NickDeChip/bottlebrush/pkg/object"
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

func evalIndexAssignment(node *ast.IndexAssignmentStatement, env *object.Environment) object.Object {
	val := Eval(node.Value, env)
	if isError(val) {
		return val
	}

	idx := Eval(node.Index, env)
	if isError(idx) {
		return idx
	}

	if idx.Type() != object.INTEGER {
		return newError("index operator not supported: %s; line=%d; col=%d", idx.Type(), node.Index.Line(), node.Index.Col())
	}

	oldVal, ok := env.Get(node.Name.Value)
	if !ok {
		return newError("the identifier %s is not declared; line=%d; col=%d", node.Name.Value, node.Name.Token.Line, node.Name.Token.Col)
	}

	if oldVal.Object.Type() != object.LIST {
		return newError("unexpected index assignment type: %s; line=%d; col=%d", val.Type(), node.Line(), node.Col())
	}

	oldVal.Object.(*object.List).Elements[idx.(*object.Integer).Value] = val

	env.Set(node.Name.Value, oldVal)
	return nil
}
