package evaluator

import (
	"github.com/NickDeChip/bottlebrush/pkg/ast"
	"github.com/NickDeChip/bottlebrush/pkg/object"
)

func evalList(node *ast.ListLiteral, env *object.Environment) object.Object {
	elements := evalExpresions(node.Elements, env)
	if len(elements) == 1 && isError(elements[0]) {
		return elements[0]
	}

	return &object.List{Elements: elements}
}
