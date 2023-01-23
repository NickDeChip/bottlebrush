package evaluator

import (
	"github.com/NickDeChip/bottle-brush/pkg/ast"
	"github.com/NickDeChip/bottle-brush/pkg/object"
)

func evalInfix(node *ast.InfixExpression, env *object.Environment) object.Object {
	left := Eval(node.Left, env)
	if isError(left) {
		return left
	}
	right := Eval(node.Right, env)
	if isError(right) {
		return right
	}
	return evalInfixExpression(node.Operator, left, right, node.Token.Line, node.Token.Col)
}

func evalInfixExpression(operator string, left, right object.Object, line, col int) object.Object {
	switch {
	case left.Type() == object.INTEGER && right.Type() == object.INTEGER:
		return evalIntegerInfixExpression(operator, left, right, line, col)
	case left.Type() == object.FLOAT && right.Type() == object.FLOAT:
		return evalFloatInfixExpression(operator, left, right, line, col)
	case left.Type() == object.STRING && right.Type() == object.STRING:
		return evalStringInfixExpression(operator, left, right, line, col)

	default:
		return newError("unkown operator: %s %s %s; line=%d; col=%d", left.Type(), operator, right.Type(), line, col)
	}
}

func evalIntegerInfixExpression(operator string, left, right object.Object, line, col int) object.Object {
	leftval := left.(*object.Integer).Value
	rightval := right.(*object.Integer).Value

	switch operator {
	case "+":
		return &object.Integer{Value: leftval + rightval}
	case "-":
		return &object.Integer{Value: leftval - rightval}
	case "*":
		return &object.Integer{Value: leftval * rightval}
	case "/":
		return &object.Integer{Value: leftval / rightval}
	case "%":
		return &object.Integer{Value: leftval % rightval}
	default:
		return newError("unkown operator: %s %s %s; line=%d; col=%d", left.Type(), operator, right.Type(), line, col)
	}
}
func evalFloatInfixExpression(operator string, left, right object.Object, line, col int) object.Object {
	leftval := left.(*object.Float).Value
	rightval := right.(*object.Float).Value

	switch operator {
	case "+":
		return &object.Float{Value: leftval + rightval}
	case "-":
		return &object.Float{Value: leftval - rightval}
	case "*":
		return &object.Float{Value: leftval * rightval}
	case "/":
		return &object.Float{Value: leftval / rightval}
	default:
		return newError("unkown operator: %s %s %s; line=%d; col=%d", left.Type(), operator, right.Type(), line, col)
	}
}

func evalStringInfixExpression(operator string, left, right object.Object, line, col int) object.Object {
	leftval := left.(*object.String).Value
	rightval := right.(*object.String).Value

	switch operator {
	case "+":
		return &object.String{Value: leftval + rightval}
	default:
		return newError("unkown operator: %s %s %s; line=%d; col=%d", left.Type(), operator, right.Type(), line, col)
	}
}
