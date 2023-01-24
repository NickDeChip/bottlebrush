package evaluator

import (
	"fmt"

	"github.com/NickDeChip/bottle-brush/pkg/ast"
	"github.com/NickDeChip/bottle-brush/pkg/object"
)

var (
	NULL  = &object.Null{}
	TRUE  = &object.Bool{Value: true}
	FALSE = &object.Bool{Value: false}
)

func Eval(node ast.Node, env *object.Environment) object.Object {
	switch node := node.(type) {
	case *ast.Program:
		return evalProgram(node, env)
	case *ast.BlockStatement:
		return evalBlock(node, env)
	case *ast.ExpressionStatement:
		return Eval(node.Expression, env)
	case *ast.VarStatement:
		return evalVar(node, env)
	case *ast.AssignmentStatement:
		return evalAssignment(node, env)
	case *ast.StringLiteral:
		return &object.String{Value: node.Value}
	case *ast.FloatLiteral:
		return &object.Float{Value: node.Value}
	case *ast.IntegerLiteral:
		return &object.Integer{Value: node.Value}
	case *ast.IfExspression:
		return evalIf(node, env)
	case *ast.Identifier:
		return evalIdentifier(node, env)
	case *ast.FunctionLiteral:
		return evalFn(node, env)
	case *ast.CallExpression:
		return evalCallExpression(node, env)
	case *ast.ReturnStatement:
		return evalReturn(node, env)
	case *ast.InfixExpression:
		return evalInfix(node, env)
	case *ast.PrefixExpression:
		return evalPrefix(node, env)
	case *ast.Bool:
		return evalBool(node.Value)
	}
	return nil
}

func evalProgram(prog *ast.Program, env *object.Environment) object.Object {
	var result object.Object

	for _, stmt := range prog.Statements {
		result = Eval(stmt, env)
		switch result := result.(type) {
		case *object.ReturnValue:
			return result.Value
		case *object.Error:
			return result
		}
	}

	return result
}

func evalExpresions(exps []ast.Expression, env *object.Environment) []object.Object {
	var result []object.Object

	for _, e := range exps {
		evaluated := Eval(e, env)
		if isError(evaluated) {
			return []object.Object{evaluated}
		}
		result = append(result, evaluated)
	}

	return result
}

func isError(obj object.Object) bool {
	if obj != nil {
		return obj.Type() == object.ERROR
	}
	return false
}

func newError(format string, a ...interface{}) *object.Error {
	return &object.Error{
		Message: fmt.Sprintf(format, a...),
	}
}
