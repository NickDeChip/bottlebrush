package evaluator

import (
	"github.com/NickDeChip/bottlebrush/pkg/ast"
	"github.com/NickDeChip/bottlebrush/pkg/object"
)

func evalIndex(node *ast.IndexExspression, env *object.Environment) object.Object {
	left := Eval(node.Left, env)
	if isError(left) {
		return left
	}

	index := Eval(node.Index, env)
	if isError(index) {
		return index
	}

	switch {
	case left.Type() == object.LIST && index.Type() == object.INTEGER:
		return evalListIndex(left, index, node.Index.Line(), node.Index.Col())

	default:
		return newError("index operator not supported: %s; line=%d; col=%d", left.Type(), node.Left, node.Col())
	}

}

func evalListIndex(list, index object.Object, line, col int) object.Object {
	listObject := list.(*object.List)
	idx := index.(*object.Integer).Value
	max := int64(len(listObject.Elements) - 1)

	if idx < 0 || idx > max {
		return newError("index %d out of range; line=%d; col=%d", idx, line, col)
	}

	return listObject.Elements[idx]
}
