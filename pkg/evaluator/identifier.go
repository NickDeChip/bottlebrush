package evaluator

import (
	"github.com/NickDeChip/bottlebrush/pkg/ast"
	"github.com/NickDeChip/bottlebrush/pkg/evaluator/builtins"
	"github.com/NickDeChip/bottlebrush/pkg/object"
)

func evalIdentifier(node *ast.Identifier, env *object.Environment) object.Object {
	if val, ok := env.Get(node.Value); ok {
		return val.Object
	}
	if builtin, ok := builtins.Builtins[node.Value]; ok {
		return builtin
	}

	return newError("identifier not found: %s; line=%d; col=%d", node.Value, node.Line(), node.Col())

}
