package evaluator

import (
	"github.com/NickDeChip/bottle-brush/pkg/ast"
	"github.com/NickDeChip/bottle-brush/pkg/object"
)

func Eval(node ast.Node, env *object.Environment) object.Object {
	switch node := node.(type) {
	case *ast.Program:
		return evalProgram(node, env)
	case *ast.ExpressionStatement:
		return Eval(node.Expression, env)
	case *ast.VarStatement:
		val := Eval(node.Value, env)
		if isError(val) {
			return val
		}
		env.Set(node.Name.Value, val)
	case *ast.StringLiteral:
		return &object.String{Value: node.Value}
	}
	return nil
}

func evalProgram(prog *ast.Program, env *object.Environment) object.Object {
	var result object.Object

	for _, stmt := range prog.Statements {
		result = Eval(stmt, env)
	}

	return result
}

func isError(obj object.Object) bool {
	return false
}
