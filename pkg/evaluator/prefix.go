package evaluator

import (
	"github.com/NickDeChip/bottle-brush/pkg/ast"
	"github.com/NickDeChip/bottle-brush/pkg/object"
)

func evalPrefix(node *ast.PrefixExpression, env *object.Environment) object.Object {
	right := Eval(node.Right, env)
	if isError(right) {
		return right
	}
	return evalPrefixExpression(node.Operator, right, node.Token.Line, node.Token.Col)
}

func evalPrefixExpression(operator string, right object.Object, line, col int) object.Object {
	switch operator {
	case "-":
		return evalMinusPrefixOperatorExpression(right, line, col)
	case "!":
		return evalBangPrefixOperatorExpression(right, line, col)
	default:
		return newError("unkown operator: %s%s; line=%d; col=%d", operator, right.Type(), line, col)
	}
}

func evalMinusPrefixOperatorExpression(right object.Object, line, col int) object.Object {
	switch {
	case right.Type() == object.INTEGER:
		value := right.(*object.Integer).Value
		return &object.Integer{Value: -value}
	case right.Type() == object.FLOAT:
		value := right.(*object.Float).Value
		return &object.Float{Value: -value}
	default:
		return newError("unkown operator: -%s; line=%d; col=%d", right.Type(), line, col)
	}
}

func evalBangPrefixOperatorExpression(right object.Object, line, col int) object.Object {
	switch {
	case right.Type() == object.BOOL:
		value := right.(*object.Bool).Value
		return getBool(!value)

	default:
		return newError("unkown operator: !%s; line=%d; col=%d", right.Type(), line, col)
	}
}
